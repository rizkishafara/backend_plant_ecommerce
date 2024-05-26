package model

import (
	"tanaman/db"
	"tanaman/utils"
)

func GetProfileUser(id, email string) utils.Respon {
	dbEngine := db.ConnectDB()

	var Respon utils.Respon
	getProfile, err := dbEngine.QueryString(`SELECT uuid as id, email, fullname, photo FROM users WHERE uuid=(?) AND email=(?)`, id, email)
	if err != nil {
		Respon.Success = false
		Respon.Message = err.Error()
		return Respon
	}
	if getProfile == nil {
		Respon.Success = false
		Respon.Message = "User tidak terdaftar"
		return Respon
	}
	var dataResponse = make(map[string]interface{})

	if getProfile != nil {
		dataResponse["id"] = getProfile[0]["id"]
		dataResponse["email"] = getProfile[0]["email"]
		dataResponse["fullname"] = getProfile[0]["fullname"]
		dataResponse["photo"] = getProfile[0]["photo"]
	}

	Respon.Success = true
	Respon.Data = dataResponse
	Respon.Message = "Berhasil Get Profile"
	return Respon
}
func UpdateProfile(fullname, photo, id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	var query string

	if photo != "" {
		query = "UPDATE users SET fullname = '" + fullname + "', photo ='" + photo + "' WHERE uuid = '" + id + "'"
	} else {
		query = "UPDATE users SET fullname = '" + fullname + "' WHERE uuid = '" + id + "'"
	}

	_, err := dbEngine.QueryString(query)
	if err != nil {
		Respon.Success = false
		Respon.Message = err.Error()
		return Respon
	}
	Respon.Success = true
	Respon.Message = "Berhasil Update Profile"
	return Respon
}
func GetCountVoucher(id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon
	getCountVoucher, err := dbEngine.QueryString(`SELECT COUNT(claim.uuid) AS count FROM claim 
												INNER JOIN logs_claim_voucher ON claim.uuid=logs_claim_voucher.claim_id
												WHERE claim.user_id=(?) AND claim.active='1'`, id)
	if err != nil {
		Respon.Success = false
		Respon.Message = err.Error()
		return Respon
	}
	var datares = make(map[string]interface{})
	if getCountVoucher != nil {
		datares["count"] = getCountVoucher[0]["count"]
	}
	Respon.Success = true
	Respon.Data = datares
	Respon.Message = "Berhasil get count voucher"
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
		Respon.Success = false
		Respon.Message = err.Error()
		return Respon
	}
	var datares = make(map[string]interface{})
	if totalLoyalty != nil {
		datares["count"] = totalLoyalty[0]["count"]
	}

	Respon.Success = true
	Respon.Data = datares
	Respon.Message = "Berhasil get count loyalty"
	return Respon
}
