package manager

import (
	"fmt"
	"github.com/HankWang95/words-gin/app/db"
	"github.com/HankWang95/words-gin/app/form"
	"github.com/smartwalle/dbs"
	"log"
)

const (
	ADD_WORD_SUCCESS = 1000 + iota
	WORD_IS_NOT_EXIST
)

func AddWord(Word, Translation string) int {
	var db = db.GetDBSession()
	ib := dbs.NewInsertBuilder()
	ib.SET("word", Word)
	ib.SET("translation", Translation)
	_, err := ib.Exec(db)
	if err != nil {
		log.Fatal(err)
	}
	return ADD_WORD_SUCCESS

}

func SearchWord4Uid(uid int64) (*form.WordForm, int) {

	var db = db.GetDBSession()
	var WordForm *form.WordForm

	sb := dbs.NewSelectBuilder()
	sb.Selects("w.word", "w.translation")
	sb.Selects("d.added_date", "d.last_study_date")
	sb.From("words", " AS w")
	sb.LeftJoin("date", " AS d ON w.uid = d.word_id")
	sb.Where("w.uid = ?", uid)
	fmt.Println(sb.ToSQL())

	err := sb.Scan(db, &WordForm)

	if err != nil {
		log.Fatal(err)
		return nil, 0
	}

	if WordForm == nil {
		return nil, WORD_IS_NOT_EXIST
	}

	return WordForm, 0
}

func SearchWord4Word(word string) (*form.WordForm, int) {
	var db = db.GetDBSession()
	var WordForm *form.WordForm

	sb := dbs.NewSelectBuilder()
	sb.Selects("w.word", "w.translation")
	sb.Selects("d.added_date", "d.last_study_date")
	sb.From("words", " AS w")
	sb.LeftJoin("date", " AS d ON w.uid = d.word_id")
	sb.Where("w.word = ? ", word)
	fmt.Println(sb.ToSQL())

	err := sb.Scan(db, &WordForm)

	if err != nil {
		fmt.Println(err)
		return nil, WORD_IS_NOT_EXIST
	}

	if WordForm == nil {
		return nil, WORD_IS_NOT_EXIST
	}

	return WordForm, 0

}

func AllWordsList() ([]*form.WordForm, error) {

	var db = db.GetDBSession()
	var WordList []*form.WordForm
	//var Word model.Words

	sb := dbs.NewSelectBuilder()
	sb.Selects("w.word", "w.translation")
	sb.Selects("d.added_date", "d.last_study_date")
	sb.From("words", " AS w")
	sb.LeftJoin("date", " AS d ON w.uid = d.word_id")

	fmt.Println(sb.ToSQL())
	err := sb.Scan(db, &WordList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return WordList, nil

}

func CreateWordsList() ([]*form.WordForm, error) {

	var db = db.GetDBSession()
	var WordList []*form.WordForm

	sb := dbs.NewSelectBuilder()
	sb.Selects("w.word", "w.translation")
	sb.From("words", " AS w")
	sb.LeftJoin("date", " AS d ON w.uid = d.word_id")
	sb.Where("TO_DAYS(NOW()) - TO_DAYS(added_date) >= ?", 1)

	fmt.Println(sb.ToSQL())
	err := sb.Scan(db, &WordList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return WordList, nil
}

// TODO: 加入到各个manager里
func SetRedisWord(word string) error {
	rSess := db.GetRedisSession()
	defer rSess.Close()
	res := rSess.SET(word, "True", "EX", 1000)
	if res.Error != nil {
		return res.Error
	}
	if res.MustString() == "OK" {
		fmt.Println(res.Data)
		return nil
	}
	return nil

}
