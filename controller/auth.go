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
		response.Success = false
		response.Message = "Email dan Password wajib diisi"
		return c.JSON(response)
	}

	if !utils.IsValidEmail(email) {
		response.Success = false
		response.Message = "Email tidak valid"
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
	phototype := c.FormValue("phototype")

	if !utils.IsValidEmail(email) {
		response.Success = false
		response.Message = "Email tidak valid"
		return c.JSON(response)
	}

	if email == "" || password == "" || fullname == "" {
		response.Success = false
		response.Message = "Data wajib diisi"
		return c.JSON(response)
	}

	createUUid := uuid.New()
	if photo != "" {
		allowedTypes := map[string]bool{
			"jpg":  true,
			"jpeg": true,
		}
		if !allowedTypes[phototype] {
			return c.Status(fiber.StatusBadRequest).SendString("Unsupported file type")
		}

		// Dekode data base64 menjadi byte
		fileBytes, err := base64.StdEncoding.DecodeString(photo)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid base64 data")
		}

		fileID := uuid.New()

		newFileName := fmt.Sprintf("profile_%s.%s", fileID, phototype)

		data := model.Register(email, password, fullname, time.Now().Format("2006-01-02"), createUUid.String(), newFileName)

		if data.Success {

			filePath := filepath.Join("./uploads/profile", newFileName)
			if err := ioutil.WriteFile(filePath, fileBytes, 0644); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Unable to save file")
			}
		}
		return c.JSON(data)
	} else {
		data := model.Register(email, password, fullname, time.Now().Format("2006-01-02"), createUUid.String(), "")

		return c.JSON(data)
	}
}
func ForgotPassword(c *fiber.Ctx) error {
	email := c.FormValue("email")

	if email == "" {
		response.Success = false
		response.Message = "Email dan Password wajib diisi"
		return c.JSON(response)
	}

	if !utils.IsValidEmail(email) {
		response.Success = false
		response.Message = "Email tidak valid"
		return c.JSON(response)
	}
	return c.JSON(model.ForgotPassword(email))
}
func UpdatePassword(c *fiber.Ctx) error {
	param := c.FormValue("param")
	password := c.FormValue("password")

	if param == "" || password == "" {
		response.Success = false
		response.Message = "value null"
		return c.JSON(response)
	}

	return c.JSON(model.UpdatePassword(param, password))
}
