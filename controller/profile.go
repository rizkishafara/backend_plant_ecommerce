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
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GetProfileUser(c *fiber.Ctx) error {

	fmt.Println("GetProfileUser")
	email := utils.GetValJWT(c.Locals("user").(*jwt.Token), "email")
	id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")
	fmt.Println("email", email)
	fmt.Println("id", id)

	if id == "" || email == "" {
		response.Status = 404
		response.Message = "User not registered"
		return c.JSON(response)
	}

	return c.JSON(model.GetProfileUser(id, email))
}
func UpdateProfileUser(c *fiber.Ctx) error {
	id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")

	if id == "" {
		response.Status = 404
		response.Message = "User not registered"
		return c.JSON(response)
	}

	fullname := c.FormValue("fullname")
	photo := c.FormValue("photo")
	phototype := c.FormValue("phototype")

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
		data := model.UpdateProfile(fullname, newFileName, id, time.Now().Format("2006-01-02"))

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
		data := model.UpdateProfile(fullname, "", id, time.Now().Format("2006-01-02"))
		return c.JSON(data)
	}
}
func GetCountVoucher(c *fiber.Ctx) error {

	email := utils.GetValJWT(c.Locals("user").(*jwt.Token), "email")
	id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")

	if id == "" || email == "" {
		response.Status = 404
		response.Message = "JWT Invalid"
		return c.JSON(response)
	}

	return c.JSON(model.GetCountVoucher(id))
}
func GetCountLoyalty(c *fiber.Ctx) error {

	email := utils.GetValJWT(c.Locals("user").(*jwt.Token), "email")
	id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")

	if id == "" || email == "" {
		response.Status = 404
		response.Message = "JWT Invalid"
		return c.JSON(response)
	}

	return c.JSON(model.GetCountLoyalty(id))
}
