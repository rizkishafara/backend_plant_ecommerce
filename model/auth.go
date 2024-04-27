package model

import (
	"tanaman/db"
	"tanaman/utils"
	"time"
)

func Login(email, password string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	cekemail, err := dbEngine.QueryString(`SELECT email FROM users WHERE email=(?)`, email)
	if err != nil {
		// log.Fatal(err)
		Respon.Success = false
		Respon.Message = err.Error()
		return Respon
	}

	if cekemail == nil {
		Respon.Success = false
		Respon.Message = "Email tidak terdaftar"
		return Respon
	}

	datauser, err := dbEngine.QueryString("SELECT id, email, fullname FROM users WHERE email = ? AND password = ?", email, password)

	if err != nil {
		// log.Fatal(err)
		Respon.Success = false
		Respon.Message = err.Error()
		return Respon
	}

	if datauser == nil {
		Respon.Success = false
		Respon.Message = "Email atau Password salah!"
		return Respon
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	expiredtime := time.Now().Local().In(loc).Add(time.Hour * 24).Format("2006-01-02 15:04:05")

	result, code := utils.Generatejwt(email, datauser[0]["id_req_pemohon"], expiredtime)
	if code != 200 {
		Respon.Success = false
		Respon.Message = result
		return Respon
	}

	datares := make(map[string]interface{})
	datares["datauser"] = datauser
	datares["token"] = result
	Respon.Success = true
	Respon.Data = datares
	Respon.Message = "Berhasil Login"
	return Respon
}
