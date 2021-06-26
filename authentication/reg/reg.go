package reg

import (
	"encoding/json"
	"fmt"
	"regexp"
)

var (
	Account = &_Regexp{
		name: "account",
		expr: `^[a-zA-Z]\w{5,31}$`,
	}
	Password = &_Regexp{
		name: "password",
		expr: `^\S{6,}$`,
	}
	PhoneNumber = &_Regexp{
		name: "phone number",
		expr: `^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`,
	}
	PhoneVerificationCode = &_Regexp{
		name: "phone verification code",
		expr: `^\d{6}$`,
	}
	Email = &_Regexp{
		name: "email",
		expr: `[\w]+(\.[\w]+)*@[\w]+(\.[\w])+`,
	}
)

func init() {
	err := Account.Compile("")
	if err != nil {
		panic(err)
	}
	Password.Compile("")
	PhoneNumber.Compile("")
	PhoneVerificationCode.Compile("")
	Email.Compile("")
}

type Regexp interface {
	Compile(string) error
	Match(s string) []byte
}

type _Regexp struct {
	*regexp.Regexp
	name string
	expr string
	json []byte
}

func (r *_Regexp) Compile(s string) (err error) {
	if s == "" {
		s = r.expr
	}
	r.Regexp, err = regexp.Compile(s)
	if err != nil {
		return err
	}
	r.json, err = json.Marshal(map[string]interface{}{
		"error":  fmt.Sprintf("Invalid %s format.", r.name),
		"regexp": r.Regexp.String(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *_Regexp) Match(s string) []byte {
	if r.Regexp.MatchString(s) {
		return nil
	}
	return r.json
}
