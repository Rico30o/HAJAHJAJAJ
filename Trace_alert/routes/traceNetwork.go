package routes

import (
	"fmt"
	"instapay/db"
	"instapay/middleware/loggers"
	"instapay/model"

	"github.com/gofiber/fiber/v2"
)

func Tracenetwork(c *fiber.Ctx) error {
	network := &model.Request_trace{}

	if err := c.BodyParser(&network); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": map[string]interface{}{
				"Error": []map[string]interface{}{
					{
						"Source":      "ALERT_FINANCIAL_CRIME",
						"ReasonCode":  "BAD_REQUEST",
						"Description": "Invalid request body",
						"Recoverable": false,
						"Details":     "The request contains a bad payload",
					},
				},
			},
		})

	}

	var transactions1 []model.Transaction_Response
	var transactions2 []model.Trace_AccountAlert
	var transactions3 []model.NetworkResponse

	// Fetch data from the database for each table
	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_transaction_alert WHERE txnid = ?`, network.Txnid).Scan(&transactions1).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}

	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_account_alert`).Scan(&transactions2).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}

	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_alert WHERE sourcetxntype = ?`, network.Sourcetxntype).Scan(&transactions3).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}

	// Check if no data was found for the given date
	if len(transactions1) == 0 || len(transactions2) == 0 || len(transactions3) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": map[string]interface{}{
				"Error": []map[string]interface{}{
					{
						"Source":      "ALERT_FINANCIAL_CRIME",
						"ReasonCode":  "BAD_REQUEST",
						"Description": "Invalid request body",
						"Recoverable": false,
						"Details":     "The request contains a bad payload",
					},
				},
			},
		})
	}

	// alert := fiber.Map{
	// 	"id":        transactions1[0].ID,
	// 	"Time":      transactions1[0].Time,
	// 	"networkID": transactions1[0].Networkalertid,
	// }

	// return c.JSON(fiber.Map{
	// 	"alerts":            alert,
	// 	"accountAlerts":     transactions1,
	// 	"transactionAlerts": transactions2,
	// 	"network":           transactions3,
	// })

	type alert struct {
		Alertss           fiber.Map                    `json:"alerts"`
		AccountAlerts     []model.Transaction_Response `json:"accountAlerts"`
		TransactionAlerts []model.Trace_AccountAlert   `json:"transactionAlerts"`
		Network           []model.NetworkResponse      `json:"network"`
	}

	responseBody := alert{
		Alertss: fiber.Map{
			"id":        transactions1[0].ID,
			"Time":      transactions1[0].Time,
			"networkID": transactions1[0].Networkalertid,
		},
		AccountAlerts:     transactions1,
		TransactionAlerts: transactions2,
		Network:           transactions3,
	}

	Vizurl := transactions3[0].Vizurl
	Sourcetxnid := transactions3[0].Sourcetxnid
	Sourcetxntype := transactions3[0].Sourcetxntype
	Length := transactions3[0].Length
	Generations := transactions3[0].Generations
	Totalvalue := transactions3[0].Totalvalue
	Sourcevalue := transactions3[0].Sourcevalue
	Uniqueaccounts := transactions3[0].Uniqueaccounts
	Meandwelltime := transactions3[0].Meandwelltime
	Mediandwelltime := transactions3[0].Mediandwelltime
	Meanmulescore := transactions3[0].Meanmulescore
	Elapsedtime := transactions3[0].Elapsedtime
	Numnotinvestigated := transactions3[0].Numnotinvestigated
	Parentalertid := transactions3[0].Parentalertid
	Decisiondate := transactions3[0].Decisiondate
	Mostrecentfeedback := transactions3[0].Mostrecentfeedback

	logMessage := fmt.Sprintf("%s: %s %s %d %d %d %d %d %s %s %.2f %s %d %s %s %s", Vizurl, Sourcetxnid, Sourcetxntype, Length, Generations, Totalvalue, Sourcevalue, Uniqueaccounts, Meandwelltime, Mediandwelltime, Meanmulescore, Elapsedtime, Numnotinvestigated, Parentalertid, Decisiondate, Mostrecentfeedback)
	loggers.Tracenetwork(c.Path(), "folderName", logMessage, Sourcetxnid, Sourcetxntype, Length, Generations, Totalvalue, Sourcevalue, Uniqueaccounts, Meandwelltime, Mediandwelltime, Meanmulescore, Elapsedtime, Numnotinvestigated, Parentalertid, Decisiondate, Mostrecentfeedback)

	return c.JSON(responseBody)

}

//200
//404
//400
//405
//429

//401
//403
