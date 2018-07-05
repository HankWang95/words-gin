package service

import (
	"github.com/HankWang95/words-gin/app/form"
	"github.com/HankWang95/words-gin/app/manager"
)

func SearchWords4Uid(uid int64) (result *form.WordForm, status int) {
	result, status = manager.SearchWord4Uid(uid)
	return
}

func AddWord(Word, Translation string) int {
	status := manager.AddWord(Word, Translation)
	return status
}

func SearchWordsList() (result *[]form.WordForm) {
	result = manager.SearchWordsList()
	return
}
