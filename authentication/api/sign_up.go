package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/qq51529210/cloud-service/authentication/db"
	"github.com/qq51529210/cloud-service/authentication/reg"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
)

// POST /sign_up?type=xx
func PostSignUp(c *router.Context) bool {
	switch c.Req.Form.Get(UrlQueryNameType) {
	case UrlQueryNameTypeValuePhone:
		// type=phone
		return phoneSignUp(c)
	}
	return true
}

func phoneSignUp(c *router.Context) bool {
	// Username sign up.
	account := c.Req.PostForm.Get(PostFormNameAccount)
	password := c.Req.PostForm.Get(PostFormNamePassword)
	phoneNumber := c.Req.PostForm.Get(PostFormNamePhoneNumber)
	phoneVerificationCode := c.Req.PostForm.Get(PostFormNamePhoneVerificationCode)
	// Check format.
	if !matchRegexp(c, reg.Account, account) ||
		!matchRegexp(c, reg.Password, password) ||
		!matchRegexp(c, reg.PhoneNumber, phoneNumber) ||
		!matchRegexp(c, reg.PhoneVerificationCode, phoneVerificationCode) {
		return false
	}
	// Check phone number and code pair.
	code, err := db.GetPhoneCode(phoneNumber)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, dbJsonError)
		return false
	}
	if phoneVerificationCode != code {
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("Invalid code %s", phoneVerificationCode),
		})
		return false
	}
	// Create new user.
	var user db.User
	user.Phone = phoneNumber
	// Existed?
	err = user.SelectIdByPhone()
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			c.WriteJSON(http.StatusInternalServerError, dbJsonError)
			return false
		}
	} else {
		c.WriteJSON(http.StatusForbidden, map[string]interface{}{
			"error": fmt.Sprintf("Phone number %s existed", phoneNumber),
		})
		return false
	}
	// Create.
	user.Account.String = account
	user.Account.Valid = true
	user.Password.String = c.SHA1(password)
	user.Password.Valid = true
	user.Name.String = account
	user.Name.Valid = true
	_, err = user.Insert()
	if err != nil {
		if db.IsExistedError(err) {
			c.WriteJSON(http.StatusForbidden, map[string]interface{}{
				"error": fmt.Sprintf("Phone number %s existed", phoneNumber),
			})
		} else {
			log.Error(err)
			c.WriteJSON(http.StatusInternalServerError, dbJsonError)
		}
		return false
	}
	// 201
	c.Res.WriteHeader(http.StatusCreated)
	return true
}
