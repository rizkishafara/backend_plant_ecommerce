package model

import (
	"encoding/json"
	"tanaman/db"
	"tanaman/utils"
)

func AddProduct(images, name, description, price, discount, category_id, token string) (utils.Respon, error) {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	var imgs map[string]interface{}
	errj := json.Unmarshal([]byte(images), &imgs)
	if errj != nil {
		Respon.Status = 500
		Respon.Message = errj.Error()
		return Respon, errj
	}

	_, err := dbEngine.QueryString(`INSERT INTO products (token) VALUES (?)`, token)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon, err
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon, nil
}
