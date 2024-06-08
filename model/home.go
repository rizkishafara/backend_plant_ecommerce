package model

import (
	"fmt"
	"strconv"
	"tanaman/db"
	"tanaman/utils"
)

func GetBanner() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon
	getbanner, err := dbEngine.QueryString(`SELECT banner.uuid,banner.title,images.file_name,banner.product_id FROM banner LEFT JOIN images ON images.uuid=banner.image_id`)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	arraybanner := make([]interface{}, 0, len(getbanner))
	for i := 0; i < len(getbanner); i++ {
		banner := make(map[string]interface{})

		banner["id"] = getbanner[i]["uuid"]
		banner["title"] = getbanner[i]["title"]
		banner["images"] = fmt.Sprintf("%s/banner/%s", Config.ServerHost, getbanner[i]["file_name"])
		banner["product_id"] = getbanner[i]["product_id"]
		arraybanner = append(arraybanner, banner)
	}
	Respon.Status = 200
	Respon.Data = arraybanner
	Respon.Message = "success"
	return Respon
}
func GetPopularCategory() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon
	getcategory, err := dbEngine.QueryString(`SELECT ref_category_product.id,ref_category_product.category_name,ref_category_product.description,images.file_name FROM ref_category_product LEFT JOIN images ON images.uuid=ref_category_product.image_id ORDER BY ref_category_product.id LIMIT 3`)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	arraycategory := make([]interface{}, 0, len(getcategory))
	for i := 0; i < len(getcategory); i++ {
		category := make(map[string]interface{})

		category["id"] = getcategory[i]["id"]
		category["category_name"] = getcategory[i]["category_name"]
		category["description"] = getcategory[i]["description"]
		category["image"] = fmt.Sprintf("%s/category/%s", Config.ServerHost, getcategory[i]["file_name"])
		arraycategory = append(arraycategory, category)
	}
	Respon.Status = 200
	Respon.Data = arraycategory
	Respon.Message = "success"
	return Respon
}
func GetNewestProduct() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon
	getproduct, err := dbEngine.QueryString(`WITH MinImage AS (
												SELECT 
													pi.product_id AS prodid,
													MIN(img.id) AS min_image_id
												FROM 
													images AS img
												INNER JOIN 
													product_image AS PI 
												ON 
													img.uuid = pi.image_id
												GROUP BY 
													pi.product_id
											)
											SELECT 
												prod.uuid,
												prod.product_name,
												img.file_name,
												prod.price,
												prod.discount,
												prod.date_created,
												cat.category_name
											FROM
												product AS prod
											LEFT JOIN 
												MinImage AS mi 
											ON 
												mi.prodid = prod.uuid
											LEFT JOIN 
												images AS img 
											ON 
												img.id = mi.min_image_id
											LEFT JOIN
												ref_category_product AS cat
											ON 
												cat.id=prod.category_id
											ORDER BY 
												prod.date_created DESC;
`)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	arrayproduct := make([]interface{}, 0, len(getproduct))
	for i := 0; i < len(getproduct); i++ {
		product := make(map[string]interface{})
		
		intPrice, _ := strconv.Atoi(getproduct[i]["price"])
		intDiscount, _ := strconv.Atoi(getproduct[i]["discount"])
		intPriceDiscount := intPrice - intDiscount

		product["id"] = getproduct[i]["uuid"]
		product["title"] = getproduct[i]["product_name"]
		product["image"] = fmt.Sprintf("%s/product/%s", Config.ServerHost, getproduct[i]["file_name"])
		product["price"] = intPrice
		product["price_discount"] = intPriceDiscount
		product["category"] = getproduct[i]["category_name"]
		arrayproduct = append(arrayproduct, product)
	}
	Respon.Status = 200
	Respon.Data = arrayproduct
	Respon.Message = "success"
	return Respon
}
