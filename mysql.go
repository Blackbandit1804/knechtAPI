package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type MySql struct {
	Dsn string
	DB  *sql.DB
}

func NewMySql(dsn string) (*MySql, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySql{ dsn, db }, nil
}

func (this *MySql) Close() {
	this.DB.Close()
}

func (this *MySql) Query(statement string, values ...interface{}) (*sql.Rows, error) {
	stm, err := this.DB.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	return stm.Query(values...)
}