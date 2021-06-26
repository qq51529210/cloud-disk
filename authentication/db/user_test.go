package db

import (
	"testing"
)

func Test_User(t *testing.T) {
	err := InitMysql("root:123123@tcp(127.0.0.1:3306)/anthentication")
	if err != nil {
		t.Fatal(err)
	}
	testAccount := "test account"
	testName := "test name"
	testPassword := "test password"
	testPhone := "test phone"
	// Try delete.
	var user1 User
	user1.Account.String = testAccount
	user1.Account.Valid = true
	_, err = user1.DeleteByAccount()
	if err != nil {
		t.Fatal(err)
	}
	// Insert
	user1.Name.String = testName
	user1.Name.Valid = true
	user1.Password.String = testPassword
	user1.Password.Valid = true
	user1.Phone = testPhone
	_, err = user1.Insert()
	if err != nil {
		t.Fatal(err)
	}
	// Existed
	var user2 User
	user2.Account.String = testAccount
	user2.Account.Valid = true
	user2.Phone = "1"
	_, err = user2.Insert()
	if err == nil {
		t.FailNow()
	} else {
		if !IsExistedError(err) {
			t.Fatal(err)
		}
	}
	user2.Account.Valid = false
	user2.Phone = testPhone
	_, err = user2.Insert()
	if err == nil {
		t.FailNow()
	} else {
		if !IsExistedError(err) {
			t.Fatal(err)
		}
	}
	// Select
	var user3 User
	user3.Account.String = testAccount
	user3.Account.Valid = true
	err = user3.SelectByAccountOrPhone()
	if err != nil {
		t.Fatal(err)
	}
	if user3.Name.String != testName ||
		user3.Password.String != testPassword ||
		user3.Phone != testPhone {
		t.FailNow()
	}
	user3.Account.String = ""
	user3.Account.Valid = false
	user3.Phone = testPhone
	err = user3.SelectByAccountOrPhone()
	if err != nil {
		t.Fatal(err)
	}
	if user3.Account.String != testAccount ||
		user3.Name.String != testName ||
		user3.Password.String != testPassword ||
		user3.Phone != testPhone {
		t.FailNow()
	}
}
