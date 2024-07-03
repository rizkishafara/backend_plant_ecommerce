package model

import (
	"fmt"
	"strconv"
	"strings"
	"tanaman/db"
	"tanaman/utils"
)

func GetProduct(sizeData, page, search, minPrice, maxPrice, sort string, idCategory []int, idSize string) utils.Respon {
	var Respon utils.Respon

	stringWhere := "WHERE prod.uuid != ''"

	//Search
	if search != "" {
		stringWhere += fmt.Sprintf(" AND (prod.product_name ILIKE '%%%s%%' OR prod.description ILIKE '%%%s%%' OR cat.category_name ILIKE '%%%s%%')", search, search, search)
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
	if len(idSize) != 0 {

		idSize = strings.Replace(idSize, "[", "", -1)
		idSize = strings.Replace(idSize, "]", "", -1)

		stringWhere += fmt.Sprintf(" AND size.id IN (%s)", idSize)

	}

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
	var stringSort string
	if sort == "a-z" {
		stringSort += fmt.Sprintf(" ORDER BY prod.product_name ASC")
	} else if sort == "z-a" {
		stringSort += fmt.Sprintf(" ORDER BY prod.product_name DESC")
	} else if sort == "low-expensive" {
		stringSort += fmt.Sprintf(" ORDER BY prod.price ASC")
	} else if sort == "expensive-low" {
		stringSort += fmt.Sprintf(" ORDER BY prod.price DESC")
	} else if sort == "newest" {
		stringSort += fmt.Sprintf(" ORDER BY prod.date_created DESC")
	} else if sort == "oldest" {
		stringSort += fmt.Sprintf(" ORDER BY prod.date_created ASC")
	} else {
		stringSort += fmt.Sprintf(" ORDER BY prod.date_created DESC")
	}

	//Pagination
	var stringLimit string
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
		stringLimit += fmt.Sprintf(" LIMIT %d OFFSET %d", sizeInt, offset)
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
							cat.category_name,
							ARRAY_AGG(DISTINCT size.size) AS sizes
						FROM 
							product AS prod 
						LEFT JOIN MinImage AS mi ON mi.prodid = prod.uuid 
						LEFT JOIN images AS img ON img.id = mi.min_image_id 
						LEFT JOIN ref_category_product AS cat ON cat.id = prod.category_id::integer
						LEFT JOIN product_size AS ps ON ps.product_id = prod.uuid
    					LEFT JOIN ref_size AS size ON size.id = ps.size_id::integer
						%s 
						GROUP BY
    						prod.uuid, prod.product_name, img.file_name, prod.price, prod.discount, prod.date_created, cat.category_name
						%s%s`, stringWhere, stringSort, stringLimit)

	dbEngine := db.ConnectDB()

	// countProduct, err := dbEngine.QueryString(`SELECT COUNT(id) AS count FROM product`)
	queryCount := fmt.Sprintf(`SELECT 
									prod.uuid AS id,
									ARRAY_AGG(DISTINCT size.size) AS sizes
								FROM 
									product AS prod 
								LEFT JOIN ref_category_product AS cat ON cat.id = prod.category_id::integer
								LEFT JOIN product_size AS ps ON ps.product_id = prod.uuid
    							LEFT JOIN ref_size AS size ON size.id = ps.size_id::integer
								%s
								GROUP BY prod.uuid`, stringWhere)
	countProduct, err := dbEngine.QueryString(queryCount)
	fmt.Println(countProduct)
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
	respData["total"] = len(countProduct)
	Respon.Status = 200
	Respon.Data = respData
	Respon.Message = "success"
	return Respon
}
func GetProductDetail(uuid string) utils.Respon {
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
						WHERE 
							prod.uuid = '%s'`, uuid)

	dbEngine := db.ConnectDB()

	getproduct, err := dbEngine.QueryString(query)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	if len(getproduct) == 0 {
		Respon.Status = 404
		Respon.Message = "Product not found"
		return Respon
	}

	getSize, err := dbEngine.QueryString(`select ps.size_id,rs.size from product_size as ps 
										left join ref_size as rs on ps.size_id::integer = rs.id 
										where ps.product_id=?
										order by rs.id asc`, getproduct[0]["uuid"])
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	getImage, err := dbEngine.QueryString(`SELECT img.file_name FROM product_image AS pi
										LEFT JOIN images AS img ON pi.image_id = img.uuid
										WHERE pi.product_id=?`, getproduct[0]["uuid"])
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	var images []string

	for i := 0; i < len(getImage); i++ {
		images = append(images, fmt.Sprintf("%s/file/product/%s", Config.ServerHost, getImage[i]["file_name"]))
	}

	product := make(map[string]interface{})

	intPrice, _ := strconv.Atoi(getproduct[0]["price"])
	intDiscount, _ := strconv.Atoi(getproduct[0]["discount"])
	intPriceDiscount := intPrice - intDiscount

	product["id"] = getproduct[0]["uuid"]
	product["images"] = images
	product["title"] = getproduct[0]["product_name"]
	product["price"] = intPrice
	product["after_discount"] = intPriceDiscount
	product["size"] = getSize
	product["description"] = getproduct[0]["description"]
	product["category"] = getproduct[0]["category_name"]
	// product["date_created"] = getproduct[0]["date_created"]

	Respon.Status = 200
	Respon.Data = product
	Respon.Message = "success"
	return Respon
}
func GetProductReview(product_id string) utils.Respon {
	var Respon utils.Respon

	query := fmt.Sprintf(`SELECT usr.fullname, usr.photo, rvw.comment
						FROM product_review AS rvw
						LEFT JOIN order_detail AS od ON rvw.order_id=od.uuid
						LEFT JOIN product AS prod ON od.product_id=prod.uuid
						LEFT JOIN trans_order AS trans ON od.trans_order_id=trans.uuid
						LEFT JOIN users AS usr ON trans.user_id=usr.uuid
						WHERE prod.uuid='%s'
						ORDER BY rvw.date_create DESC`, product_id)

	dbEngine := db.ConnectDB()

	getReview, err := dbEngine.QueryString(query)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if len(getReview) == 0 {
		Respon.Status = 404
		Respon.Message = "Review not found"
		return Respon
	}

	Respon.Status = 200
	Respon.Data = getReview
	Respon.Message = "success"
	return Respon
}
func GetProductCategory() utils.Respon {
	var Respon utils.Respon

	query := `SELECT ref_category_product.id,ref_category_product.category_name,ref_category_product.description,images.file_name FROM ref_category_product LEFT JOIN images ON images.uuid=ref_category_product.image_id ORDER BY ref_category_product.category_name ASC`

	dbEngine := db.ConnectDB()

	getCategory, err := dbEngine.QueryString(query)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	arraycategory := make([]interface{}, 0, len(getCategory))
	for i := 0; i < len(getCategory); i++ {
		category := make(map[string]interface{})

		intid, _ := strconv.Atoi(getCategory[i]["id"])

		category["id"] = intid
		category["category_name"] = getCategory[i]["category_name"]
		category["description"] = getCategory[i]["description"]
		category["image"] = fmt.Sprintf("%s/file/category/%s", Config.ServerHost, getCategory[i]["file_name"])
		arraycategory = append(arraycategory, category)
	}

	Respon.Status = 200
	Respon.Data = arraycategory
	Respon.Message = "success"
	return Respon
}
func GetProductSize() utils.Respon {
	var Respon utils.Respon

	query := `SELECT id,size FROM ref_size ORDER BY size DESC`

	dbEngine := db.ConnectDB()

	getSize, err := dbEngine.QueryString(query)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Data = getSize
	Respon.Message = "success"
	return Respon
}
