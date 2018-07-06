package service

import (
	"github.com/HankWang95/words-gin/app/manager"
	"github.com/HankWang95/words-gin/app/form"
)

func SearchHotelUser() []*form.HotelUserForm {
	res,err := manager.SearchHotelUser()
	if err != nil {
		return nil
	}
	return res
}
