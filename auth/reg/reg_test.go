package reg

import (
	"regexp"
	"testing"
)

func init() {
	Init(&Config{})
}

func test(t *testing.T, r *regexp.Regexp, ok string, fail []string) {
	if !r.MatchString(ok) {
		t.FailNow()
	}
	for _, s := range fail {
		if r.MatchString(s) {
			t.FailNow()
		}
	}
}

func Test_Password(t *testing.T) {
	test(t, Password, "sdf_:'xc13$5", []string{"", "1123"})
}

func Test_Phone(t *testing.T) {
	test(t, Phone, "+86-18912341234", []string{"", "18912341234", "+86-123456", "+86-123abcdcf456", "+86-123.123"})
}

func Test_VerificationCode(t *testing.T) {
	test(t, VerificationCode, "123123", []string{"", "12345", "abcdfds", "123.123"})
}

func Test_Email(t *testing.T) {
	test(t, Email, "123@aa.bb", []string{"", "12345", "abcdfds", "123.123", "123@qq", ".abc@qq", ".abc@qq."})
}
