package service

import (
	"github.com/HankWang95/words-gin/app/form"
	"github.com/HankWang95/words-gin/app/manager"
)

func SearchWords4Uid(uid int64) (result *form.WordForm, status int) {
	result, status = manager.SearchWord4Uid(uid)
	return
}

func SearchWord4Word(word string) (result *form.WordForm, status int) {
	result, status = manager.SearchWord4Word(word)
	if status == manager.WORD_IS_NOT_EXIST || result == nil {
		status = AddWord(word, "等待翻译api接入")

	}
	return
}

//翻译api接口方法
//func TranslateWord(word string) *model.Translation {
//
//}

func AddWord(Word, Translation string) int {
	status := manager.AddWord(Word, Translation)
	return status
}

func AllWordsList() (result []*form.WordForm) {
	result, _ = manager.AllWordsList()
	return
}

func CreateWordsList() (result []*form.WordForm) {
	result, _ = manager.CreateWordsList()
	return
}

func SetRedisWord(word string) error {
	return manager.SetRedisWord(word)
}
