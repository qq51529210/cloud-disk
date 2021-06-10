package api

import (
	"fmt"
	"net/http"

	"github.com/qq51529210/cloud-service/authentication/db"
	"github.com/qq51529210/cloud-service/util"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
	"github.com/qq51529210/uuid"
)

var (
	// SignIn page cache handler.
	SignInPage router.CacheHandler
	// Cookie
	CookieName   = "session_id"
	CookieDomain = ""
	CookiePath   = "/"
	CookieMaxAge = 0
	// Form name.
	FormUsername = "username"
	FormPassword = "password"
	FormNumber   = "number"
	FormCode     = "code"
	// Query name.
	QuerySignType         = "type"
	QuerySignTypePhone    = "phone"
	QuerySignTypeUsername = "username"
	QueryRedirectUrl      = "redirect_url"
	QueryAppID            = "app_id"
	//
	PhoneCodeLength = 6
)

// GET /sign_in
func GetSignIn(c *router.Context) bool {
	// Check cookie token.
	cookie, _ := c.Req.Cookie(CookieName)
	if cookie != nil {
		ok, err := db.HasToken(cookie.Value)
		if err != nil {
			log.Error(err)
			return false
		}
		if ok {
			return ssoRedirect(c)
		}
	}
	// Check header bearer token
	token := c.BearerToken()
	if token != "" {
		ok, err := db.HasToken(token)
		if err != nil {
			log.Error(err)
			return false
		}
		if ok {
			return ssoRedirect(c)
		}
	}
	// Not token, return sign_in page.
	return SignInPage.Handle(c)
}

// Redirect after signed in.
func ssoRedirect(c *router.Context) bool {
	// SSO redirect.
	redirectUrl := c.Req.Form.Get(QueryRedirectUrl)
	if redirectUrl == "" {
		return true
	}
	appId := c.Req.Form.Get(QueryAppID)
	if appId != "" {
		// Generate jwt by app alg and key.
		jwtInfo, err := db.GetAppJwtInfo(appId)
		if err != nil {
			log.Error(err)
			return false
		}
		ssoJWT, err := genHSJWT(jwtInfo.Alg, jwtInfo.Key)
		if err != nil {
			log.Error(err)
			return false
		}
		// Redirect with jwt
		http.Redirect(c.Res, c.Req, fmt.Sprintf(redirectUrl, ssoJWT), http.StatusSeeOther)
	}
	return true
}

// POST /sign_in
func PostSignIn(c *router.Context) bool {
	var ok bool
	// Handle signIn and return token.
	_type := c.Req.Form.Get(QuerySignType)
	switch _type {
	case QuerySignTypePhone:
		ok = phoneSignIn(c)
	case "", QuerySignTypeUsername:
		ok = usernameSignIn(c)
	default:
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("Invalid sign in type %s.", _type),
		})
	}
	// SignIn failed.
	if !ok {
		return false
	}
	token := uuid.LowerV1WithoutHyphen()
	// Cache token.
	err := db.SetToken(token)
	if err != nil {
		log.Error(err)
		return false
	}
	// Set cookie.
	http.SetCookie(c.Res, &http.Cookie{
		Name:     CookieName,
		Value:    token,
		Path:     CookiePath,
		Domain:   CookieDomain,
		HttpOnly: true,
		Secure:   true,
		MaxAge:   CookieMaxAge,
	})
	return ssoRedirect(c)
}

// Sign in with username and password.
func usernameSignIn(c *router.Context) bool {
	username := c.Req.PostForm.Get(FormUsername)
	password := c.Req.PostForm.Get(FormPassword)
	// Check by regular expression.
	if !usernameRegexp.MatchString(username) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid username format.",
			"regexp": UsernameRegexp,
		})
		return false
	}
	if !passwordRegexp.MatchString(password) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid password format.",
			"regexp": PasswordRegexp,
		})
		return false
	}
	// Database.
	user, err := db.GetUserByName(username)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query database error.",
		})
		return false
	}
	// Wrong password.
	if c.SHA1(password) != user.Password {
		c.WriteJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Password incorrect.",
		})
		return false
	}
	// Invalid state.
	if user.State != 0 {
		c.WriteJSON(http.StatusForbidden, map[string]interface{}{
			"error": fmt.Sprintf("Invalid user state %s.", user.StateString()),
		})
		return false
	}
	return true
}

// Sign in with phone number and code.
func phoneSignIn(c *router.Context) bool {
	number := c.Req.PostForm.Get(FormNumber)
	code := c.Req.PostForm.Get(FormCode)
	// Check by regular expression.
	if !phoneNumberRegexp.MatchString(number) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid phone number format.",
			"regexp": PhoneNumberRegexp,
		})
		return false
	}
	if !phoneCodeRegexp.MatchString(code) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid phone code format.",
			"regexp": PhoneCodeRegexp,
		})
		return false
	}
	// Check sms number.
	dbCode, err := db.GetPhoneCode(number)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query database error.",
		})
		return false
	}
	if dbCode == "" || dbCode != code {
		c.WriteJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Wrong phone code.",
		})
	}
	// Database.
	user, err := db.GetUserByPhone(number)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query database error.",
		})
		return false
	}
	// Invalid state.
	if user.State != 0 {
		c.WriteJSON(http.StatusForbidden, map[string]interface{}{
			"error": fmt.Sprintf("Invalid user state %s.", user.StateString()),
		})
		return false
	}
	return true
}

// GET /sign_up
func GetSignUp(c *router.Context) bool {
	return true
}

// POST /sign_up
func PostSignUp(c *router.Context) bool {
	_type := c.Req.Form.Get(QuerySignType)
	switch _type {
	case QuerySignTypePhone:
		return phoneSignUp(c)
	case "", QuerySignTypeUsername:
		return usernameSignUp(c)
	default:
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("Invalid sign up type %s.", _type),
		})
		return false
	}
}

// Sign up with username and password.
func usernameSignUp(c *router.Context) bool {
	username := c.Req.PostForm.Get(FormUsername)
	password := c.Req.PostForm.Get(FormPassword)
	// Check by regular expression.
	if !usernameRegexp.MatchString(username) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid username format.",
			"regexp": UsernameRegexp,
		})
		return false
	}
	if !passwordRegexp.MatchString(password) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid password format.",
			"regexp": PasswordRegexp,
		})
		return false
	}
	// Database.
	user := db.User{
		Id:       uuid.LowerV1WithoutHyphen(),
		Name:     username,
		Password: c.SHA1(password),
		State:    0,
	}
	err := db.CreateUser(&user)
	if err != nil {
		var str string
		if err != db.ErrUserExists {
			str = "Query database error."
		} else {
			str = fmt.Sprintf("Username %s exists", username)
		}
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": str,
		})
		return false
	}
	// 201
	c.Res.WriteHeader(http.StatusCreated)
	return true
}

// Sign up with phone number and code.
func phoneSignUp(c *router.Context) bool {
	return true
}

// POST /phone_code
func PostPhoneCode(c *router.Context) bool {
	number := c.Req.PostForm.Get(FormNumber)
	// Check by regular expression.
	if !phoneNumberRegexp.MatchString(number) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid phone number format.",
			"regexp": PhoneNumberRegexp,
		})
		return false
	}
	// Create a random code.
	code := c.RandomNumber(PhoneCodeLength)
	// Cache it.
	err := db.SetPhoneCode(number, code)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query database error.",
		})
		return false
	}
	// Send code to phone
	err = util.SendSMS(number, code)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query database error.",
		})
		return false
	}
	// 201
	c.Res.WriteHeader(http.StatusCreated)
	return true
}
