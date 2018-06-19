package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

// 数据库实例
var (
	DB *sqlx.DB
)

// Init 初始化
func Init() {

	db, err := sqlx.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(3)
	DB = db
}