package model

import (
	"fmt"
	"tanaman/config"
	"tanaman/db"
	"tanaman/utils"
)

var Config = config.LoadConfig(".")
func GetProfileUser(id, email string) utils.Respon {
	dbEngine := db.ConnectDB()

	var Respon utils.Respon
	getProfile, err := dbEngine.QueryString(`SELECT uuid as id, email, fullname, photo FROM users WHERE uuid=(?) AND email=(?)`, id, email)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if getProfile == nil {
		Respon.Status = 404
		Respon.Message = "User tidak terdaftar"
		return Respon
	}
	var dataResponse = make(map[string]interface{})

	if getProfile != nil {
		dataResponse["id"] = getProfile[0]["id"]
		dataResponse["email"] = getProfile[0]["email"]
		dataResponse["fullname"] = getProfile[0]["fullname"]
		// dataResponse["photo"] = getProfile[0]["photo"]
		dataResponse["photo"] = fmt.Sprintf("%s/profile/%s", Config.ServerHost, getProfile[0]["photo"])
	}

	Respon.Status = 200
	Respon.Data = dataResponse
	Respon.Message = "success"
	return Respon
}
func UpdateProfile(fullname, photo, id, date string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	oldPhoto, _ := dbEngine.QueryString(`SELECT photo FROM users WHERE uuid=(?)`, id)

	var query string

	if photo != "" && fullname != "" {
		query = "UPDATE users SET fullname = '" + fullname + "', photo ='" + photo + ", date_update = '" + date + "' WHERE uuid = '" + id + "'"
	} else if photo != "" && fullname == "" {
		query = "UPDATE users SET photo ='" + photo + ", date_update = '" + date + "' WHERE uuid = '" + id + "'"
	} else if fullname != "" {
		query = "UPDATE users SET fullname = '" + fullname + ", date_update = '" + date + "' WHERE uuid = '" + id + "'"
	} else {
		Respon.Status = 404
		Respon.Message = "Tidak ada yang diubah"
		return Respon
	}

	_, err := dbEngine.QueryString(query)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	if photo != "" && oldPhoto[0]["photo"] != "" {
		go utils.DeleteFile(oldPhoto[0]["photo"], "profile")
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func GetCountVoucher(id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon
	getCountVoucher, err := dbEngine.QueryString(`SELECT COUNT(claim.uuid) AS count FROM claim 
												INNER JOIN logs_claim_voucher ON claim.uuid=logs_claim_voucher.claim_id
												WHERE claim.user_id=(?) AND claim.active='1'`, id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	var datares = make(map[string]interface{})
	if getCountVoucher != nil {
		datares["count"] = getCountVoucher[0]["count"]
	}
	Respon.Status = 200
	Respon.Data = datares
	Respon.Message = "success"
	return Respon
}
func GetCountLoyalty(id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	totalLoyalty, err := dbEngine.QueryString(`SELECT 
													(COALESCE(loyalty, 0) - COALESCE(claim, 0)) AS count
												FROM
													(
														SELECT
															(SELECT COALESCE(SUM(quantity), 0) FROM loyalty WHERE user_id = (?)) AS loyalty,
															(SELECT COALESCE(SUM(nominal_loyalty), 0) FROM claim WHERE user_id = (?)) AS claim
													) AS counts`, id, id)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	var datares = make(map[string]interface{})
	if totalLoyalty != nil {
		datares["count"] = totalLoyalty[0]["count"]
	}

	Respon.Status = 200
	Respon.Data = datares
	Respon.Message = "success"
	return Respon
}
