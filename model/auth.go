package model

import (
	"fmt"
	"strings"
	"tanaman/db"
	"tanaman/utils"
	"time"
)

func Register(email, password, fullname, datecreate, uuid, photo string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	cekemail, err := dbEngine.QueryString(`SELECT email FROM users WHERE email=(?)`, email)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	if cekemail != nil {
		Respon.Status = 404
		Respon.Message = "Email sudah terdaftar"
		return Respon
	}
	passnew, err := utils.HashPassword(password)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	_, err = dbEngine.QueryString("INSERT INTO users (uuid,email,PASSWORD,fullname,photo,date_create) VALUES (?,?,?,?,?,?)", uuid, email, passnew, fullname, photo, datecreate)

	if err != nil {
		// log.Fatal(err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func Login(email, password string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	cekemail, err := dbEngine.QueryString(`SELECT email, password FROM users WHERE email=(?)`, email)
	if err != nil {
		// log.Fatal(err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	if cekemail == nil {
		Respon.Status = 404
		Respon.Message = "Email tidak terdaftar"
		return Respon
	}
	errc := utils.CheckPasswordHash(password, cekemail[0]["password"])
	if !errc {
		Respon.Status = 404
		Respon.Message = err.Error()
		return Respon
	}

	datauser, err := dbEngine.QueryString("SELECT uuid, email FROM users WHERE email = ? ", email)

	if err != nil {
		// log.Fatal(err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	if datauser == nil {
		Respon.Status = 404
		Respon.Message = "Email atau Password salah!"
		return Respon
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	expiredtime := time.Now().Local().In(loc).Add(time.Hour * 24).Format("2006-01-02 15:04:05")

	result, code := utils.Generatejwt(email, datauser[0]["uuid"], expiredtime)
	if code != 200 {
		Respon.Status = code
		Respon.Message = result
		return Respon
	}

	datares := make(map[string]interface{})
	if datauser != nil {
		datares["email"] = datauser[0]["email"]
		datares["id"] = datauser[0]["uuid"]
		datares["token"] = result
	}
	Respon.Status = 200
	Respon.Data = datares
	Respon.Message = "success"
	return Respon
}
func ForgotPassword(email string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	cekemail, err := dbEngine.QueryString(`SELECT uuid,email FROM users WHERE email=(?)`, email)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	if cekemail == nil {
		Respon.Status = 404
		Respon.Message = "Email tidak terdaftar"
		return Respon
	}
	fmt.Println(cekemail)
	fmt.Println(cekemail[0]["uuid"])

	newparam := fmt.Sprintf("%s,%s,%s", email, cekemail[0]["uuid"], time.Now())

	encrypted, err := utils.Encrypt(newparam)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	fmt.Println("Encrypted:", encrypted)

	errs := utils.SendForgotPassword(email, encrypted)
	if errs != nil {
		Respon.Status = 500
		Respon.Message = errs.Error()
		return Respon
	}
	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func UpdatePassword(param, password string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon
	decrypted, err := utils.Decrypt(param)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	fmt.Println("Decrypted:", decrypted)

	email := strings.Split(decrypted, ",")[0]
	uuid := strings.Split(decrypted, ",")[1]

	passnew, err := utils.HashPassword(password)
	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	_, err = dbEngine.QueryString("UPDATE users SET password = ? WHERE email = ? AND uuid = ?", passnew, email, uuid)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
