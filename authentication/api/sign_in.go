package api

import router "github.com/qq51529210/http-router"

// GET /sign_in
func GetSignIn(c *router.Context) bool {
	// 	// Check cookie token.
	// 	cookie, _ := c.Req.Cookie(CookieName)
	// 	if cookie != nil {
	// 		ok, err := db.HasToken(cookie.Value)
	// 		if err != nil {
	// 			log.Error(err)
	// 			return false
	// 		}
	// 		if ok {
	// 			return ssoRedirect(c)
	// 		}
	// 	}
	// 	// Check header bearer token
	// 	token := c.BearerToken()
	// 	if token != "" {
	// 		ok, err := db.HasToken(token)
	// 		if err != nil {
	// 			log.Error(err)
	// 			return false
	// 		}
	// 		if ok {
	// 			return ssoRedirect(c)
	// 		}
	// 	}
	// 	// Not token, return sign_in page.
	// 	return SignInPage.Handle(c)
	return true
}

// // Redirect after signed in.
// func ssoRedirect(c *router.Context) bool {
// 	// SSO redirect.
// 	redirectUrl := c.Req.Form.Get(UrlQueryRedirectUrl)
// 	if redirectUrl == "" {
// 		return true
// 	}
// 	appId := c.Req.Form.Get(UrlQueryAppID)
// 	if appId != "" {
// 		// Generate jwt by app alg and key.
// 		jwtInfo, err := db.GetAppJwtInfo(appId)
// 		if err != nil {
// 			log.Error(err)
// 			return false
// 		}
// 		ssoJWT, err := genHSJWT(jwtInfo.Alg, jwtInfo.Key)
// 		if err != nil {
// 			log.Error(err)
// 			return false
// 		}
// 		// Redirect with jwt
// 		http.Redirect(c.Res, c.Req, fmt.Sprintf(redirectUrl, ssoJWT), http.StatusSeeOther)
// 	}
// 	return true
// }

// POST /sign_in?type=xx
func PostSignIn(c *router.Context) bool {
	// 	var ok bool
	// 	// Handle signIn and return token.
	// 	_type := c.Req.Form.Get(UrlQueryType)
	// 	switch _type {
	// 	case UrlQueryPhone:
	// 		ok = phoneSignIn(c)
	// 	case "", UrlQueryUsername:
	// 		ok = usernameSignIn(c)
	// 	default:
	// 		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
	// 			"error": fmt.Sprintf("Invalid sign in type %s.", _type),
	// 		})
	// 	}
	// 	// SignIn failed.
	// 	if !ok {
	// 		return false
	// 	}
	// 	token := uuid.LowerV1WithoutHyphen()
	// 	// Cache token.
	// 	err := db.SetToken(token)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return false
	// 	}
	// 	// Set cookie.
	// 	http.SetCookie(c.Res, &http.Cookie{
	// 		Name:     CookieName,
	// 		Value:    token,
	// 		Path:     CookiePath,
	// 		Domain:   CookieDomain,
	// 		HttpOnly: true,
	// 		Secure:   true,
	// 		MaxAge:   CookieMaxAge,
	// 	})
	// 	return ssoRedirect(c)
	return true
}

// // Sign in with username/email/phone and password.
// func usernameSignIn(c *router.Context) bool {
// 	username := c.Req.PostForm.Get(FormNameUsername)
// 	password := c.Req.PostForm.Get(FormNamePassword)
// 	// Check by regular expression.
// 	if !usernameRegexp.MatchString(username) {
// 		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
// 			"error":  "Invalid username format.",
// 			"regexp": UsernameRegexp,
// 		})
// 		return false
// 	}
// 	if !passwordRegexp.MatchString(password) {
// 		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
// 			"error":  "Invalid password format.",
// 			"regexp": PasswordRegexp,
// 		})
// 		return false
// 	}
// 	// Database.
// 	user, err := db.GetUserByNameOrEmailOrPhone(username)
// 	if err != nil {
// 		log.Error(err)
// 		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": "Query database error.",
// 		})
// 		return false
// 	}
// 	// Wrong password.
// 	if c.SHA1(password) != user.Password {
// 		c.WriteJSON(http.StatusUnauthorized, map[string]interface{}{
// 			"error": "Password incorrect.",
// 		})
// 		return false
// 	}
// 	// Invalid state.
// 	if user.State != 0 {
// 		c.WriteJSON(http.StatusForbidden, map[string]interface{}{
// 			"error": fmt.Sprintf("Invalid user state %s.", user.StateString()),
// 		})
// 		return false
// 	}
// 	return true
// }

// // Sign in with phone number and code.
// func phoneSignIn(c *router.Context) bool {
// 	number := c.Req.PostForm.Get(FormNamePhoneNumber)
// 	code := c.Req.PostForm.Get(FormNamePhoneCode)
// 	// Check by regular expression.
// 	if !phoneNumberRegexp.MatchString(number) {
// 		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
// 			"error":  "Invalid phone number format.",
// 			"regexp": PhoneNumberRegexp,
// 		})
// 		return false
// 	}
// 	if !phoneCodeRegexp.MatchString(code) {
// 		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
// 			"error":  "Invalid phone code format.",
// 			"regexp": PhoneCodeRegexp,
// 		})
// 		return false
// 	}
// 	// Check sms number.
// 	dbCode, err := db.GetPhoneCode(number)
// 	if err != nil {
// 		log.Error(err)
// 		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": "Query database error.",
// 		})
// 		return false
// 	}
// 	if dbCode == "" || dbCode != code {
// 		c.WriteJSON(http.StatusUnauthorized, map[string]interface{}{
// 			"error": "Wrong phone code.",
// 		})
// 	}
// 	// Database.
// 	user, err := db.GetUserByPhone(number)
// 	if err != nil {
// 		log.Error(err)
// 		c.WriteJSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": "Query database error.",
// 		})
// 		return false
// 	}
// 	// Invalid state.
// 	if user.State != 0 {
// 		c.WriteJSON(http.StatusForbidden, map[string]interface{}{
// 			"error": fmt.Sprintf("Invalid user state %s.", user.StateString()),
// 		})
// 		return false
// 	}
// 	return true
// }

// // POST /phone_code
// func PostPhoneCode(c *router.Context) bool {
// 	number := c.Req.PostForm.Get(FormNamePhoneNumber)
// 	// Check format.
// 	if !matchRegexp(c, reg.PhoneNumber, number) {
// 		return false
// 	}
// 	// Create a random code and cache it.
// 	code := c.RandomNumber(PhoneCodeLength)
// 	err := db.SetPhoneCode(number, code)
// 	if err != nil {
// 		log.Error(err)
// 		c.WriteJSON(http.StatusInternalServerError, dbJsonError)
// 		return false
// 	}
// 	// Send code to phone
// 	err = util.SendSMS(number, code)
// 	if err != nil {
// 		log.Error(err)
// 		c.WriteJSON(http.StatusInternalServerError, smsJsonError)
// 		return false
// 	}
// 	// 201
// 	c.Res.WriteHeader(http.StatusCreated)
// 	return true
// }

// // GET /account_activation?id=xx
// func GetActivation(c *router.Context) bool {
// 	return false
// }
