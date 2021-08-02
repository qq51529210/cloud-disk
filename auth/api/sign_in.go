package api

import (
	"database/sql"
	"fmt"
	"net/http"

	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/db"
	"github.com/qq51529210/micro-services/auth/reg"
	"github.com/qq51529210/uuid"
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
			return signInRedirect(c)
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
			return signInRedirect(c)
		}
	}
	// Not token, return sign_in page.
	return signInPage.Handle(c)
}

// Redirect after signed in.
func signInRedirect(c *router.Context) bool {
	// SSO redirect.
	redirectUrl := c.Req.Form.Get(UrlQueryNameRedirectUrl)
	if redirectUrl == "" {
		return true
	}
	appId := c.Req.Form.Get(UrlQueryNameAppID)
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

// POST /sign_in?type=xx
func PostSignIn(c *router.Context) bool {
	// Handle signIn and return token.
	switch c.Req.Form.Get(UrlQueryNameType) {
	case UrlQueryNameTypeValuePhone:
		return phoneSignIn(c)
	default:
		return false
	}
}

// Sign in with phone number and code.
func phoneSignIn(c *router.Context) bool {
	phoneNumber := c.Req.PostForm.Get(PostFormNamePhoneNumber)
	phoneVerificationCode := c.Req.PostForm.Get(PostFormNamePhoneVerificationCode)
	// Check format.
	if !matchRegexp(c, reg.PhoneNumber, phoneNumber) ||
		!matchRegexp(c, reg.PhoneVerificationCode, phoneVerificationCode) {
		return false
	}
	// Check phone verification code.
	code, err := db.GetPhoneCode(phoneNumber)
	if err != nil {
		log.Error(err)
		writeJSON(c, http.StatusInternalServerError, dbErrorJson)
		return false
	}
	if phoneVerificationCode != code {
		writeJSON(c, http.StatusUnauthorized, verificationCodeErrorJson)
		return false
	}
	// Database.
	var user db.User
	user.Phone = phoneNumber
	err = user.SelectByAccountOrPhone()
	if err != nil {
		if err == sql.ErrNoRows {
			writeJSON(c, http.StatusNotFound, userNotExistsErrorJson)
		} else {
			writeJSON(c, http.StatusInternalServerError, dbErrorJson)
		}
		return false
	}
	return true
}

func setSignInSession(c *router.Context) bool {
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
	return signInRedirect(c)
}
