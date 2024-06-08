package model

import (
	"fmt"
	"tanaman/db"
	"tanaman/utils"

	"github.com/google/uuid"
)

func AddProduct(uuid, name, description, price, discount, category_id,date string) (utils.Respon, error) {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`INSERT INTO product (uuid, product_name, description, price, discount, category_id, date_created) VALUES (?,?,?,?,?,?,?)`, uuid, name, description, price, discount, category_id,date)

	if err != nil {
		return Respon, err
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon, nil
}
func AddProductImage(file_name, date, product_id string) {
	dbEngine := db.ConnectDB()
	uuidPivot := uuid.New()
	uuidImage := uuid.New()

	_, err := dbEngine.QueryString(`INSERT INTO images (uuid, file_name, date_create) VALUES (?,?,?)`, uuidImage.String(), file_name, date)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = dbEngine.QueryString(`INSERT INTO product_image (uuid, image_id, product_id, date_create) VALUES (?,?,?,?)`, uuidPivot.String(), uuidImage.String(), product_id, date)

	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
