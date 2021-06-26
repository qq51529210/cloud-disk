package api

import (
	"net/http"
	"net/url"
	"testing"

	router "github.com/qq51529210/http-router"
)

func Test_postPhoneVerificationCode(t *testing.T) {
	// DB.
	testInitDB(t)
	// Request.
	var req http.Request
	var res testResponseWriter
	req.Form = make(url.Values)
	req.Form.Set(UrlQueryNameType, UrlQueryNameTypeValuePhone)
	req.PostForm = make(url.Values)
	req.PostForm.Set(PostFormNamePhoneNumber, "18712345678")
	res.Reset()
	if !postPhoneVerificationCode(&router.Context{Req: &req, Res: &res}) {
		t.FailNow()
	}
	if res.statusCode != http.StatusCreated {
		t.FailNow()
	}
	// Invalid phone number format.
	req.PostForm.Set(PostFormNamePhoneNumber, "1871234")
	res.Reset()
	if postPhoneVerificationCode(&router.Context{Req: &req, Res: &res}) {
		t.FailNow()
	}
	if res.statusCode != http.StatusBadRequest {
		t.FailNow()
	}
}
