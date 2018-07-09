package service

import (
	"github.com/HankWang95/words-gin/app/manager"
	"github.com/HankWang95/words-gin/app/form"
)

func SearchWords4Uid(uid int64) (result *form.WordForm, status int) {
	result, status = manager.SearchWord4Uid(uid)
	return
}

func SearchWord4Word(word string) (result *form.WordForm, status int){
	result, status = manager.SearchWord4Word(word)
	return
}



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



