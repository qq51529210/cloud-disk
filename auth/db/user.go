package db

import "database/sql"

var (
	stmtUserSelect *sql.Stmt
	stmtUserInsert *sql.Stmt
	stmtUserDelete *sql.Stmt
)

func initUserStmt(db *sql.DB) {
	var err error
	stmtUserSelect, err = db.Prepare(`select id,password,name from user where account=?`)
	if err != nil {
		panic(err)
	}
	stmtUserDelete, err = db.Prepare(`delete from user where id=?`)
	if err != nil {
		panic(err)
	}
	stmtUserInsert, err = db.Prepare(`insert into user(account,password,name) values(?,?,?)`)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id       int64
	Account  string
	Password string
	Name     string
}

func (m *User) Select() error {
	return stmtUserSelect.QueryRow(m.Account).Scan(&m.Id, &m.Password, &m.Name)
}

func (m *User) Delete() (sql.Result, error) {
	return stmtUserDelete.Exec(m.Id)
}

func (m *User) Insert() (sql.Result, error) {
	return stmtUserInsert.Exec(m.Account, m.Password, m.Name)
}
