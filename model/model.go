package model

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var Db *sql.DB

// func init() {
// 	db, err := sql.Open("mysql", "root:dev2312@tcp(172.17.0.1:3306)/books")
// 	if err != nil {
// 		fmt.Println("打开失败")
// 		log.Fatalln(err.Error())
// 	}
// 	fmt.Println("打开成功")
// 	Db = db
// }

import (
 	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db sqlx.Db
func init(){
 	db, err := sql.Open("mysql", "root:dev2312@tcp(172.17.0.1:3306)/books")
	Db=db
}