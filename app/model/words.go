package model

import "time"

type Words struct {
	Word        string `json:"word" sql:"word"`
	Id          int64  `json:"id" sql:"uid"`
	Translation
	Date

}

type Translation struct {
	V string
	N string

}

type User struct {
	Id       int64
	UserName string
	PassWord string
}

type Date struct {
	AddedDate     time.Time `json:"added_date" sql:"added_date"`
	LastStudyDate time.Time `json:"last_study_date" sql:"last_study_date"`
}




