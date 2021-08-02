package db

import (
	"database/sql"
)

var (
	stmtUserInsertPhonePassword        *sql.Stmt
	stmtUserInsertAccountPasswordEmail *sql.Stmt
	stmtUserSelectByAccountOrEmail     *sql.Stmt
	stmtUserSelectIdByPhone            *sql.Stmt
)

func initUserStmt(db *sql.DB) error {
	var err error
	stmtUserInsertPhonePassword, err = db.Prepare(`insert into user(password,phone) values(?,?)`)
	if err != nil {
		return err
	}
	stmtUserInsertAccountPasswordEmail, err = db.Prepare(`insert into user(account,password,email) values(?,?,?)`)
	if err != nil {
		return err
	}
	stmtUserSelectIdByPhone, err = db.Prepare(`select id from user where phone=?`)
	if err != nil {
		return err
	}
	stmtUserSelectByAccountOrEmail, err = db.Prepare(`select id,account,password,phone,name from user where account=? or email=?`)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Id       int64
	Account  sql.NullString
	Password sql.NullString
	Phone    sql.NullString
	Email    sql.NullString
	Name     sql.NullString
}

func (u *User) InsertPhonePassword() (sql.Result, error) {
	return stmtUserInsertPhonePassword.Exec(
		u.Password.String,
		u.Phone.String,
	)
}

func (u *User) InsertAccountPasswordEmail() (sql.Result, error) {
	return stmtUserInsertAccountPasswordEmail.Exec(
		u.Account.String,
		u.Password.String,
		u.Email.String,
	)
}

func (u *User) SelectIdByPhone() error {
	return stmtUserSelectIdByPhone.QueryRow(u.Phone.String).Scan(&u.Id)
}

func (u *User) SelectIdByAccountOrEmail() error {
	return stmtUserSelectByAccountOrEmail.QueryRow(u.Account.String, u.Email.String).Scan(&u.Id)
}
