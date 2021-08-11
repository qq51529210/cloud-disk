package api

import (
	"fmt"
	"net/http"

	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/db"
	"github.com/qq51529210/micro-services/auth/reg"
	"github.com/qq51529210/micro-services/util"
	"github.com/qq51529210/uuid"
)

type postUserReq struct {
	Account          string `json:"account,omitempty"`
	Password         string `json:"password,omitempty"`
	VerificationCode string `json:"verificationCode,omitempty"`
	Type             string `json:"type,omitempty"` // phone or email
}

func postUsers(c *router.Context) bool {
	// Read model data.
	var m1 postUserReq
	if !readJSON(c, &m1) {
		return false
	}
	// Validate fields.
	if !validateAccount(c, m1.Type, m1.Account) {
		return false
	}
	if !reg.Password.MatchString(m1.Password) {
		c.WriteJSONBytes(http.StatusBadRequest, errPasswordFormat)
		return false
	}
	// Validate code.
	code, err := db.GetVerificationCode(m1.Account)
	if err != nil {
		log.Error(err)
		c.WriteJSONBytes(http.StatusInternalServerError, errQueryDB)
		return false
	}
	if m1.VerificationCode != code {
		c.WriteJSONBytes(http.StatusBadRequest, errVerificationCode)
		return false
	}
	// Update database.
	var m2 db.User
	m2.Account = m1.Account
	m2.Password = c.SHA1(m1.Password)
	m2.Name = fmt.Sprintf("user_%s", util.Format10To62(uuid.SnowflakeID()))
	_, err = m2.Insert()
	if err != nil {
		if db.IsExistedError(err) {
			c.WriteJSON(http.StatusConflict, map[string]string{"error": fmt.Sprintf("Account %s existed", m2.Account)})
		} else {
			log.Error(err)
			c.WriteJSONBytes(http.StatusInternalServerError, errQueryDB)
		}
		return false
	}
	c.Res.WriteHeader(http.StatusCreated)
	return true
}
