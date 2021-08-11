package reg

import "regexp"

var (
	Phone            *regexp.Regexp
	Email            *regexp.Regexp
	VerificationCode *regexp.Regexp
	Password         *regexp.Regexp
)

type Config struct {
	Phone            string `json:"phone,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	VerificationCode string `json:"verificationCode,omitempty"`
}

func Init(c *Config) {
	// Phone
	if c.Phone == "" {
		Phone = regexp.MustCompile(`^\+86-1[3-8]\d{9}$`)
	} else {
		Phone = regexp.MustCompile(c.Phone)
	}
	// Email
	if c.Email == "" {
		Email = regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+$`)
	} else {
		Email = regexp.MustCompile(c.Email)
	}
	// VerificationCode
	if c.VerificationCode == "" {
		VerificationCode = regexp.MustCompile(`^\d{6}$`)
	} else {
		VerificationCode = regexp.MustCompile(c.VerificationCode)
	}
	// Password
	if c.VerificationCode == "" {
		Password = regexp.MustCompile(`^.{6,}$`)
	} else {
		Password = regexp.MustCompile(c.VerificationCode)
	}
}
