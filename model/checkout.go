package model

import (
	"strconv"
	"tanaman/db"
	"tanaman/utils"
)

func PostCheckout(uuid, user_id, status, order_number, payment_method_id, notes, resi, sub_total, shipping_cost, discount, total_payment, date_create, json_alamat string) utils.Respon {

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
	data := make(map[string]interface{})
	data["order_id"] = order_id
	Respon.Status = 200
	Respon.Data = data
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
func GetProductOrder(user_id, order_id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	getProductOrder, err := dbEngine.QueryString(`SELECT 
												od.uuid,
												od.product_name,
												od.size,
												od.quantity,
												od.discount_price,
												od.image,
												od.product_id
											FROM order_detail od
											WHERE od trans_order_id = ?`, order_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if getProductOrder == nil {
		Respon.Status = 404
		Respon.Message = "Product not found"
		return Respon
	}

	product := make([]interface{}, 0, len(getProductOrder))
	for _, prod := range getProductOrder {
		intQTY, _ := strconv.Atoi(prod["quantity"])
		intDiscount, _ := strconv.Atoi(prod["discount_price"])
		intSubTotal := intQTY * intDiscount
		prd := make(map[string]interface{})
		prd["uuid"] = prod["uuid"]
		prd["product_name"] = prod["product_name"]
		prd["size"] = prod["size"]
		prd["quantity"] = intQTY
		prd["discount_price"] = intDiscount
		prd["subtot_price"] = intSubTotal
		prd["image"] = prod["image"]
		prd["product_id"] = prod["product_id"]
		product = append(product, prd)
	}

	Respon.Status = 200
	Respon.Message = "success"
	Respon.Data = getProductOrder

	return Respon
}
