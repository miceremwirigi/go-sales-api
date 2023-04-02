package cashierControllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/miceremwirigi/go-sales-api/config/dbConnection"
	"github.com/miceremwirigi/go-sales-api/models"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			},
		)
	}
	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier Name is required",
			},
		)
	}
	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode is required",
			},
		)
	}

	// save cashier in db
	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"messae":  "Cashier created successfully",
			"data":    cashier,
		},
	)
}

func CashiersList(c *fiber.Ctx) error {
	var cashier []models.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier list api",
			"data":    cashier,
		},
	)

}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Select("id,name, passcode, created_at, updated_at").Where("id = ?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["cashierName"] = cashier.Name
	cashierData["cashierPasscode"] = cashier.Passcode
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt

	// if cashier is not in database throw error
	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
				"error":   map[string]interface{}{},
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": fmt.Sprintf("Success. cashier id %s found", cashierId),
			"data":    cashierData,
		},
	)
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id = ?", cashierId)

	// check if cashier with this id exists
	if cashier.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			},
		)
	}

	var updateCashier models.Cashier
	err := c.BodyParser(&updateCashier)
	if err != nil {
		return err
	}

	if updateCashier.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier name is required",
			},
		)
	}

	// query to update the record
	cashier.Name = updateCashier.Name
	cashier.Passcode = updateCashier.Passcode
	cashier.UpdatedAt = time.Now()

	db.DB.Save(&cashier)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": fmt.Sprintf("Cashier id %s updated susscessfully", cashierId),
		},
	)
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Where("id = ?", cashierId).First(&cashier)
	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			},
		)
	}

	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": fmt.Sprintf("Cashier id %s deleted susscessfully", cashierId),
		},
	)
}
