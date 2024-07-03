package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"tanaman/model"
	"tanaman/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func PostCheckout(c *fiber.Ctx) error {
	user_id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")
	order_number := generateOrderNumber()
	payment_id := c.FormValue("payment_id")
	// json_alamat := c.FormValue("json_alamat")
	notes := c.FormValue("notes")
	resi := "RESI/56789876"
	sub_total := c.FormValue("sub_total")
	discount := c.FormValue("discount")
	idshipping := c.FormValue("shipping_id")
	alamat:= c.FormValue("alamat")

	shippingcost := model.GetShippingCost(idshipping)

	intSubTotal, _ := strconv.Atoi(sub_total)
	intDiscount, _ := strconv.Atoi(discount)
	intShippingCost, _ := strconv.Atoi(shippingcost)
	total_payment := intSubTotal + intShippingCost - intDiscount

	stotal_payment := fmt.Sprintf("%d", total_payment)

	productcart := c.FormValue("product")

	var prodcart []string
	_ = json.Unmarshal([]byte(productcart), &prodcart)

	for _, prod := range prodcart {
		// fmt.Println(prod)
		getprod := model.GetCart(user_id, prod)

		prdct := getprod.Data

		// fmt.Println(prdct)
		arrayproduct, _ := prdct.([]interface{})

		for _, prd := range arrayproduct {

			product_id := prd.(map[string]interface{})["product_id"].(string)
			size := prd.(map[string]interface{})["size"].(string)
			qty := prd.(map[string]interface{})["quantity"].(string)
			product_name := prd.(map[string]interface{})["title"].(string)
			discount_price := prd.(map[string]interface{})["price_discount"].(string)
			image := prd.(map[string]interface{})["img"].(string)

			model.PostCheckoutDetail(uuid.New().String(), order_number, size, qty, product_name, discount_price, image, time.Now().Format("2006-01-02"), product_id)
		
		}

	}

	chekout := model.PostCheckout(uuid.New().String(), user_id, "1", order_number, payment_id, notes, resi, sub_total, shippingcost, discount, stotal_payment, time.Now().Format("2006-01-02"), alamat)
	return c.JSON(chekout)
}

func GetProductOrder(c *fiber.Ctx) error {
	user_id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")
	order_id := c.Params("order_id")

	return c.JSON(model.GetProductOrder(user_id, order_id))

}

func generateOrderNumber() string {
	// Get current date and time
	currentTime := time.Now()
	dateTimeFormat := currentTime.Format("02012006150405") // ddmmyyyyhhmmss

	// Generate a unique identifier (xxx part)
	rand.Seed(time.Now().UnixNano())
	uniqueID := rand.Intn(1000) // random number between 0 and 999

	// Combine them into the desired format
	orderNumber := fmt.Sprintf("INV/%s%03d", dateTimeFormat, uniqueID)

	return orderNumber
}
