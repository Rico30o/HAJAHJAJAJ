package routes

import (
	"instapay/db"
	"instapay/model"
	errorhandling "instapay/model/errorHandling"

	"github.com/gofiber/fiber/v2"
)

func GetAccInfo(c *fiber.Ctx) error {

	// Parse the request body
	info := &model.AccountAlertRequest1{}
	resp := &model.ResponseInfo1{}

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
	} else if fetchErr := db.DB.Debug().Raw(`SELECT * FROM trace_alerts WHERE accountid = ?`, info.Accountid).Scan(&resp).Error; fetchErr != nil {
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
	var alertid []model.ResponseInfo1

	if len(alertid) == 0 {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	return c.JSON(resp)
}

//200
//400
//404
//405
//429

//401
//403
