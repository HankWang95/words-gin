package form

import (
	"time"
	"github.com/HankWang95/words-gin/app/model"
)

type WordForm struct {
	Word          string    `form:"word" sql:"word" json:"word"`
	Translation   string    `form:"translation" sql:"translation" json:"translation"`
	AddedDate     time.Time    `form:"added_date" sql:"added_date" json:"added_date"`
	LastStudyDate time.Time `form:"last_study_date" sql:"last_study_date" json:"last_study_date"`
}


type HotelUserForm struct {
	model.Hotel
	model.HotelUser
}

