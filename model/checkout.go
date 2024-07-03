package model

import (
	"tanaman/db"
	"tanaman/utils"
)

func Checkout(uuid, user_id,status, order_number, payment_id, notes, resi, sub_total, discount, idshipping, shippingcost, date_create string, json_alamat []string) utils.Respon {

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
