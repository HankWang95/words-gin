package manager

import (
	"fmt"
	"github.com/HankWang95/words-gin/app/db"
	"github.com/HankWang95/words-gin/app/model"
	"github.com/smartwalle/dbs"
	"github.com/smartwalle/time4go"
	"github.com/HankWang95/words-gin/app/form"
)

const (
	ADD_WORD_SUCCESS    = 1001
	WORD_IS_EXIST       = 1002
	SEARCH_WORD_SUCCESS = 1011
	SEARCH_WORD_FAIL    = 1012
)

func AddWord(Word, Translation string) int {
	var db = db.GetDBSession()
	var W *model.Words
	//rows, err := db.Query("SELECT uid FROM words WHERE word=? limit 1", Word)
	sb := dbs.NewSelectBuilder()
	sb.Selects("uid")
	sb.From("words")
	sb.Limit(1)
	sb.Where("word = ?", Word)
	fmt.Println(sb.ToSQL())
	//sb.Scan(db, &W)

	var tx = dbs.MustTx(db)

	if err := sb.ScanTx(tx, &W); err != nil {
		return SEARCH_WORD_FAIL
	}

	if W != nil {
		tx.Rollback()
		return WORD_IS_EXIST
	}

	ib := dbs.NewInsertBuilder()
	ib.Table("words")
	ib.SET("word", Word)
	ib.SET("translation", Translation)

	fmt.Println(ib.ToSQL())
	res, err := ib.ExecTx(tx)
	wordId,err := res.LastInsertId()
	if err != nil{
		tx.Rollback()
	}
	if err != nil {
		return SEARCH_WORD_FAIL
	}
	ib = dbs.NewInsertBuilder()
	ib.Table("date")
	ib.SET("added_date",time4go.Now())
	ib.SET("last_study_date",time4go.Now())
	ib.SET("word_id",wordId)
	fmt.Println(ib.ToSQL())

	if _, err := ib.ExecTx(tx); err != nil {
		return SEARCH_WORD_FAIL
	}
	tx.Commit()
	return ADD_WORD_SUCCESS

	//stmt, err := db.Prepare("INSERT words SET word=?,translation=?")
	//if err != nil {
	//	fmt.Println("SQL语句错误")
	//}
	////defer stmt.Close()
	//stmt.Exec(Word, Translation)
	//for rows.Next(){
	//	rows.Scan(&uid)
	//}
	//if err != nil {
	//	fmt.Println(err)

}

func SearchWord4Uid(uid int64) (*form.WordForm, int) {
	//var db = db.GetDBSession()
	//var Word = form.WordForm{}
	//err := db.QueryRow("SELECT word,translation FROM words WHERE uid=?", uid).Scan(&Word.Word, &Word.Translation)
	//if err != nil {
	//	fmt.Println(err)
	//	return &Word, SEARCH_WORD_FAIL
	//}
	var db = db.GetDBSession()
	var WordForm *form.WordForm

	//var sb = dbs.NewSelectBuilder()
	//sb.Selects("rsp.id", "rsp.room_id", "rsp.begin_date", "rsp.end_date", "rsp.price", "rsp.remark", "rsp.status", "rsp.created_on", "rsp.updated_on")
	//sb.Selects("r.name AS room_name", "r.code AS room_code")
	//sb.From("wwwwww", "AS rsp")
	//sb.LeftJoin("dddddd", "AS r ON r.id = rsp.room_id")
	//sb.Where("r.hotel_id = ?", 1)
	//fmt.Println(sb.ToSQL())

	sb := dbs.NewSelectBuilder()
	sb.Selects("w.word", "w.translation")
	sb.Selects("d.added_date", "d.last_study_date")
	sb.From("words", " AS w")
	sb.LeftJoin("date"," AS d ON w.uid = d.word_id")
	sb.Where("w.uid = ?", uid)
	fmt.Println(sb.ToSQL())

	err := sb.Scan(db, &WordForm)

	if err != nil{
		fmt.Println(err)
		return nil, SEARCH_WORD_FAIL
	}

	if WordForm == nil {
		return nil, SEARCH_WORD_FAIL
	}
	fmt.Println(WordForm)

	return WordForm, SEARCH_WORD_SUCCESS
}

func SearchWordsList() ([]*form.WordForm, error) {

	var db = db.GetDBSession()
	var WordList []*form.WordForm
	//var Word model.Words

	sb := dbs.NewSelectBuilder()
	sb.Selects("w.word", "w.translation")
	sb.Selects("d.added_date", "d.last_study_date")
	sb.From("words", " AS w")
	sb.LeftJoin("date"," AS d ON w.uid = d.word_id")

	fmt.Println(sb.ToSQL())
	err := sb.Scan(db, &WordList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//rows,_ := sb.Query(db)
	//fmt.Println(rows,*rows)
	//
	//for rows.Next(){
	//	rows.Scan(&Word.Word, &Word.Translation)
	//	WordList = append(WordList, Word)
	//}
	//
	//	fmt.Println(Word)
	//	fmt.Println(&Word)

	//}
	return WordList, nil

	//var db = db.GetDBSession()
	//var WordsList = make([]form.WordForm, 0)
	//var Word = form.WordForm{}
	//
	//rows, err := db.Query("SELECT word,translation FROM words ")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	//defer rows.Close()
	//
	//for i := 0; rows.Next(); i++ {
	//	if err := rows.Scan(&Word.Word, &Word.Translation); err != nil {
	//		return nil
	//	}
	//	WordsList = append(WordsList, Word)
	//}
	//return &WordsList

}
