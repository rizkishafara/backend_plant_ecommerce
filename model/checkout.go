package model

import (
	"tanaman/db"
	"tanaman/utils"
)

func Checkout(user_id, order_number, payment_id, json_alamat, notes, resi, sub_total, discount, idshipping, date_create string) utils.Respon {

	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(``)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	Respon.Status = 200
	Respon.Message = "success"

	return Respon
}
