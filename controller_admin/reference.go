package controller_admin

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	admin "tanaman/model/Admin"
	"tanaman/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddCategoryProduct(c *fiber.Ctx) error {
	catUUID := uuid.New().String()

	category := c.FormValue("category")
	date := time.Now().Format("2006-01-02")
	image_id := uuid.New().String()
	description := c.FormValue("description")
	image := c.FormValue("image")

	if category == "" || image == "" {
		response.Status = 404
		response.Message = "All fields are required"
		return c.JSON(response)
	}

	fileBytes, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		response.Status = 500
		response.Message = "Invalid base64 data"
		return c.JSON(response)
	}
	fileType := http.DetectContentType(fileBytes)
	fmt.Println("fileType ", fileType)
	if !utils.AllowedTypes[fileType] {
		response.Status = 500
		response.Message = "Unsupported file type"
		return c.JSON(response)
	}

	fileID := uuid.New()

	fileType = strings.Split(fileType, "/")[1]

	newFileName := fmt.Sprintf("category_%s.%s", fileID, fileType)

	err = admin.AddCategoryProductImage(image_id, newFileName, time.Now().Format("2006-01-02"))
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		return c.JSON(response)
	}

	_, err = utils.SaveFile(newFileName, fileType, "category", fileBytes)

	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		return c.JSON(response)
	}

	err = admin.AddCategoryProduct(catUUID, category, date, image_id, description)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		return c.JSON(response)
	}

	response.Status = 200
	response.Message = "success"

	return c.JSON(response)
}
