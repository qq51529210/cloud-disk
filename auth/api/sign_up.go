package api

import (
	"database/sql"
	"net/http"

	"github.com/qq51529210/cloud-service/authentication/db"
	"github.com/qq51529210/cloud-service/authentication/reg"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
)

// POST /sign_up?type=xx
func PostSignUp(c *router.Context) bool {
	switch c.Req.Form.Get(UrlQueryNameType) {
	case UrlQueryNamePhoneType:
		return phoneSignUp(c)
	case "", UrlQueryNameEmailType:
		return emailSignUp(c)
	default:
		return false
	}
}

// POST /sign_up?type=email
func emailSignUp(c *router.Context) bool {
	var user db.User
	user.Account.String = c.Req.PostForm.Get(PostFormNameAccount)
	user.Password.String = c.Req.PostForm.Get(PostFormNamePassword)
	user.Email.String = c.Req.PostForm.Get(PostFormNameEmail)
	// Check format.
	if !matchRegexp(c, reg.Account, user.Account.String) ||
		!matchRegexp(c, reg.Password, user.Password.String) ||
		!matchRegexp(c, reg.Email, user.Email.String) {
		return false
	}
	// Account or email existed?
	err := user.SelectIdByAccountOrEmail()
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			writeJSON(c, http.StatusInternalServerError, dbErrorJson)
			return false
		}
	} else {
		writeJSON(c, http.StatusForbidden, accountOrEmailExistedErrorJson)
		return false
	}
	// Insert.
	user.Password.String = c.SHA1(user.Password.String)
	_, err = user.InsertAccountPasswordEmail()
	if err != nil {
		if db.IsExistedError(err) {
			writeJSON(c, http.StatusForbidden, accountOrEmailExistedErrorJson)
		} else {
			log.Error(err)
			writeJSON(c, http.StatusInternalServerError, dbErrorJson)
		}
		return false
	}
	return signUp(c)
}

// POST /sign_up?type=phone
func phoneSignUp(c *router.Context) bool {
	var user db.User
	user.Password.String = c.Req.PostForm.Get(PostFormNamePassword)
	user.Phone.String = c.Req.PostForm.Get(PostFormNamePhoneNumber)
	phoneVerificationCode := c.Req.PostForm.Get(PostFormNamePhoneVerificationCode)
	// Check format.
	if !matchRegexp(c, reg.Password, user.Password.String) || !matchRegexp(c, reg.PhoneNumber, user.Phone.String) ||
		!matchRegexp(c, reg.PhoneVerificationCode, phoneVerificationCode) {
		return false
	}
	// Check phone verification code.
	code, err := db.GetPhoneVerificationCode(user.Phone.String)
	if err != nil {
		log.Error(err)
		writeJSON(c, http.StatusInternalServerError, dbErrorJson)
		return false
	}
	if phoneVerificationCode != code {
		writeJSON(c, http.StatusBadRequest, verificationCodeErrorJson)
		return false
	}
	// Phone number existed?
	err = user.SelectIdByPhone()
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			writeJSON(c, http.StatusInternalServerError, dbErrorJson)
			return false
		}
	} else {
		writeJSON(c, http.StatusForbidden, phoneExistedErrorJson)
		return false
	}
	// Insert.
	user.Password.String = c.SHA1(user.Password.String)
	_, err = user.InsertPhonePassword()
	if err != nil {
		if db.IsExistedError(err) {
			writeJSON(c, http.StatusForbidden, phoneExistedErrorJson)
		} else {
			log.Error(err)
			writeJSON(c, http.StatusInternalServerError, dbErrorJson)
		}
		return false
	}
	return signUp(c)
}

func signUp(c *router.Context) bool {
	c.Res.WriteHeader(http.StatusCreated)
	return true
}
