package model

import (
	"fmt"
	"strconv"
	"tanaman/db"
	"tanaman/utils"
)

func GetListHistoryOrder(user_id string) utils.Respon {
	var Respon utils.Respon
	dbEngine := db.ConnectDB()
	result, err := dbEngine.QueryString(`SELECT 
											t.uuid AS order_id, 
											t.status, 
											t.order_number,
											t.total_payment,
											d.uuid AS detail_order_id, 
											d.size,
											d.quantity,
											d.product_name,
											d.discount_price,
											d.image,
											d.price
										FROM 
											trans_order t
										JOIN 
											detail_order d ON t.trans_order_id = d.trans_order_id
										WHERE 
											d.detail_order_id = (
												SELECT 
													MIN(detail_order_id) 
												FROM 
													detail_order 
												WHERE 
													trans_order_id = t.trans_order_id
											); `)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	history := make([]map[string]interface{}, 0, len(result))
	for i := 0; i < len(result); i++ {
		order := make(map[string]interface{})

		intStatus, _ := strconv.Atoi(result[i]["status"])
		intTotalPayment, _ := strconv.Atoi(result[i]["total_payment"])
		intQTY, _ := strconv.Atoi(result[i]["quantity"])
		intPrice, _ := strconv.Atoi(result[i]["price"])
		intDiscount, _ := strconv.Atoi(result[i]["discount_price"])

		order["order_id"] = result[i]["order_id"]
		order["status"] = intStatus
		order["order_number"] = result[i]["order_number"]
		order["total_payment"] = intTotalPayment
		order["detail_order_id"] = result[i]["detail_order_id"]
		order["size"] = result[i]["size"]
		order["quantity"] = intQTY
		order["product_name"] = result[i]["product_name"]
		order["discount_price"] = intDiscount
		order["image"] = fmt.Sprintf("%s/file/product/%s", Config.ServerHost, result[i]["image"])
		order["price"] = intPrice
		history = append(history, order)
	}

	Respon.Status = 200
	Respon.Data = history
	Respon.Message = "success"
	return Respon
}
