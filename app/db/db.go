package app

import (
	_"github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

var db  sql.DB



// use raw SQL
func NewDBLink() error{
	var err error
	&db, err = sql.Open("mysql", "gin:1234@tcp(127.0.0.1:3306)/words?charset=utf8")
	if err != nil{
		fmt.Println("链接数据库失败")
		return err
	}
	return nil
}


func GetDBSession() *sql.DB {
	return &db
}





//
//
//func AddWord(db *sql.DB,word string) error {
//
//	var ib = dbs.NewInsertBuilder()
//	ib.SET("word", word)
//	ib.Table("words")
//
//	res, err := ib.Exec(db)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//

	//fmt.Println(res.LastInsertId())


	//stmt, err := db.Prepare("INSERT words SET word=?")
	//if err != nil {
	//	fmt.Println("SQL语句错误")
	//}
	//
	//res, err := stmt.Exec(word)
	//if err != nil{
	//	fmt.Println("添加单词失败")
	//}
	//res.LastInsertId()
//	return nil
//
//}
//
//type Words struct {
//	Uid  int `sql:"uid"`
//	Word string `sql:"word"`
//}
//
//
//func SearchWord(db *sql.DB, id int) (*Words, error){
//	var sb = dbs.NewSelectBuilder()
//	sb.Selects("uid", "word")
//	sb.From("words")
//	sb.Where("uid = ?", id)
//
//	var result *Words
//	err := sb.Scan(db, &result)
//	return result, err

	//var word string
	//err := db.QueryRow("SELECT word FROM words WHERE uid=?", id).Scan(&word)
	//return word, err


