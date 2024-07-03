package model

import (
	"fmt"
	"strconv"
	"tanaman/db"
	"tanaman/utils"
)

func AddToCart(uuid, product_id, qty, size_id, user_id, date string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	cekExist, _ := dbEngine.QueryString(`SELECT * FROM trans_cart WHERE product_id = ? AND user_id = ? AND size_id=?`, product_id, user_id, size_id)

	if cekExist != nil {
		intqty, _ := strconv.Atoi(cekExist[0]["quantity"])
		qty, _ := strconv.Atoi(qty)
		intqty = intqty + qty
		sqty := fmt.Sprintf("%d", intqty)
		UpdateCart(cekExist[0]["uuid"], sqty, date)
	} else {
		_, err := dbEngine.QueryString(`INSERT INTO trans_cart (uuid, product_id, quantity, size_id, user_id, date_create) VALUES (?,?,?,?,?,?)`, uuid, product_id, qty, size_id, user_id, date)

		if err != nil {
			Respon.Status = 500
			Respon.Message = err.Error()
			return Respon
		}
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}

func UpdateCart(chart_id, qty, date string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	query := fmt.Sprintf("UPDATE trans_cart SET")

	if qty != "" {
		query += fmt.Sprintf(" quantity = %s", qty)
	} else {
		Respon.Status = 404
		Respon.Message = "No data to update"
		return Respon
	}

	query += fmt.Sprintf(", date_update = '%s' WHERE uuid = '%s'", date, chart_id)

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
func GetCart(user_id, chart_id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	var cart string

	if chart_id != "" {
		cart = fmt.Sprintf("AND tc.uuid = %s", user_id)
	}

	result, err := dbEngine.QueryString(`WITH MinImage AS (
						SELECT 
							pi.product_id AS prodid, 
							MIN(img.id) AS min_image_id 
						FROM 
							images AS img 
							INNER JOIN product_image AS PI ON img.uuid = pi.image_id 
						GROUP BY 
							pi.product_id
						) 
							SELECT prod.uuid,prod.product_name,img.file_name,prod.price,prod.discount,tc.quantity,size.size,tc.uuid AS cart_id
						FROM trans_cart AS tc
						LEFT JOIN product AS prod ON prod.uuid = tc.product_id
						LEFT JOIN MinImage AS mi ON mi.prodid = prod.uuid
						LEFT JOIN images AS img ON img.id = mi.min_image_id
						LEFT JOIN ref_category_product AS cat ON cat.id = prod.category_id::integer
						LEFT JOIN ref_size AS size ON size.id = tc.size_id::integer
						WHERE tc.user_id = ?`+cart+` ORDER BY tc.id ASC `, user_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	arrayproduct := make([]interface{}, 0, len(result))
	for i := 0; i < len(result); i++ {
		intPrice, _ := strconv.Atoi(result[i]["price"])
		intDiscount, _ := strconv.Atoi(result[i]["discount"])
		intQTY, _ := strconv.Atoi(result[i]["quantity"])
		intPriceDiscount := intPrice - intDiscount

		product := make(map[string]interface{})
		product["product_id"] = result[i]["uuid"]
		product["cart_id"] = result[i]["cart_id"]
		product["title"] = result[i]["product_name"]
		product["image"] = fmt.Sprintf("%s/file/product/%s", Config.ServerHost, result[i]["file_name"])
		product["price"] = intPrice
		product["img"] = result[i]["file_name"]
		product["price_discount"] = intPriceDiscount
		product["quantity"] = intQTY
		product["size"] = result[i]["size"]
		arrayproduct = append(arrayproduct, product)
	}
	

	Respon.Status = 200
	Respon.Message = "success"
	Respon.Data = arrayproduct
	return Respon
}
func GetCountCart(user_id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	result, err := dbEngine.QueryString(`SELECT COUNT(uuid) AS total FROM trans_cart WHERE user_id = ?`, user_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	Respon.Data = result[0]["total"]
	return Respon
}
