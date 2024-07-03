package model

import (
	"tanaman/db"
	"tanaman/utils"
)

func PostCheckout(uuid, user_id, status, order_number, payment_method_id, notes, resi, sub_total, shipping_cost, discount, total_payment, date_create string, json_alamat []string) utils.Respon {

	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`INSERT INTO trans_order (uuid, user_id, status, order_number, payment_method_id, notes, resi, sub_total, shipping_cost, discount, total_payment, date_create) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`, uuid, user_id, status, order_number, payment_method_id, notes, resi, sub_total, shipping_cost, discount, total_payment, date_create)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	Respon.Status = 200
	Respon.Message = "success"

	return Respon
}
func PostCheckoutDetail(uuid, order_id, size, quantity, product_name, discount_price, image, date_create, product_id string) utils.Respon {

	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`INSERT INTO order_detail (uuid, trans_order_id, size, quantity, product_name, discount_price, image, date_create, product_id) VALUES (?,?,?,?,?,?,?,?,?)`, uuid, order_id, size, quantity, product_name, discount_price, image, date_create, product_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	Respon.Status = 200
	Respon.Message = "success"

	return Respon
}
func GetShippingCost(idshipping string) string {
	dbEngine := db.ConnectDB()
	shipping_cost, err := dbEngine.QueryString(`SELECT cost FROM ref_shipping WHERE id = ?`, idshipping)
	if err != nil {
		return "0"
	}
	return shipping_cost[0]["cost"]
}
