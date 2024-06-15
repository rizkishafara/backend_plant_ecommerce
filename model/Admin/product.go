package admin

import (
	"fmt"
	"tanaman/db"
	"tanaman/utils"

	"github.com/google/uuid"
)

func AddProduct(uuid, name, description, price, discount, category_id, date string) (utils.Respon, error) {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`INSERT INTO product (uuid, product_name, description, price, discount, category_id, date_created) VALUES (?,?,?,?,?,?,?)`, uuid, name, description, price, discount, category_id, date)

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
func AddSizeProduct(size_id, date, product_id string) {
	dbEngine := db.ConnectDB()
	uuidPivot := uuid.New()

	_, err := dbEngine.QueryString(`INSERT INTO product_size (uuid, size_id, product_id, date_create) VALUES (?,?,?,?)`, uuidPivot.String(), size_id, product_id, date)

	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func UpdateProduct(uuid, name, description, price, discount, category_id, date string) (utils.Respon, error) {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`UPDATE product SET product_name=?, description=?, price=?, discount=?, category_id=?, date_update=? WHERE uuid=?`, name, description, price, discount, category_id, date, uuid)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon, err
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon, nil
}
func DeleteProduct(uuid string) (utils.Respon, error) {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString(`DELETE FROM product WHERE uuid=?`, uuid)

	if err != nil {
		return Respon, err
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon, nil
}
func DeleteAllProductImage(product_id string) (utils.Respon, error) {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	getimageId, err := dbEngine.QueryString(`SELECT image_id FROM product_image WHERE product_id=?`, product_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon, err
	}
	for i := 0; i < len(getimageId); i++ {
		getnameimage, err := dbEngine.QueryString(`SELECT file_name FROM images WHERE uuid=?`, getimageId[i]["image_id"])
		if err == nil && len(getnameimage) > 0 {
			utils.DeleteFile(getnameimage[0]["file_name"], "product")
		}
		_, err = dbEngine.QueryString(`DELETE FROM images WHERE uuid=?`, getimageId[i]["image_id"])
		if err != nil {
			Respon.Status = 500
			Respon.Message = err.Error()
			return Respon, err
		}
	}

	_, err = dbEngine.QueryString(`DELETE FROM product_image WHERE product_id=?`, product_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon, err
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon, nil
}
