package model

import (
	"fmt"
)

type Article struct {
	Id      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
}

// //查询一条数据
// func ArticleGet(id int64) (Article, error) {
// 	mod := Article{}
// 	row := Db.QueryRow("select id,title,author,content from article where id=? limit 1", id)
// 	if err := row.Scan(&mod.Id, &mod.Title, &mod.Author, &mod.Content); err != nil {
// 		fmt.Printf("查询失败%v\n", err)
// 		return mod, err
// 	}
// 	return mod, nil
// }

// //查询多条数据
// func ArticleList() ([]Article, error) {
// 	mods := make([]Article, 0, 10)
// 	rows, err := Db.Query("select id,title,author,content from article")
// 	if err != nil {
// 		fmt.Printf("查询多行失败%v\n", err)
// 		return mods, err
// 	}

// 	defer func() {
// 		if rows != nil {
// 			rows.Close()
// 		}
// 	}()

// 	for rows.Next() {
// 		var i int
// 		err = rows.Scan(&mods[i].Id, &mods[i].Title, &mods[i].Author, &mods[i].Content)
// 		if err != nil {
// 			fmt.Printf("Scan failed, err :%v\n", err)
// 			return mods, err
// 		}
// 		i++
// 		return mods, nil
// 	}
// 		return mods, nil
// }


//查询1条数据用Get
//db.Get(1,2,3...)
//1.查询的结构
//2.sql语句
//3.sql里的参数
func ArticleGet(id int64) (Article, error) {

//查询多条数据Select
//db.Select(1,2,3...)

//非查询的执行
// 1.sql
//2 参数
//db.Exec(1,2...)
func ArticleList() ([]Article, error) {
	var mods  []Article
	err := Db.
}