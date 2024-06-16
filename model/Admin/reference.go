package admin

import (
	"tanaman/db"
	"tanaman/utils"
)

func AddCategoryProduct(uuid, category, date, image_id, description string) error {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`INSERT INTO ref_category_product (uuid, category_name, date_create, image_id, description) VALUES (?,?,?,?,?)`, uuid, category, date, image_id, description)

	if err != nil {
		return err
	}

	Respon.Status = 200
	Respon.Message = "success"
	return nil
}
func AddCategoryProductImage(uuid, file_name, date string) error {
	dbEngine := db.ConnectDB()

	_, err := dbEngine.QueryString(`INSERT INTO images (uuid, file_name, date_create) VALUES (?,?,?)`, uuid, file_name, date)

	if err != nil {
		return err
	}

	return nil
}
