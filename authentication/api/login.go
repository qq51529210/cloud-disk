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
	// Login page cache handler.
	LoginPage    router.CacheHandler
	CookieName   = "session_id"
	CookieDomain = ""
	CookiePath   = "/"
	CookieMaxAge = 0
)

// Handle GET /login
func GetLogin(c *router.Context) bool {
	// Check cookie token.
	cookie, _ := c.Req.Cookie(CookieName)
	if cookie != nil {
		ok, err := db.HasToken(cookie.Value)
		if err != nil {
			log.Error(err)
			return false
		}
		if ok {
			return login(c, cookie.Value)
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
			return login(c, cookie.Value)
		}
	}
	// Not token, return login page.
	return LoginPage.Handle(c)
}

func login(c *router.Context, token string) bool {
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
	// SSO redirect.
	redirectUrl := c.Req.Form.Get("redirect_url")
	if redirectUrl != "" {
		// Is our system.
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
	}
	return true
}

// Handle POST /login
func PostLogin(c *router.Context) bool {
	// Handle login and return token.
	var token string
	switch c.Req.Form.Get("type") {
	case "mobile":
		token = phoneLogin(c)
	default:
		token = userLogin(c)
	}
	// Login failed.
	if token == "" {
		return false
	}
	return login(c, token)
}

// Handle username login.
func userLogin(c *router.Context) string {
	username := c.Req.PostForm.Get("username")
	password := c.Req.PostForm.Get("password")
	// Check form value by regular expression.
	if !usernameRegexp.MatchString(username) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Wrong username.",
			"code":  1,
		})
		return ""
	}
	if !passwordRegexp.MatchString(password) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Wrong password.",
			"code":  2,
		})
		return ""
	}
	// Database.
	user, err := db.GetUserByName(username)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query error.",
			"code":  3,
		})
		return ""
	}
	// Wrong password.
	if c.SHA1(password) != user.Password {
		c.WriteJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Wrong password.",
			"code":  2,
		})
		return ""
	}
	// Invalid state.
	if user.State != 0 {
		c.WriteJSON(http.StatusForbidden, map[string]interface{}{
			"error": "Invalid user state.",
			"code":  4,
		})
		return ""
	}
	// Create uuid token.
	return uuid.LowerV1WithoutHyphen()
}

// Handle phone login.
func phoneLogin(c *router.Context) string {
	phone := c.Req.PostForm.Get("phone")
	sms := c.Req.PostForm.Get("sms")
	// Check form value by regular expression.
	if !phoneNumberRegexp.MatchString(phone) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Wrong phone number.",
			"code":  1,
		})
		return ""
	}
	if !phoneSMSRegexp.MatchString(sms) {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Wrong sms number.",
			"code":  2,
		})
		return ""
	}
	// Check sms number.
	smsNumber, err := db.GetSMSNumber(phone)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query error.",
			"code":  3,
		})
	}
	if smsNumber == "" || smsNumber != sms {
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Wrong sms number.",
			"code":  2,
		})
	}
	// Database.
	user, err := db.GetUserByPhone(phone)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Query error",
			"code":  3,
		})
		return ""
	}
	// Invalid state.
	if user.State != 0 {
		c.WriteJSON(http.StatusForbidden, map[string]interface{}{
			"error": "Invalid user state",
			"code":  4,
		})
		return ""
	}
	// Create uuid token.
	return uuid.LowerV1WithoutHyphen()
}
