package manager

import (
	"fmt"
	"github.com/HankWang95/words-gin/app/db"
	"github.com/HankWang95/words-gin/app/form"
)


const(
	ADD_WORD_SUCCESS = 1001
	WORD_IS_EXIST = 1002
	SEARCH_WORD_SUCCESS = 1011
	SEARCH_WORD_FAIL = 1012
)

func AddWord(Word, Translation string) int{
	var db = db.GetDBSession()
	var uid int64
	err := db.QueryRow("SELECT uid FROM words WHERE word=?",Word).Scan(&uid)
	if err != nil {
		fmt.Println(err)

		stmt, err := db.Prepare("INSERT words SET word=?,translation=?")
		if err != nil {
			fmt.Println("SQL语句错误")
		}
		stmt.Exec(Word, Translation)

		return ADD_WORD_SUCCESS
	}else {
		return WORD_IS_EXIST
	}


}

func SearchWord4Uid(uid int64) (*form.WordForm,int) {
	var db  = db.GetDBSession()
	var Word = form.WordForm{}
	err := db.QueryRow("SELECT word,translation FROM words WHERE uid=?",uid).Scan(&Word.Word, &Word.Translation)
	if err != nil{
		fmt.Println(err)
		return &Word,SEARCH_WORD_FAIL
	}
	return &Word,SEARCH_WORD_SUCCESS
}

func SearchWordsList() (*[]*form.WordForm) {
	var db = db.GetDBSession()
	var WordsList  = make([]*form.WordForm,0)
	var Word = form.WordForm{}

	rows, err := db.Query("SELECT word,translation FROM words ")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i:=0; rows.Next();i++ {
		if err := rows.Scan(&Word.Word, &Word.Translation);err != nil{
			return nil
		}
		WordsList = append(WordsList, &Word)
	}
	return &WordsList


}
