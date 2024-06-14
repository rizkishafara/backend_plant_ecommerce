package controller_admin

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	admin "tanaman/model/Admin"
	"tanaman/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var response utils.Respon

func AddProduct(c *fiber.Ctx) error {
	productid := uuid.New().String()

	name := c.FormValue("name")
	description := c.FormValue("description")
	price := c.FormValue("price")
	discount := c.FormValue("discount")
	category_id := c.FormValue("category_id")
	images := c.FormValue("images")

	var img []string
	_ = json.Unmarshal([]byte(images), &img)

	for _, image := range img {
		// Dekode data base64 menjadi byte
		fileBytes, err := base64.StdEncoding.DecodeString(image)
		if err != nil {
			response.Status = 500
			response.Message = "Invalid base64 data"
			return c.JSON(response)
		}
		fileType := http.DetectContentType(fileBytes)
		allowedTypes := map[string]bool{
			"image/jpg":  true,
			"image/jpeg": true,
		}
		if !allowedTypes[fileType] {
			response.Status = 500
			response.Message = "Unsupported file type"
			return c.JSON(response)
		}

		fileID := uuid.New()

		fileType = strings.Split(fileType, "/")[1]

		newFileName := fmt.Sprintf("product_%s.%s", fileID, fileType)

		_, err = utils.SaveFile(newFileName, fileType, "product", fileBytes)

		if err != nil {
			response.Status = 500
			response.Message = err.Error()
			return c.JSON(response)
		}

		// fmt.Println(image)
		go admin.AddProductImage(newFileName, time.Now().Format("2006-01-02"), productid)

	}

	if name == "" || description == "" || price == "" || discount == "" || category_id == "" {
		response.Status = 404
		response.Message = "Data must be filled"
		return c.JSON(response)
	}

	data, err := admin.AddProduct(productid, name, description, price, discount, category_id, time.Now().Format("2006-01-02"))
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		return c.JSON(response)
	}

	return c.JSON(data)
}
func UpdateProduct(c *fiber.Ctx) error {
	productid := c.FormValue("productid")

	name := c.FormValue("name")
	description := c.FormValue("description")
	price := c.FormValue("price")
	discount := c.FormValue("discount")
	category_id := c.FormValue("category_id")

	if name == "" || description == "" || price == "" || discount == "" || category_id == "" {
		response.Status = 404
		response.Message = "Data must be filled"
		return c.JSON(response)
	}

	data, err := admin.UpdateProduct(productid, name, description, price, discount, category_id, time.Now().Format("2006-01-02"))
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		return c.JSON(response)
	}

	return c.JSON(data)
}
func DeleteProduct(c *fiber.Ctx) error {
	productid := c.FormValue("productid")

	if productid == "" {
		response.Status = 404
		response.Message = "Data must be filled"
		return c.JSON(response)
	}

	delFile, err := admin.DeleteAllProductImage(productid)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		return c.JSON(response)
	}
	if delFile.Status != 200 {
		return c.JSON(delFile)
	}

	data, err := admin.DeleteProduct(productid)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		return c.JSON(response)
	}

	return c.JSON(data)
}
