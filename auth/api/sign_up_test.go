package api

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/qq51529210/cloud-service/authentication/db"
	router "github.com/qq51529210/http-router"
)

func Test_phoneSignUp(t *testing.T) {
	// DB.
	testInitDB(t)
	phoneCode := "123321"
	var user db.User
	user.Account.String = "user_1"
	user.Password.String = "password_1"
	user.Phone.String = "18712312312"
	user.Name.String = user.Account.String
	_, err := (&user).DeleteByAccount()
	if err != nil {
		t.Fatal(err)
	}
	err = db.SetPhoneCode(user.Phone, phoneCode, int64(PhoneCodeExpire))
	if err != nil {
		t.Fatal(err)
	}
	// Request.
	var req http.Request
	var res testResponseWriter
	res.Reset()
	req.Form = make(url.Values)
	req.Form.Set(UrlQueryNameType, UrlQueryNameTypeValuePhone)
	req.PostForm = make(url.Values)
	req.PostForm.Set(PostFormNameAccount, user.Account.String)
	req.PostForm.Set(PostFormNamePassword, user.Password.String)
	req.PostForm.Set(PostFormNamePhoneNumber, user.Phone)
	req.PostForm.Set(PostFormNamePhoneVerificationCode, phoneCode)
	if !phoneSignUp(&router.Context{Req: &req, Res: &res}) {
		t.FailNow()
	}
	if res.statusCode != http.StatusCreated {
		t.FailNow()
	}
}
