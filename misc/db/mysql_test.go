package db

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySQL_Query(t *testing.T) {
	dialInfo := `john:since1999@tcp(localhost:3306)/test?charset=utf8`
	h := NewMySQL(dialInfo)
	rows, err := h.Query(`select id, name from user where id = ?`, 1)
	if nil != err {
		t.Error(err)
	}

	for rows.Next() {
		var (
			id   int
			name string
		)
		rows.Scan(&id, &name)

		fmt.Println(id, name)
	}

	t.Error("xx")
}

func TestMySQL_insert(t *testing.T) {
	dialInfo := `john:since1999@tcp(localhost:3306)/test?charset=utf8`
	h := NewMySQL(dialInfo)
	res, err := h.Do(`insert into user set name=?, password=?, role=?, salt=?`,
		"todo", "xxxx", 1, "none")
	if nil != err {
		t.Error(err)
	}
	t.Error(res, err)
}
