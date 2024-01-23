package routes

import (
	"fmt"
	"instapay/db"
	"instapay/middleware/loggers"
	"instapay/model"
	errorhandling "instapay/model/errorHandling"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_feedback int

const allowedRates_feedback = 5

func checkRateLimits_feedback() bool {
	requestCounts_feedback++
	if requestCounts_feedback > allowedRates_feedback {
		requestCounts_feedback = 0
		return true
	}
	return false
}

func Feedback(c *fiber.Ctx) error {
	var userRequest model.RequestBody

	if checkRateLimits_feedback() {

		//429
		return errorhandling.Rate_Limit_Exceeded(c, "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS")
	}

	if c.Method() != fiber.MethodPost {

		//405
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")

	}

	// Check if the media type is supported (JSON)
	if !c.Is("json") {
		//415
		return errorhandling.Unsupported_Media_Type(c, "The request media type 'application/x-www-form-urlencoded' is not supported by this resource")
	}

	accountArray := &model.RequestBodyArray{}
	// Parse the request body
	if parsErr := c.BodyParser(accountArray); parsErr != nil {
		//422
		return errorhandling.Unprocessable_Entity(c, "Unprocessable_Entity")
	}

	if err := c.BodyParser(&userRequest); err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request body is expecting an array")
	}

	// query := ` SELECT id, alertid, alerttype, entityid, decisiondate, feedback, status, feedbackid FROM public.feedback WHERE 1 = 1`
	query := ` SELECT * FROM public.feedback WHERE 1 = 1`

	if userRequest.Alertid != "" && userRequest.Alertid != "0" {
		query += " AND alertid = ?"
	}

	// Add a condition to check for a specific alerttype
	if userRequest.Alerttype != "" && userRequest.Alerttype != "0" {
		query += " AND alerttype = ?"
	}

	// Add a condition to check for a specific alerttype
	if userRequest.Entityid != "" && userRequest.Entityid != "0" {
		query += " AND entityid = ?"
	}

	// Add a condition to check for a specific alerttype
	if userRequest.Feedback != "" && userRequest.Feedback != "0" {
		query += " AND feedback = ?"
	}

	var feedbackList []model.Response

	// result := db.DB.Raw(query+" ORDER BY id", userRequest.Alertid).Scan(&feedbackList)
	result := db.DB.Raw(query+" ORDER BY id", userRequest.Alertid, userRequest.Alerttype, userRequest.Entityid, userRequest.Feedback).Scan(&feedbackList)
	if result.Error != nil {
		//409
		return errorhandling.Conflict(c, "Conflict")
	}

	if len(feedbackList) == 0 {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	// Check the status field and respond with a specific error if it's true
	if userRequest.Status {
		//403
		return errorhandling.Permision_Denied(c, "Permision_Denied")
	}

	//------------------- logs -------------------
	Feedbackid := feedbackList[0].Feedbackid
	logMessage := fmt.Sprintf("%s:", Feedbackid)
	loggers.Feedbacklogs(c.Path(), "folderName", logMessage)

	return c.JSON(fiber.Map{
		"feedbackID": Feedbackid,
	})

}

// func GenerateReferenceID(instructionID string) string {
// 	var ipsReference string
// 	for ctr := 1; ctr <= 6; ctr++ {
// 		ipsReference = instructionID[len(instructionID)-ctr:]
// 	}
// 	return fmt.Sprintf("IFT%v-%v", ipsReference, time.Now().UnixMilli())
// }

// func GenerateIBFTReferenceID(instructionID string) string {
// 	var ipsReference string
// 	for ctr := 1; ctr <= 6; ctr++ {
// 		ipsReference = instructionID[len(instructionID)-ctr:]
// 	}
// 	return fmt.Sprintf("%v", ipsReference)
// }

//400
//403
//404
//200
//405
//409
//429
//422
//415

//401
