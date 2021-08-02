package api

import (
	"net/http"

	"github.com/qq51529210/cloud-service/authentication/db"
	"github.com/qq51529210/cloud-service/authentication/reg"
	"github.com/qq51529210/cloud-service/util"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
)

// POST /verification_code?type=xx
func PostVerificationCode(c *router.Context) bool {
	switch c.Req.Form.Get(UrlQueryNameType) {
	case UrlQueryNamePhoneType:
		return postPhoneVerificationCode(c)
	case UrlQueryNameEmailType:
		return postEmailVerificationCode(c)
	}
	return false
}

// POST /verification_code?type=phone
func postPhoneVerificationCode(c *router.Context) bool {
	phoneNumber := c.Req.PostForm.Get(PostFormNamePhoneNumber)
	// Check format.
	if !matchRegexp(c, reg.PhoneNumber, phoneNumber) {
		return false
	}
	// Create a random code and cache it.
	code := c.RandomNumber(PhoneCodeLength)
	err := db.SetPhoneVerificationCode(phoneNumber, code, int64(PhoneCodeExpire))
	if err != nil {
		log.Error(err)
		writeJSON(c, http.StatusInternalServerError, dbErrorJson)
		return false
	}
	// Send verification code to phone.
	err = util.SendSMS(phoneNumber, code)
	if err != nil {
		log.Error(err)
		writeJSON(c, http.StatusInternalServerError, dbErrorJson)
		return false
	}
	// 201
	c.Res.WriteHeader(http.StatusCreated)
	return true
}

// POST /verification_code?type=email
func postEmailVerificationCode(c *router.Context) bool {
	email := c.Req.PostForm.Get(PostFormNameEmail)
	// Check format.
	if !matchRegexp(c, reg.Email, email) {
		return false
	}
	return true
}
