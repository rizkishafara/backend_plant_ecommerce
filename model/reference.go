package model

import (
	"tanaman/db"
	"tanaman/utils"
)

func GetProvince() utils.Respon {
	dbEngine := db.ConnectDB()

	var Respon utils.Respon
	getProvince, err := dbEngine.QueryString(`SELECT id, province_id as province FROM ref_province`)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if getProvince == nil {
		Respon.Status = 404
		Respon.Message = "Province not found"
		return Respon
	}

	Respon.Status = 200
	Respon.Data = getProvince
	Respon.Message = "success"
	return Respon
}
func GetCity(province_id string) utils.Respon {
	dbEngine := db.ConnectDB()

	var Respon utils.Respon
	getCity, err := dbEngine.QueryString(`SELECT id, kabkot_name as city FROM ref_kabkot WHERE prov_id=(?)`, province_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if getCity == nil {
		Respon.Status = 404
		Respon.Message = "City not found"
		return Respon
	}

	Respon.Status = 200
	Respon.Data = getCity
	Respon.Message = "success"
	return Respon
}
func GetDistrict(city_id string) utils.Respon {
	dbEngine := db.ConnectDB()

	var Respon utils.Respon
	getDistrict, err := dbEngine.QueryString(`SELECT id, kecamatan_name as district FROM ref_kecamatan WHERE kabkot_id=(?)`, city_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if getDistrict == nil {
		Respon.Status = 404
		Respon.Message = "District not found"
		return Respon
	}

	Respon.Status = 200
	Respon.Data = getDistrict
	Respon.Message = "success"
	return Respon
}
func GetVillage(district_id string) utils.Respon {
	dbEngine := db.ConnectDB()

	var Respon utils.Respon
	getVillage, err := dbEngine.QueryString(`SELECT id, kelurahan_name as village FROM ref_kelurahan WHERE kecamatan_id=(?)`, district_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if getVillage == nil {
		Respon.Status = 404
		Respon.Message = "Village not found"
		return Respon
	}

	Respon.Status = 200
	Respon.Data = getVillage
	Respon.Message = "success"
	return Respon
}
func GetPostalCode(village_id string) utils.Respon {
	dbEngine := db.ConnectDB()

	var Respon utils.Respon
	getPostalCode, err := dbEngine.QueryString(`SELECT id, kodepos as postal_code FROM ref_kodepos WHERE kelurahan_id=(?)`, village_id)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	if getPostalCode == nil {
		Respon.Status = 404
		Respon.Message = "Postal Code not found"
		return Respon
	}

	Respon.Status = 200
	Respon.Data = getPostalCode
	Respon.Message = "success"
	return Respon
}
