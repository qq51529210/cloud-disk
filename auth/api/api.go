package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/reg"
	"github.com/qq51529210/micro-services/util"
)

var (
	//
	cookieName   = "session_id"
	cookieDomain = ""
	cookiePath   = ""
	cookieMaxAge = int64(0)
	//
	verificationCodeLength  int   = 6
	verificationCodeExpired int64 = 60
	//
	errQueryDB          []byte
	errReadJSONHeader   []byte
	errReadJSONBody     []byte
	errVerificationCode []byte
	errPhoneFormat      []byte
	errEmailFormat      []byte
	errPasswordFormat   []byte
	errUnsupportedType  []byte
)

type Config struct {
	Cookie                  util.CookieConfig `json:"cookie,omitempty"`
	VerificationCodeLength  int               `json:"verificationCodeLength,omitempty"`
	VerificationCodeExpired int64             `json:"verificationCodeExpired,omitempty"`
}

// Init init http and http2 router handlers.
// Must init reg package first
func Init(r *router.Router, c *Config) {
	if c.VerificationCodeLength > 0 {
		verificationCodeLength = c.VerificationCodeLength
	}
	if c.VerificationCodeExpired > 0 {
		verificationCodeExpired = c.VerificationCodeExpired
	}
	initCookie(&c.Cookie)
	initErrRes()
	initRouter(r)
}

func initErrRes() {
	errQueryDB, _ = json.Marshal(map[string]interface{}{
		"error": "Query database",
	})
	errReadJSONHeader, _ = json.Marshal(map[string]interface{}{
		"error": fmt.Sprintf("Only supported %s and %s request",
			router.ContentTypeJSON, router.UTF8),
	})
	errReadJSONBody, _ = json.Marshal(map[string]interface{}{
		"error": "Read JSON from request body",
	})
	errVerificationCode, _ = json.Marshal(map[string]interface{}{
		"error": "Invalid verification code",
	})
	errPhoneFormat, _ = json.Marshal(map[string]interface{}{
		"error": "Invalid phone format",
		"field": "phone",
		"expr":  reg.Phone.String(),
	})
	errEmailFormat, _ = json.Marshal(map[string]interface{}{
		"error": "Invalid email format",
		"field": "email",
		"expr":  reg.Email.String(),
	})
	errPasswordFormat, _ = json.Marshal(map[string]interface{}{
		"error": "Invalid password format",
		"field": "password",
		"expr":  reg.Password.String(),
	})
	errUnsupportedType, _ = json.Marshal(map[string]interface{}{
		"error": "Unsupported type",
	})
}

func initCookie(c *util.CookieConfig) {
	if c.Name != "" {
		cookieName = c.Name
	}
	if c.Domain != "" {
		cookieDomain = c.Domain
	}
	if c.Path != "" {
		cookiePath = c.Path
	}
	if c.MaxAge != 0 {
		cookieMaxAge = c.MaxAge
	}
}

func initRouter(r *router.Router) {
	r.SetBefore(func(c *router.Context) bool {
		log.Debugf("%s %s %s", c.Req.RemoteAddr, c.Req.Method, c.Req.RequestURI)
		return true
	})
	r.AddPost("/users", postUsers)
	r.AddPost("/verification_codes", postVerificationCodes)
}

func readJSON(c *router.Context, m interface{}) bool {
	contentType := strings.ToLower(c.Req.Header.Get("Content-Type"))
	if strings.Contains(contentType, router.ContentTypeJSON) &&
		strings.Contains(contentType, router.UTF8) {
		err := json.NewDecoder(c.Req.Body).Decode(m)
		if err != nil {
			log.ErrorStack(1, err)
			c.WriteJSONBytes(http.StatusBadRequest, errReadJSONBody)
			return false
		}
		return true
	}
	c.WriteJSONBytes(http.StatusBadRequest, errReadJSONHeader)
	return false
}

func validateAccount(c *router.Context, _type, account string) bool {
	switch _type {
	case "", "phone":
		if !reg.Phone.MatchString(account) {
			c.WriteJSONBytes(http.StatusBadRequest, errPhoneFormat)
			return false
		}
	case "email":
		if !reg.Email.MatchString(account) {
			c.WriteJSONBytes(http.StatusBadRequest, errEmailFormat)
			return false
		}
	default:
		c.WriteJSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("Unsupported type %s", _type),
		})
		return false
	}
	return true
}

func sendVerificationCodeToPhone(number, code string) error {
	log.Debug(number, code)
	return nil
}

func sendVerificationCodeToEmail(email, code string) error {
	log.Debug(email, code)
	return nil
}

// var (
// 	// HTTP route paths.
// 	RoutePathSignIn           = "/sign_in"
// 	RoutePathSignUp           = "/sign_up"
// 	RoutePathVerificationCode = "/verification_code"
// 	// Cookie informations.
// 	CookieName   = "session_id"
// 	CookieDomain = ""
// 	CookiePath   = "/"
// 	CookieMaxAge = 0
// 	// Random phone code length.
// 	PhoneCodeLength = 6
// 	PhoneCodeExpire = 60
// 	// Form names.
// 	PostFormNameAccount               = "account"
// 	PostFormNamePassword              = "password"
// 	PostFormNamePhoneNumber           = "phone_number"
// 	PostFormNamePhoneVerificationCode = "verification_code"
// 	PostFormNameEmail                 = "email"
// 	// Query names.
// 	UrlQueryNameType        = "type"
// 	UrlQueryNamePhoneType   = "phone"
// 	UrlQueryNameEmailType   = "email"
// 	UrlQueryNameRedirectUrl = "redirect_url"
// 	UrlQueryNameAppID       = "app_id"
// 	// Error response json data.
// 	dbErrorJson                    []byte
// 	smsErrorJson                   []byte
// 	phoneExistedErrorJson          []byte
// 	accountOrEmailExistedErrorJson []byte
// 	verificationCodeErrorJson      []byte
// 	userNotExistsErrorJson         []byte
// 	// Sign in page handler.
// 	signInPage *router.Route
// )

// func init() {
// 	// Error response json data.
// 	dbErrorJson, _ = json.Marshal(map[string]interface{}{
// 		"error": "Query database error.",
// 	})
// 	smsErrorJson, _ = json.Marshal(map[string]interface{}{
// 		"error": "Send phone verification code error.",
// 	})
// 	phoneExistedErrorJson, _ = json.Marshal(map[string]interface{}{
// 		"error": "Phone number existed.",
// 	})
// 	accountOrEmailExistedErrorJson, _ = json.Marshal(map[string]interface{}{
// 		"error": "Account or email existed.",
// 	})
// 	verificationCodeErrorJson, _ = json.Marshal(map[string]interface{}{
// 		"error": "Invalid phone verification code.",
// 	})
// 	userNotExistsErrorJson, _ = json.Marshal(map[string]interface{}{
// 		"error": "User not exists.",
// 	})
// }

// func Init() {

// }

// func InitGRPG() *grpc.Server {
// 	return nil
// }

// func InitRouter(r *router.Router, pageDir string) error {
// 	// Add all static files.
// 	err := r.AddStatic(http.MethodGet, "/", pageDir, true)
// 	if err != nil {
// 		return err
// 	}
// 	// Get sign in page cache.
// 	p := path.Join(pageDir, "sign_in.html")
// 	signInPage = r.RouteGet(p)
// 	if signInPage == nil {
// 		return fmt.Errorf("counld not find sign in page html file in %s", p)
// 	}
// 	// get /sign_in
// 	_, err = r.AddGet(RoutePathSignIn, append([]router.HandlerFunc{parseForm}, signInPage.Handler...)...)
// 	if err != nil {
// 		return err
// 	}
// 	// post /sign_in
// 	_, err = r.AddPost(RoutePathSignIn, parseForm, PostSignIn)
// 	if err != nil {
// 		return err
// 	}
// 	// post /sign_up
// 	_, err = r.AddPost(RoutePathSignUp, parseForm, PostSignUp)
// 	if err != nil {
// 		return err
// 	}
// 	// post /verification_code
// 	_, err = r.AddPost(RoutePathVerificationCode, parseForm, PostVerificationCode)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Parse form before handle, return false if there is a error.
// func parseForm(c *router.Context) bool {
// 	err := c.Req.ParseForm()
// 	if err != nil {
// 		log.Error(err)
// 		return false
// 	}
// 	return true
// }

// func genHSJWT(alg, key string) (string, error) {
// 	header := make(map[string]interface{})
// 	payload := make(map[string]interface{})
// 	switch alg {
// 	case string(jwt.HS384):
// 		return jwt.GenerateHS384(header, payload, []byte(key))
// 	case string(jwt.HS512):
// 		return jwt.GenerateHS512(header, payload, []byte(key))
// 	default:
// 		return jwt.GenerateHS256(header, payload, []byte(key))
// 	}
// }

// func writeJSON(c *router.Context, statusCode int, json []byte) {
// 	c.Res.Header().Set("Content-Type", router.ContentTypeJSON)
// 	c.Res.WriteHeader(statusCode)
// 	c.Res.Write(json)
// }

// func sendEmail(email string) error {
// 	return nil
// }
