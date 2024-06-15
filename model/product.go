package model

import (
	"fmt"
	"strconv"
	"tanaman/db"
	"tanaman/utils"
)

func GetProduct(sizeData, page, search, minPrice, maxPrice, sort string, idCategory, idSize []int) utils.Respon {
	var Respon utils.Respon

	stringWhere := "WHERE prod.uuid != ''"

	//Search
	if search != "" {
		stringWhere += fmt.Sprintf(" AND (prod.product_name LIKE '%%%s%%' OR prod.description LIKE '%%%s%%' OR cat.category_name LIKE '%%%s%%')", search, search, search)
	}

	//Filter Category
	if len(idCategory) != 0 {
		stringWhere += fmt.Sprintf(" AND prod.category_id::integer  = %d", idCategory[0])
		for i := 1; i < len(idCategory); i++ {
			fmt.Println(idCategory[i])
			stringWhere += fmt.Sprintf(" OR prod.category_id::integer  = %d", idCategory[i])
		}
	}

	//Filter Size
	// if len(idSize) != 0 {
	// 	stringWhere += fmt.Sprintf("AND prod.size_id LIKE (%s)", idSize[0])
	// 	for i := 1; i < len(idSize); i++ {
	// 		stringWhere += fmt.Sprintf(" OR prod.size_id LIKE (%s)", idSize[i])
	// 	}
	// }

	//Filter Price
	if minPrice != "" || maxPrice != "" {
		intMinPrice, _ := strconv.Atoi(minPrice)
		intMaxPrice, _ := strconv.Atoi(maxPrice)
		if intMaxPrice == 0 {
			intMaxPrice = 999999999999999999
		}
		stringWhere += fmt.Sprintf(" AND prod.price::integer  >= %d AND prod.price::integer  <= %d", intMinPrice, intMaxPrice)
	}

	//Sort
	if sort == "a-z" {
		stringWhere += fmt.Sprintf(" ORDER BY prod.product_name ASC")
	} else if sort == "z-a" {
		stringWhere += fmt.Sprintf(" ORDER BY prod.product_name DESC")
	} else if sort == "low-expensive" {
		stringWhere += fmt.Sprintf(" ORDER BY prod.price ASC")
	} else if sort == "expensive-low" {
		stringWhere += fmt.Sprintf(" ORDER BY prod.price DESC")
	} else if sort == "newest" {
		stringWhere += fmt.Sprintf(" ORDER BY prod.date_created DESC")
	} else if sort == "oldest" {
		stringWhere += fmt.Sprintf(" ORDER BY prod.date_created ASC")
	} else {
		stringWhere += fmt.Sprintf(" ORDER BY prod.date_created DESC")
	}

	//Pagination
	if page != "" && sizeData != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			Respon.Status = 400
			Respon.Message = "Invalid page parameter"
			return Respon
		}
		sizeInt, err := strconv.Atoi(sizeData)
		if err != nil {
			Respon.Status = 400
			Respon.Message = "Invalid size parameter"
			return Respon
		}

		offset := (pageInt - 1) * sizeInt
		stringWhere += fmt.Sprintf(" LIMIT %d OFFSET %d", sizeInt, offset)
	}

	query := fmt.Sprintf(`WITH MinImage AS (
						SELECT 
							pi.product_id AS prodid, 
							MIN(img.id) AS min_image_id 
						FROM 
							images AS img 
							INNER JOIN product_image AS PI ON img.uuid = pi.image_id 
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
						LEFT JOIN MinImage AS mi ON mi.prodid = prod.uuid 
						LEFT JOIN images AS img ON img.id = mi.min_image_id 
						LEFT JOIN ref_category_product AS cat ON cat.id = prod.category_id::integer
						%s`, stringWhere)

	dbEngine := db.ConnectDB()

	countProduct, err := dbEngine.QueryString(`SELECT COUNT(id) AS count FROM product`)
	intCountProduct, _ := strconv.Atoi(countProduct[0]["count"])
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	getproduct, err := dbEngine.QueryString(query)
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
		product["image"] = fmt.Sprintf("%s/file/product/%s", Config.ServerHost, getproduct[i]["file_name"])
		product["price"] = intPrice
		product["price_discount"] = intPriceDiscount
		product["category"] = getproduct[i]["category_name"]
		arrayproduct = append(arrayproduct, product)
	}
	respData := make(map[string]interface{})

	respData["product"] = arrayproduct
	respData["total"] = intCountProduct
	Respon.Status = 200
	Respon.Data = respData
	Respon.Message = "success"
	return Respon
}
func GetProductByID(id string) utils.Respon {
	var Respon utils.Respon

	query := fmt.Sprintf(`SELECT 
							prod.uuid, 
							prod.product_name, 
							prod.description, 
							prod.price, 
							prod.discount, 
							prod.date_created, 
							cat.category_name 
						FROM 
							product AS prod 
						LEFT JOIN ref_category_product AS cat ON cat.id = prod.category_id::integer
						WHERE prod.uuid = '%s'`, id)

	dbEngine := db.ConnectDB()

	getproduct, err := dbEngine.QueryString(query)
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
		product["description"] = getproduct[i]["description"]
		product["price"] = intPrice
		product["price_discount"] = intPriceDiscount
		product["category"] = getproduct[i]["category_name"]
		arrayproduct = append(arrayproduct, product)
	}
	respData := make(map[string]interface{})

	respData["product"] = arrayproduct
	Respon.Status = 200
	Respon.Data = respData
	Respon.Message = "success"
	return Respon
}