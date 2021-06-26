package api

import (
	"encoding/json"
	"net/http"

	"github.com/qq51529210/cloud-service/authentication/reg"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/jwt"
	"github.com/qq51529210/log"
	"google.golang.org/grpc"
)

const (
	// HTTP route paths.
	RoutePathSignIn = "/sign_in"
	RoutePathSignUp = "/sign_up"
	// Form names.
	PostFormNameAccount               = "account"
	PostFormNamePassword              = "password"
	PostFormNamePhoneNumber           = "phone_number"
	PostFormNamePhoneVerificationCode = "verification_code"
	// Query names.
	UrlQueryNameType           = "type"
	UrlQueryNameTypeValuePhone = "phone"
	UrlQueryNameRedirectUrl    = "redirect_url"
	UrlQueryNameAppID          = "app_id"
)

var (
	// Cookie informations.
	CookieName   = "session_id"
	CookieDomain = ""
	CookiePath   = "/"
	CookieMaxAge = 0
	// Random phone code length.
	PhoneCodeLength = 6
	PhoneCodeExpire = 60
	// Error response json data.
	dbJsonError  []byte
	smsJsonError []byte
	// Handlers
	Handlers = make(map[string][]router.HandleFunc)
)

func init() {
	var err error
	// Error response json data.
	dbJsonError, err = json.Marshal(map[string]interface{}{
		"error": "Query database error.",
	})
	if err != nil {
		panic(err)
	}
	smsJsonError, err = json.Marshal(map[string]interface{}{
		"error": "Send phone verification code error.",
	})
	if err != nil {
		panic(err)
	}
	// Handlers
	Handlers[RoutePathSignIn] = []router.HandleFunc{ParseForm, PostSignIn}
	Handlers[RoutePathSignUp] = []router.HandleFunc{ParseForm, PostSignUp}
}

func genHSJWT(alg, key string) (string, error) {
	header := make(map[string]interface{})
	payload := make(map[string]interface{})
	switch alg {
	case string(jwt.HS384):
		return jwt.GenerateHS384(header, payload, []byte(key))
	case string(jwt.HS512):
		return jwt.GenerateHS512(header, payload, []byte(key))
	default:
		return jwt.GenerateHS256(header, payload, []byte(key))
	}
}

func InitGRPG() *grpc.Server {
	return nil
}

// If s not match regular expression r, setup c.Res and return false.
func matchRegexp(c *router.Context, r reg.Regexp, s string) bool {
	err := r.Match(s)
	if err != nil {
		c.Res.WriteHeader(http.StatusBadRequest)
		c.Res.Header().Add("Content-Type", router.ContentTypeJSON)
		c.Res.Write(err)
		return false
	}
	return true
}

// Parse form before handle, return false if there is a error.
func ParseForm(c *router.Context) bool {
	err := c.Req.ParseForm()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
