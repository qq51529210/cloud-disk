package store

import (
	"database/sql"
	"strings"
)

// User is the model of tb_user
type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Pwd sql.NullString `json:"pwd"`
}

// Insert is the function of SQL: insert into `user` (`name`,`pwd`) values (?,?)
func (m *User) Insert() (sql.Result, error) {
	return stmtUserInsert.Exec(
		m.Name,
		m.Pwd,
	)
}

// Delete is the function of SQL: delete from `user` where `id`=?
func (m *User) Delete() (sql.Result, error) {
	return stmtUserDelete.Exec(
		m.Id,
	)
}

// Update is the function of SQL: update `user` set `name`=?,`pwd`=? where `id`=?
func (m *User) Update() (sql.Result, error) {
	return stmtUserUpdate.Exec(
		m.Name,
		m.Pwd,
		m.Id,
	)
}

// Select is the function of SQL: select * from `user` where `id`=?
func (m *User) Select() error {
	err := stmtUserSelect.QueryRow(
		m.Id,
	).Scan( 
		&m.Id, 
		&m.Name, 
		&m.Pwd,
	)
	if err != nil {
		return err
	}
	return nil
}

// SelectUserPage is the function of SQL: select * from `user` limit ?,?
func SelectUserPage(offset, count interface{}) ([]*User,error) { 
	var models []*User
	rows, err := stmtUserSelectUserPage.Query(
		offset,
		count,
	)
	if err != nil {
		return nil,err
	}
	for rows.Next() { 
		var model User
		err = rows.Scan(
			&model.Id,
			&model.Name,
			&model.Pwd,
		)
		if err != nil {
			return nil,err
		} 
		models = append(models, &model)
	}
	return models,err
}

// SelectUserOrderPage is the function of SQL: select * from `user` order by sort order limit ?,?
func (m *User) SelectUserOrderPage(sort, order string, offset, count interface{}) ([]*User,error) {
	var str strings.Builder
	str.WriteString("select * from `user` order by ")
	str.WriteString(sort)
	str.WriteString(" ")
	str.WriteString(order)
	str.WriteString(" limit ?,?") 
	var models []*User
	rows, err := _db.Query(
		str.String(),
		offset,
		count,
	)
	if err != nil {
		return nil,err
	}
	for rows.Next() { 
		var model User
		err = rows.Scan(
			&model.Id,
			&model.Name,
			&model.Pwd,
		)
		if err != nil {
			return nil,err
		} 
		models = append(models, &model)
	}
	return models,err
}

// SelectUserCount is the function of SQL: select count(*) from `user`
func SelectUserCount() (int64,error) { 
	var value0 int64
	err := stmtSelectUserCount.QueryRow(
	).Scan(
		&value0,
	)
	if err != nil {
		return value0,err
	}
	return value0,nil
}

// SelectByUsername is the function of SQL: select `id`,`pwd` from `tb_user` where `name`=?
func (m *User) SelectByUsername() error {
	err := stmtUserSelectByUsername.QueryRow(
		m.Name,
	).Scan( 
		&m.Id, 
		&m.Pwd,
	)
	if err != nil {
		return err
	}
	return nil
}
