package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"tanaman/model"
	"tanaman/utils"

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
		response.Success = false
		response.Message = "User tidak terdaftar"
		return c.JSON(response)
	}

	return c.JSON(model.GetProfileUser(id, email))
}
func UpdateProfileUser(c *fiber.Ctx) error {
	id := c.FormValue("id")
	fullname := c.FormValue("fullname")
	photo := c.FormValue("photo")
	phototype := c.FormValue("phototype")

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
		data := model.UpdateProfile(fullname, newFileName, id)

		if data.Success {

			filePath := filepath.Join("./uploads/profile", newFileName)
			if err := ioutil.WriteFile(filePath, fileBytes, 0644); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Unable to save file")
			}
		}
		return c.JSON(data)
	} else {
		data := model.UpdateProfile(fullname, "", id)
		return c.JSON(data)
	}
}
func GetCountVoucher(c *fiber.Ctx) error {

	email := utils.GetValJWT(c.Locals("user").(*jwt.Token), "email")
	id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")

	if id == "" || email == "" {
		response.Success = false
		response.Message = "JWT Invalid"
		return c.JSON(response)
	}

	return c.JSON(model.GetCountVoucher(id))
}
func GetCountLoyalty(c *fiber.Ctx) error {

	email := utils.GetValJWT(c.Locals("user").(*jwt.Token), "email")
	id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")

	if id == "" || email == "" {
		response.Success = false
		response.Message = "JWT Invalid"
		return c.JSON(response)
	}

	return c.JSON(model.GetCountLoyalty(id))
}
