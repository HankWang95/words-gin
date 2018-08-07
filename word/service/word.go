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
	SearchWord4English(s string) (*word.Word, error)
	SearchWord4Chinese(s string) (*word.Word, error)
	SearchWord4Id(s int) (*word.Word, error)
}

type wordServeice struct {
	repo WordRepository
}

func NewWordService(repo WordRepository) word.WordService {
	return &wordServeice{repo: repo}
}

func (this *wordServeice) SearchWord(s interface{}) (*word.Word, error) {
	switch s.(type) {
	case string:
		return this.repo.SearchWord4English(s.(string))

	case word.Chinese:
		return this.repo.SearchWord4Chinese(s.(string))

	case int:
		return this.repo.SearchWord4Id(s.(int))
	default:
		return nil, TypeInvalid
	}
}
