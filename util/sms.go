package util

import "fmt"

func SendSMS(number, code string) error {
	fmt.Println("send", code, "to", number)
	return nil
}
