package model

import "time"

type Hotel struct {
	HotelId   int64     `json:"hotel_id" sql:"hotel_id"`
	UserId    int64     `json:"user_id" sql:"user_id"`
	Status    int       `json:"status" sql:"status"`
	CreatedOn time.Time `json:"created_on" sql:"created_on"`
	UpdatedOn time.Time `json:"updated_on" sql:"updated_on"`
	HotelName string `json:"hotel_name" sql:"name"`
}

type HotelUser struct {
	Id        int64  `json:"id" sql:"id"`
	UserName  string `json:"username" sql:"username"`
	FirstName string `json:"first_name" sql:"first_name"`
	LastName  string `json:"last_name" sql:"last_name"`

}

