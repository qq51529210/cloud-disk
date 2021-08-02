package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/qq51529210/cloud-service/authentication/reg"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/jwt"
	"github.com/qq51529210/log"
	"google.golang.org/grpc"
)

var (
	// HTTP route paths.
	RoutePathSignIn           = "/sign_in"
	RoutePathSignUp           = "/sign_up"
	RoutePathVerificationCode = "/verification_code"
	// Cookie informations.
	CookieName   = "session_id"
	CookieDomain = ""
	CookiePath   = "/"
	CookieMaxAge = 0
	// Random phone code length.
	PhoneCodeLength = 6
	PhoneCodeExpire = 60
	// Form names.
	PostFormNameAccount               = "account"
	PostFormNamePassword              = "password"
	PostFormNamePhoneNumber           = "phone_number"
	PostFormNamePhoneVerificationCode = "verification_code"
	PostFormNameEmail                 = "email"
	// Query names.
	UrlQueryNameType        = "type"
	UrlQueryNamePhoneType   = "phone"
	UrlQueryNameEmailType   = "email"
	UrlQueryNameRedirectUrl = "redirect_url"
	UrlQueryNameAppID       = "app_id"
	// Error response json data.
	dbErrorJson                    []byte
	smsErrorJson                   []byte
	phoneExistedErrorJson          []byte
	accountOrEmailExistedErrorJson []byte
	verificationCodeErrorJson      []byte
	userNotExistsErrorJson         []byte
	// Sign in page handler.
	signInPage *router.Route
)

func init() {
	// Error response json data.
	dbErrorJson, _ = json.Marshal(map[string]interface{}{
		"error": "Query database error.",
	})
	smsErrorJson, _ = json.Marshal(map[string]interface{}{
		"error": "Send phone verification code error.",
	})
	phoneExistedErrorJson, _ = json.Marshal(map[string]interface{}{
		"error": "Phone number existed.",
	})
	accountOrEmailExistedErrorJson, _ = json.Marshal(map[string]interface{}{
		"error": "Account or email existed.",
	})
	verificationCodeErrorJson, _ = json.Marshal(map[string]interface{}{
		"error": "Invalid phone verification code.",
	})
	userNotExistsErrorJson, _ = json.Marshal(map[string]interface{}{
		"error": "User not exists.",
	})
}

func InitGRPG() *grpc.Server {
	return nil
}

func InitRouter(r *router.Router, pageDir string) error {
	// Add all static files.
	err := r.AddStatic(http.MethodGet, "/", pageDir, true)
	if err != nil {
		return err
	}
	// Get sign in page cache.
	p := path.Join(pageDir, "sign_in.html")
	signInPage = r.RouteGet(p)
	if signInPage == nil {
		return fmt.Errorf("counld not find sign in page html file in %s", p)
	}
	// get /sign_in
	_, err = r.AddGet(RoutePathSignIn, append([]router.HandlerFunc{parseForm}, signInPage.Handler...)...)
	if err != nil {
		return err
	}
	// post /sign_in
	_, err = r.AddPost(RoutePathSignIn, parseForm, PostSignIn)
	if err != nil {
		return err
	}
	// post /sign_up
	_, err = r.AddPost(RoutePathSignUp, parseForm, PostSignUp)
	if err != nil {
		return err
	}
	// post /verification_code
	_, err = r.AddPost(RoutePathVerificationCode, parseForm, PostVerificationCode)
	if err != nil {
		return err
	}
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
func parseForm(c *router.Context) bool {
	err := c.Req.ParseForm()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
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

func writeJSON(c *router.Context, statusCode int, json []byte) {
	c.Res.Header().Set("Content-Type", router.ContentTypeJSON)
	c.Res.WriteHeader(statusCode)
	c.Res.Write(json)
}

func sendEmail(email string) error {
	return nil
}
