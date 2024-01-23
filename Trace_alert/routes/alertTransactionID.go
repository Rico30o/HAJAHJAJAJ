package routes

import (
	"instapay/db"
	"instapay/model"

	"github.com/gofiber/fiber/v2"
)

func GetTransacInfo(c *fiber.Ctx) error {

	info := &model.TransacAlertRequest{}
	resp := &model.ResponseTransac{}

	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "ALERT_FINANCIAL_CRIME",
						"ReasonCode":  "BAD_REQUEST",
						"Description": "We could not handle your request",
						"Recoverable": false,
						"Details":     "The request contains a bad payload",
					},
				},
			},
		})
	} else if fetchErr := db.DB.Debug().Raw(`SELECT * FROM response_transac WHERE txn_alert_id = ? `, info.TransactionAlertID).Scan(&resp).Error; fetchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "ALERT_FINANCIAL_CRIME",
						"ReasonCode":  "BAD_REQUEST",
						"Description": "We could not handle your request",
						"Recoverable": false,
						"Details":     "The request contains a bad payload",
					},
				},
			},
		})
	}

	return c.JSON(resp)
}
