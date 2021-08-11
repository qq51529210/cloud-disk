package api

import (
	"net/http"

	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/db"
)

type postVerificationCodeReq struct {
	Account string `json:"account,omitempty"`
	Type    string `json:"type,omitempty"` // phone or email
}

type postVerificationCodeRes struct {
	Expired int64 `json:"expired,omitempty"`
}

func postVerificationCodes(c *router.Context) bool {
	// Read model data.
	var m1 postVerificationCodeReq
	if !readJSON(c, &m1) {
		return false
	}
	// Validate fields.
	if !validateAccount(c, m1.Type, m1.Account) {
		return false
	}
	// Cache verification code
	code := c.RandomNumber(verificationCodeLength)
	err := db.SetVerificationCode(m1.Account, code, verificationCodeExpired)
	if err != nil {
		log.Error(err)
		c.WriteJSONBytes(http.StatusInternalServerError, errQueryDB)
		return false
	}
	switch m1.Type {
	case "", "phone":
		err = sendVerificationCodeToPhone(m1.Account, code)
	case "email":
		err = sendVerificationCodeToEmail(m1.Account, code)
	}
	if err != nil {
		log.Error(err)
	}
	c.WriteJSON(http.StatusCreated, &postVerificationCodeRes{
		Expired: verificationCodeExpired,
	})
	return false
}

// // POST /verification_code?type=phone
// func postPhoneVerificationCode(c *router.Context) bool {
// 	phoneNumber := c.Req.PostForm.Get(PostFormNamePhoneNumber)
// 	// Check format.
// 	if !matchRegexp(c, reg.PhoneNumber, phoneNumber) {
// 		return false
// 	}
// 	// Create a random code and cache it.
// 	code := c.RandomNumber(PhoneCodeLength)
// 	err := db.SetPhoneVerificationCode(phoneNumber, code, int64(PhoneCodeExpire))
// 	if err != nil {
// 		log.Error(err)
// 		writeJSON(c, http.StatusInternalServerError, dbErrorJson)
// 		return false
// 	}
// 	// Send verification code to phone.
// 	err = util.SendSMS(phoneNumber, code)
// 	if err != nil {
// 		log.Error(err)
// 		writeJSON(c, http.StatusInternalServerError, dbErrorJson)
// 		return false
// 	}
// 	// 201
// 	c.Res.WriteHeader(http.StatusCreated)
// 	return true
// }

// // POST /verification_code?type=email
// func postEmailVerificationCode(c *router.Context) bool {
// 	email := c.Req.PostForm.Get(PostFormNameEmail)
// 	// Check format.
// 	if !matchRegexp(c, reg.Email, email) {
// 		return false
// 	}
// 	return true
// }
