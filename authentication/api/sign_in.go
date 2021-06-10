package api

import (
	"fmt"
	"net/http"

	"github.com/qq51529210/cloud-service/authentication/db"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
	"github.com/qq51529210/uuid"
)

var (
	// SignIn page cache handler.
	SignInPage   router.CacheHandler
	CookieName   = "session_id"
	CookieDomain = ""
	CookiePath   = "/"
	CookieMaxAge = 0
)

// Handle GET /sign_in
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

// Handle sso redirect.
func ssoRedirect(c *router.Context) bool {
	// SSO redirect.
	redirectUrl := c.Req.Form.Get("redirect_url")
	if redirectUrl == "" {
		return true
	}
	appId := c.Req.Form.Get("app_id")
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

// Handle POST /sign_in
func PostSignIn(c *router.Context) bool {
	var ok bool
	// Handle signIn and return token.
	_type := c.Req.Form.Get("type")
	switch _type {
	case "mobile":
		ok = phoneSignIn(c)
	case "", "username":
		ok = userSignIn(c)
	default:
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Errorf("Invalid sign in type.", _type),
		})
	}
	// SignIn failed.
	if !ok {
		return false
	}
	// Set cookie.
	http.SetCookie(c.Res, &http.Cookie{
		Name:     CookieName,
		Value:    uuid.LowerV1WithoutHyphen(),
		Path:     CookiePath,
		Domain:   CookieDomain,
		HttpOnly: true,
		Secure:   true,
		MaxAge:   CookieMaxAge,
	})
	return ssoRedirect(c)
}

// Handle username sign in.
func userSignIn(c *router.Context) bool {
	username := c.Req.PostForm.Get("username")
	password := c.Req.PostForm.Get("password")
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
			"error": fmt.Errorf("Invalid user state.", user.StateString()),
		})
		return false
	}
	return true
}

// Handle phone sign in.
func phoneSignIn(c *router.Context) bool {
	phone := c.Req.PostForm.Get("phone")
	sms := c.Req.PostForm.Get("sms")
	// Check by regular expression.
	if !phoneNumberRegexp.MatchString(phone) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid phone number format.",
			"regexp": PhoneNumberRegexp,
		})
		return false
	}
	if !phoneSMSRegexp.MatchString(sms) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error":  "Invalid phone sms format.",
			"regexp": PhoneSMSRegexp,
		})
		return false
	}
	// Check sms number.
	smsNumber, err := db.GetSMSNumber(phone)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query database error.",
		})
		return false
	}
	if smsNumber == "" || smsNumber != sms {
		c.WriteJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Wrong sms code.",
		})
	}
	// Database.
	user, err := db.GetUserByPhone(phone)
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
			"error": fmt.Errorf("Invalid user state.", user.StateString()),
		})
		return false
	}
	return true
}
