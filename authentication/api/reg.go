package api

import "regexp"

var (
	UsernameRegexp    = ``
	PasswordRegexp    = ``
	PhoneNumberRegexp = ``
	PhoneSMSRegexp    = ``
	usernameRegexp    *regexp.Regexp
	passwordRegexp    *regexp.Regexp
	phoneNumberRegexp *regexp.Regexp
	phoneSMSRegexp    *regexp.Regexp
)

func CompileRegExp() error {
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
	exp, err = regexp.Compile(PhoneSMSRegexp)
	if err != nil {
		return err
	}
	phoneSMSRegexp = exp
	return nil
}
