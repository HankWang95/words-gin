package mysql

import (
	"github.com/HankWang95/words-gin/word"
	"github.com/HankWang95/words-gin/word/service"
	"github.com/smartwalle/dbs"
)

const (
	k_DB_WORD = "words"
)

type wordRepository struct {
	db dbs.DB
}

func NewWordRepository(db dbs.DB) service.WordRepository {
	return &wordRepository{db: db}
}

func (this *wordRepository) SearchWordWithEnglish(s string) (result *word.Word, err error) {
	var sb = dbs.NewSelectBuilder()
	sb.Selects("w.id", "w.english", "w.chinese")
	sb.From(k_DB_WORD, "AS w")
	sb.Where("w.english = ?", s)
	err = sb.Scan(this.db, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (this *wordRepository) SearchWordWithChinese(s string) (result *word.Word, err error) {
	var sb = dbs.NewSelectBuilder()
	sb.Selects("w.id", "w.english", "w.chinese")
	sb.From(k_DB_WORD, "AS w")
	sb.Where("w.chinese = ?", s)
	err = sb.Scan(this.db, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (this *wordRepository) SearchWordWithId(s int) (result *word.Word, err error) {
	var sb = dbs.NewSelectBuilder()
	sb.Selects("w.id", "w.english", "w.chinese")
	sb.From(k_DB_WORD, "AS w")
	sb.Where("w.id = ?", s)
	err = sb.Scan(this.db, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (this *wordRepository) AddWord(word *word.Word) error {
	return nil
}
