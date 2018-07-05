package model

type Words struct {
	Word        string `json:"word"`
	Id          int64  `json:"id"`
	Translation string `json:"translation"`
}

type User struct {
	Id       int64
	UserName string
	PassWord string
}
