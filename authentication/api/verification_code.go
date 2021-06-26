package api

import (
	"net/http"

	"github.com/qq51529210/cloud-service/authentication/db"
	"github.com/qq51529210/cloud-service/authentication/reg"
	"github.com/qq51529210/cloud-service/util"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
)

// POST /verification_code
func PostVerificationCode(c *router.Context) bool {
	switch c.Req.Form.Get(UrlQueryNameType) {
	case UrlQueryNameTypeValuePhone:
		return postPhoneVerificationCode(c)
	}
	return true
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
	err := db.SetPhoneCode(phoneNumber, code, int64(PhoneCodeExpire))
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, dbJsonError)
		return false
	}
	// Send SMS.
	err = util.SendSMS(phoneNumber, code)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, smsJsonError)
		return false
	}
	// 201
	c.Res.WriteHeader(http.StatusCreated)
	return true
}
