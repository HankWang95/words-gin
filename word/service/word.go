package service

import (
	"github.com/HankWang95/words-gin/word"
	"github.com/smartwalle/errors"
)

var (
	WordNotExist = errors.New("1000", "单词不存在")
	TypeInvalid  = errors.New("2000", "类型不合法")
)

type WordRepository interface {
	SearchWordWithEnglish(s string) (*word.Word, error)
	SearchWordWithChinese(s string) (*word.Word, error)
	AddWord(word *word.Word) error
}

type wordServeice struct {
	repo WordRepository
}

func NewWordService(repo WordRepository) word.WordService {
	return &wordServeice{repo: repo}
}

func (this *wordServeice) SearchWordWithEnglish(word string) (result *word.Word, err error) {
	result, err = this.repo.SearchWordWithEnglish(word)
	if result == nil {
		newWord, err := searchWordFromYouDao(word)
		if err != nil {
			return nil, err
		}
		if newWord != nil {
			err = addWord(newWord)
			if err != nil {
				return
			}
		}
	}
	return
}

func (this *wordServeice) SearchWordWithChinese(word string) (*word.Word, error) {
	return this.repo.SearchWordWithChinese(word)
}

func addWord(word *word.Word) error {

	return nil
}

func searchWordFromYouDao(word string) (result *word.Word, err error) {
	// todo: youdao API
	return
}
