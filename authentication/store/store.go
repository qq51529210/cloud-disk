package store

import (
	"database/sql"
)

var (
	_db *sql.DB
	stmtUserInsert *sql.Stmt
	stmtUserDelete *sql.Stmt
	stmtUserUpdate *sql.Stmt
	stmtUserSelect *sql.Stmt
	stmtUserSelectUserPage *sql.Stmt
	stmtSelectUserCount *sql.Stmt
	stmtUserSelectByUsername *sql.Stmt
)

func Init(db *sql.DB) {
	var err error
	stmtUserInsert, err = db.Prepare("insert into `user` (`name`,`pwd`) values (?,?)")
	if err != nil {
		panic(err)
	}
	stmtUserDelete, err = db.Prepare("delete from `user` where `id`=?")
	if err != nil {
		panic(err)
	}
	stmtUserUpdate, err = db.Prepare("update `user` set `name`=?,`pwd`=? where `id`=?")
	if err != nil {
		panic(err)
	}
	stmtUserSelect, err = db.Prepare("select * from `user` where `id`=?")
	if err != nil {
		panic(err)
	}
	stmtUserSelectUserPage, err = db.Prepare("select * from `user` limit ?,?")
	if err != nil {
		panic(err)
	}
	stmtSelectUserCount, err = db.Prepare("select count(*) from `user`")
	if err != nil {
		panic(err)
	}
	stmtUserSelectByUsername, err = db.Prepare("select `id`,`pwd` from `tb_user` where `name`=?")
	if err != nil {
		panic(err)
	}
	_db = db
}
