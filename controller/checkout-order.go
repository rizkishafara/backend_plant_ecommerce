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
	resi := c.FormValue("resi")
	sub_total := c.FormValue("sub_total")
	discount := c.FormValue("discount")
	idshipping := c.FormValue("idshipping")

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

		prdct:= getprod.Data

		fmt.Println(prdct)

		// for i:=0; i<len(prdct); i++ {
		// 	model.PostCheckoutDetail(uuid.New().String(), order_number, prdct[i].Size, prdct[i].Quantity, prdct[i].ProductName, prdct[i].DiscountPrice, prdct[i].Image, time.Now().Format("2006-01-02"), prdct[i].ProductID)
		// }
		
	}

	var alamat []string
	chekout := model.PostCheckout(uuid.New().String(), user_id, "0", order_number, payment_id, notes, resi, sub_total, shippingcost, discount, stotal_payment, time.Now().Format("2006-01-02"), alamat)
	return c.JSON(chekout)
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
