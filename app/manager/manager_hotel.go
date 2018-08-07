package manager

import (
	"fmt"
	"github.com/HankWang95/words-gin/app/db"
	"github.com/HankWang95/words-gin/app/form"
	"github.com/smartwalle/dbs"
)

func SearchHotelUser() ([]*form.HotelUserForm, error) {
	var res []*form.HotelUserForm

	tx := dbs.MustTx(db.NewDBLink())

	sb := dbs.NewSelectBuilder()
	sb.From("hotel_user", "AS h")
	sb.Selects("h.hotel_id, h.user_id, h.status, h.created_on, h.updated_on")
	sb.Selects("h_h.name")
	sb.Selects("u.id, u.username, u.first_name, u.last_name")
	sb.LeftJoin("user_user", "AS u ON u.id = h.user_id")
	sb.LeftJoin("hotel_hotel", "AS h_h ON h_h.owner_id = h.user_id")
	fmt.Println(sb.ToSQL())
	if err := sb.ScanTx(tx, &res); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
