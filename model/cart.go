package model

import (
	"fmt"
	"tanaman/db"
	"tanaman/utils"
)

func AddToCart(uuid, product_id, qty, size_id, user_id, date string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`INSERT INTO trans_cart (uuid, product_id, quantity, size_id, user_id, date_create) VALUES (?,?,?,?,?,?)`, uuid, product_id, qty, size_id, user_id, date)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func UpdateCart(chart_id, qty, size_id, date string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	query := fmt.Sprintf("UPDATE trans_cart SET")

	if qty != "" && size_id != "" {
		query += fmt.Sprintf("quantity = %s, size_id=%s", qty, size_id)
	} else if qty != "" {
		query += fmt.Sprintf("quantity = %s", qty)
	} else if size_id != "" {
		query += fmt.Sprintf("size_id = %s", size_id)
	} else {
		Respon.Status = 404
		Respon.Message = "No data to update"
		return Respon
	}

	query += fmt.Sprintf(", date_update = %s WHERE uuid = %s", date, chart_id)

	_, err := dbEngine.QueryString(query)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func DeleteCart(chart_id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`DELETE FROM trans_cart WHERE uuid = ?`, chart_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
