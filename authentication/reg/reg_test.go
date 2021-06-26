package reg

import "testing"

func test(t *testing.T, r Regexp, ok string, fail []string) {
	if ret := r.Match(ok); ret != nil {
		t.FailNow()
	}
	for _, s := range fail {
		if ret := r.Match(s); ret == nil {
			t.FailNow()
		}
	}
}

func Test_Account(t *testing.T) {
	test(t, Account, "sdf_123", []string{"", "123_abc", "123-abc", "sdf_123+|:'"})
}

func Test_Password(t *testing.T) {
	test(t, Password, "sdf_:'xc13$5", []string{"", "1123"})
}

func Test_PhoneNumber(t *testing.T) {
	test(t, PhoneNumber, "18912341234", []string{"", "1891234123489", "123456", "123abcdcf456", "123.123"})
}

func Test_PhoneVerificationCode(t *testing.T) {
	test(t, PhoneVerificationCode, "123123", []string{"", "12345", "abcdfds", "123.123"})
}

func Test_Email(t *testing.T) {
	test(t, Email, "123@aa.bb", []string{"", "12345", "abcdfds", "123.123", "123@qq", ".abc@qq", ".abc@qq."})
}
