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

func CompileRegExp() {
	usernameRegexp = regexp.MustCompile(UsernameRegexp)
	passwordRegexp = regexp.MustCompile(PasswordRegexp)
	phoneNumberRegexp = regexp.MustCompile(PhoneNumberRegexp)
	phoneSMSRegexp = regexp.MustCompile(PhoneSMSRegexp)
}
