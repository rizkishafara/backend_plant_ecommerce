package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"tanaman/model"
	"tanaman/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var response utils.Respon

func Login(c *fiber.Ctx) error {
	fmt.Println("Login")
	// return c.JSON(response)
	// BODY
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		response.Status = 404
		response.Message = "Email and Password must be filled"
		return c.JSON(response)
	}

	if !utils.IsValidEmail(email) {
		response.Status = 404
		response.Message = "Email is not valid"
		return c.JSON(response)
	}

	data := model.Login(email, password)

	return c.JSON(data)
}
func Register(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fullname := c.FormValue("fullname")
	photo := c.FormValue("photo")
	phone := c.FormValue("phone")
	phototype := c.FormValue("phototype")

	fmt.Println("email masuk: ",email)

	if !utils.IsValidEmail(email) {
		response.Status = 404
		response.Message = "Email is not valid"
		return c.JSON(response)
	}

	if email == "" || password == "" || fullname == "" {
		response.Status = 404
		response.Message = "Email, Password and Fullname must be filled"
		return c.JSON(response)
	}

	createUUid := uuid.New()
	if photo != "" {
		allowedTypes := map[string]bool{
			"jpg":  true,
			"jpeg": true,
		}
		if !allowedTypes[phototype] {
			response.Status = 500
			response.Message = "Unsupported file type"
			return c.JSON(response)
		}

		// Dekode data base64 menjadi byte
		fileBytes, err := base64.StdEncoding.DecodeString(photo)
		if err != nil {
			response.Status = 500
			response.Message = "Invalid base64 data"
			return c.JSON(response)
		}

		fileID := uuid.New()

		newFileName := fmt.Sprintf("profile_%s.%s", fileID, phototype)

		data := model.Register(email, password, fullname, time.Now().Format("2006-01-02"), createUUid.String(), newFileName, phone)

		if data.Status == 200 {

			filePath := filepath.Join("./uploads/profile", newFileName)
			if err := ioutil.WriteFile(filePath, fileBytes, 0644); err != nil {
				response.Status = 500
				response.Message = "Unable to save file"
				return c.JSON(response)
			}
		}
		return c.JSON(data)
	} else {
		data := model.Register(email, password, fullname, time.Now().Format("2006-01-02"), createUUid.String(), "", phone)

		return c.JSON(data)
	}
}
func ForgotPassword(c *fiber.Ctx) error {
	email := c.FormValue("email")

	if email == "" {
		response.Status = 404
		response.Message = "Email must be filled"
		return c.JSON(response)
	}

	if !utils.IsValidEmail(email) {
		response.Status = 404
		response.Message = "Email is not valid"
		return c.JSON(response)
	}
	return c.JSON(model.ForgotPassword(email))
}
func UpdatePassword(c *fiber.Ctx) error {
	param := c.FormValue("param")
	password := c.FormValue("password")

	if param == "" || password == "" {
		response.Status = 404
		response.Message = "value null"
		return c.JSON(response)
	}

	return c.JSON(model.UpdatePassword(param, password))
}
func GetFile(c *fiber.Ctx) error {
	jenis := c.Params("jenis")
	file := c.Params("file")
	return c.SendFile(fmt.Sprintf("./uploads/%s/%s", jenis, file))
}
