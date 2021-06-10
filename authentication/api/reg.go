package api

import "regexp"

var (
	UsernameRegexp    = ``
	PasswordRegexp    = ``
	PhoneNumberRegexp = ``
	PhoneCodeRegexp   = ``
	usernameRegexp    *regexp.Regexp
	passwordRegexp    *regexp.Regexp
	phoneNumberRegexp *regexp.Regexp
	phoneCodeRegexp   *regexp.Regexp
)

func InitRegExp() error {
	exp, err := regexp.Compile(UsernameRegexp)
	if err != nil {
		return err
	}
	usernameRegexp = exp
	exp, err = regexp.Compile(PasswordRegexp)
	if err != nil {
		return err
	}
	passwordRegexp = exp
	exp, err = regexp.Compile(PhoneNumberRegexp)
	if err != nil {
		return err
	}
	phoneNumberRegexp = exp
	exp, err = regexp.Compile(PhoneCodeRegexp)
	if err != nil {
		return err
	}
	phoneCodeRegexp = exp
	return nil
}
