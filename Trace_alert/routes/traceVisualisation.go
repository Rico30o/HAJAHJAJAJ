package routes

import (
	"fmt"
	"instapay/db"
	"instapay/middleware/loggers"
	"instapay/model"
	errorhandling "instapay/model/errorHandling"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_tracevisual int

const allowedRates_tracevisual = 5

func checkRateLimits_tracevisual() bool {
	requestCounts_tracevisual++
	if requestCounts_tracevisual > allowedRates_tracevisual {
		requestCounts_tracevisual = 0
		return true
	}
	return false
}

func NetworkAlertID(c *fiber.Ctx) error {
	var userRequest model.Tracevisualisationsalert

	// cert := util.LoadCertificate(envRouting.SSLSigning)
	if checkRateLimits_tracevisual() {
		return c.Status(fiber.StatusTooManyRequests).JSON(model.ErrorResponses{
			Errors: struct {
				Error []model.ErrorDetail `json:"Error"`
			}{
				Error: []model.ErrorDetail{{
					Source:      "Gateway",
					ReasonCode:  "RATE_LIMIT_EXCEEDED",
					Description: "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS",
					Recoverable: true,
				}},
			},
		})
	}

	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "TRACE_FINANCIAL_CRIME",
						"ReasonCode":  "METHOD_NOT_ALLOWED",
						"Description": "Only POST method allowed",
						"Recoverable": false,
						"Details":     nil,
					},
				},
			},
		})
	}

	accountArray := &model.RequestBodyArray{}
	// Parse the request body
	if parsErr := c.BodyParser(accountArray); parsErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(model.ErrorResponses{
			Errors: struct {
				Error []model.ErrorDetail `json:"Error"`
			}{
				Error: []model.ErrorDetail{{
					Source:      "FEEDBACK_FINANCIAL_CRIME",
					ReasonCode:  "UNPROCESSABLE_ENTITY",
					Description: "Expects a single JSON object and not an array",
					Recoverable: false,
				}},
			},
		})
	}

	if err := c.BodyParser(&userRequest); err != nil {

		return errorhandling.Bad_Request(c, "The request body is expecting an array")
	}

	query := ` SELECT * FROM public.trace_visualisation WHERE 1 = 1`

	if userRequest.Networkalertid != "" && userRequest.Networkalertid != "0" {
		query += " AND networkalertid = ?"
	}

	// Add a condition to check for a specific alerttype
	if userRequest.Format != "" && userRequest.Format != "0" {
		query += " AND format = ?"
	}

	// Add a condition to check for a specific alerttype
	if userRequest.Type != "" && userRequest.Type != "0" {
		query += " AND type = ?"
	}

	// Add a condition to check for a specific alerttype
	if userRequest.Colourmode != "" && userRequest.Colourmode != "0" {
		query += " AND colourmode = ?"
	}

	var feedbackList []model.Traceresponse

	// result := db.DB.Raw(query+" ORDER BY id", userRequest.Alertid).Scan(&feedbackList)
	result := db.DB.Raw(query+" ORDER BY id", userRequest.Networkalertid, userRequest.Format, userRequest.Type, userRequest.Colourmode).Scan(&feedbackList)
	if result.Error != nil {

		return errorhandling.Conflict(c, "Alert ID does not match the specified entity")

	}

	if userRequest.Status {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"Errors": map[string]interface{}{
				"Error": []map[string]interface{}{
					{
						"Source":      "Gateway",
						"ReasonCode":  "PERMISSION_DENIED",
						"Description": "Invalid customer for third party",
						"Recoverable": false,
						"Details":     nil,
					},
				},
			},
		})
	}

	if len(feedbackList) == 0 {
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}
	//------------------- logs -------------------
	Traceid := feedbackList[0].Traceid
	logMessage := fmt.Sprintf("%s: ", Traceid)
	loggers.Tracevisuallogs(c.Path(), "folderName", logMessage)

	return c.JSON(feedbackList)

}

//429/
//400/
//405/
//500
//403/
//404/
//200/

//401/
