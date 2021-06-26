package db

import (
	"database/sql"
)

var (
	stmtUserInsert                 *sql.Stmt
	stmtUserSelectByAccountOrPhone *sql.Stmt
	stmtUserSelectIdByPhone        *sql.Stmt
	stmtUserDeleteByAccount        *sql.Stmt
)

func initUserStmt(db *sql.DB) error {
	var err error
	stmtUserInsert, err = db.Prepare(`insert into user(account,password,phone,name) values(?,?,?,?)`)
	if err != nil {
		return err
	}
	stmtUserSelectByAccountOrPhone, err = db.Prepare(`select id,account,password,phone,name from user where account=? or phone=?`)
	if err != nil {
		return err
	}
	stmtUserSelectIdByPhone, err = db.Prepare(`select id from user where phone=?`)
	if err != nil {
		return err
	}
	stmtUserDeleteByAccount, err = db.Prepare(`delete from user where account=?`)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Id       int64
	Account  sql.NullString
	Password sql.NullString
	Phone    string
	Name     sql.NullString
}

func (u *User) Insert() (sql.Result, error) {
	return stmtUserInsert.Exec(
		u.Account,
		u.Password,
		u.Phone,
		u.Name,
	)
}

func (u *User) SelectByAccountOrPhone() error {
	return stmtUserSelectByAccountOrPhone.QueryRow(u.Account, u.Phone).Scan(
		&u.Id,
		&u.Account,
		&u.Password,
		&u.Phone,
		&u.Name,
	)
}

func (u *User) SelectIdByPhone() error {
	return stmtUserSelectIdByPhone.QueryRow(u.Phone).Scan(&u.Id)
}

func (u *User) DeleteByAccount() (sql.Result, error) {
	return stmtUserDeleteByAccount.Exec(u.Account)
}
