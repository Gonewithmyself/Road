package db

import "database/sql"

type MySQL struct {
	db *sql.DB
}

func (db *MySQL) Do(sql string, args ...interface{}) (sql.Result, error) {
	stmt, _ := db.db.Prepare(sql)
	return stmt.Exec(args...)
}

func (db *MySQL) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	stmt, _ := db.db.Prepare(sql)
	return stmt.Query(args...)
}

func NewMySQL(dialInfo string) *MySQL {
	d, err := sql.Open("mysql", dialInfo)
	if nil != err {
		panic(err)
	}
	return &MySQL{
		db: d,
	}
}
