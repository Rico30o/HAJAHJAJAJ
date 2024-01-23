package routes

import (
	"fmt"
	"instapay/db"
	"instapay/middleware/loggers"
	"instapay/model"

	"github.com/gofiber/fiber/v2"
)

var requestCountsalert_transaction int

const allowedRatesalert_transaction = 5

func checkRateLimitsalert_transaction() bool {
	requestCountsalert_transaction++
	if requestCountsalert_transaction > allowedRatesalert_transaction {
		requestCountsalert_transaction = 0
		return true
	}
	return false
}

func Alerttransaction(c *fiber.Ctx) error {
	var userRequest model.Transaction_Body

	if checkRateLimitsalert_transaction() {
		return c.Status(fiber.StatusTooManyRequests).JSON(model.ErrorResponses{
			Errors: struct {
				Error []model.ErrorDetail `json:"Error"`
			}{
				Error: []model.ErrorDetail{{
					Source:      "Gateway",
					ReasonCode:  "RATE_LIMIT_EXCEEDED",
					Description: "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS",
					Recoverable: true,
					Details:     nil,
				}},
			},
		})

	}

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": map[string]interface{}{
				"Error": []map[string]interface{}{
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

	// Initialize an empty slice to store the retrieved alerts
	var transactions []model.Transaction_Response

	query := `SELECT * FROM public.alerts_transaction WHERE 1 = 1`
	countQuery := "SELECT COUNT(*) FROM public.alerts_transaction WHERE 1=1"
	countFilteredQuery := "SELECT COUNT(*) FROM public.alerts_transaction WHERE 1=1"

	// Add conditions based on the request body
	if userRequest.Since != "" {
		query += " AND DATE_TRUNC('day', decisiondate) = $1"
		countFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = $1"
	}

	if userRequest.Filter != "" {
		query += " AND " + userRequest.Filter
		countFilteredQuery += " AND " + userRequest.Filter
	}

	query += ` ORDER BY idalert`

	// var totalCount int64
	var filteredCount int64

	var count int64
	resultCount := db.DB.Raw(countQuery).Count(&count)
	if resultCount.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error counting transactions",
		})
	}

	// Execute the query and retrieve the count for filtered transactions
	resultFilteredCount := db.DB.Raw(countFilteredQuery, userRequest.Since).Count(&filteredCount)
	if resultFilteredCount.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error counting filtered transactions",
		})
	}

	////limit count and display data

	if userRequest.Limit > 0 {
		if int64(userRequest.Limit) > filteredCount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		query += fmt.Sprintf(" LIMIT %d", userRequest.Limit)
	}

	// Execute the query and retrieve the results
	err := db.DB.Debug().Raw(query, userRequest.Since).Scan(&transactions).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": map[string]interface{}{
				"Error": []map[string]interface{}{
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

	// Check if no data was found for the given date
	if len(transactions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}

	var trans model.AlertResponse

	// Return the transactions as JSON response

	Count := count
	ID := transactions[0].ID
	Txnid := transactions[0].Txnid
	Networkalertid := transactions[0].Networkalertid
	Networkid := transactions[0].Networkid
	Time := transactions[0].Time
	Txntime := transactions[0].Txntime
	Sourceid := transactions[0].Sourceid
	Destid := transactions[0].Destid
	Sourcebankid := transactions[0].Sourcebankid
	Sourcebankname := transactions[0].Sourcebankname
	Destbankid := transactions[0].Destbankid
	Destbankname := transactions[0].Destbankname
	Value := transactions[0].Value
	Remitinfo := transactions[0].Remitinfo
	Generation := transactions[0].Generation
	Currency := transactions[0].Currency
	Service := transactions[0].Service
	Dwelltime := transactions[0].Dwelltime
	Tracetype := transactions[0].Tracetype
	Mulescore := transactions[0].Mulescore
	Parentalertid := trans.Parentalertid
	Decisiondate := trans.Decisiondate
	Mostrecentfeedback := trans.Mostrecentfeedback
	logMessage := fmt.Sprintf("%s: %d %s %s %s %s %s %s %s %s %s %s %s %d %s %d %s %s %s %s %.2f %s %s %s", ID, Count, Txnid, Networkalertid, Networkid, Time, Txntime, Sourceid, Destid, Sourcebankid, Sourcebankname, Destbankid, Destbankname, Value, Remitinfo, Generation, Currency, Service, Dwelltime, Tracetype, Mulescore, Parentalertid, Decisiondate, Mostrecentfeedback)
	loggers.Alerttransaction(c.Path(), "folderName", logMessage, Count, Txnid, Networkalertid, Networkid, Time, Txntime, Sourceid, Destid, Sourcebankid, Sourcebankname, Destbankid, Destbankname, Value, Remitinfo, Generation, Currency, Service, Dwelltime, Tracetype, Mulescore, Parentalertid, Decisiondate, Mostrecentfeedback)

	return c.JSON(transactions)

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

//404
//400
//200
//429
//405

//401
//403

// // Used for balance inquiries and validating the account, it also generates a reference ID.
// func BalanceInquiry(c *fiber.Ctx, accountNumber, instructionID string) igateModel.AccountValidationResponse {
// 	accountInfo := &igateModel.RequestAccountNumber{
// 		AccountNumber: accountNumber,
// 		InstructionID: instructionID,
// 	}

// 	requestAccountNumber, marshalErr := json.Marshal(accountInfo)
// 	if marshalErr != nil {
// 		return igateModel.AccountValidationResponse{
// 			RetCode:      http.StatusBadRequest,
// 			Description:  "marshal error",
// 			ResponseCode: marshalErr.Error(),
// 		}
// 	}

// 	// This will get the endpoint from DB
// 	ServiceEP := util.GetServiceEP("CheckAccount_igate", strings.ToLower(envRouting.Environment))
// 	// fmt.Println("SERVICE:", ServiceEP)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodPost, ServiceEP, bytes.NewBuffer(requestAccountNumber))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Merchant-ID", "QVBJMDAwMDU=")

// 	fmt.Println("REQUEST:", req)
// 	if err != nil {
// 		return igateModel.AccountValidationResponse{
// 			RetCode:      http.StatusBadRequest,
// 			Description:  "request error",
// 			ResponseCode: err.Error(),
// 		}
// 	}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return igateModel.AccountValidationResponse{
// 			RetCode:      http.StatusBadRequest,
// 			Description:  "client response error",
// 			ResponseCode: err.Error(),
// 		}
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return igateModel.AccountValidationResponse{
// 			RetCode:      http.StatusBadRequest,
// 			Description:  "resopnse body error",
// 			ResponseCode: err.Error(),
// 		}
// 	}

// 	fmt.Println(string(body))
// 	response := &igateModel.AccountValidationResponse{}

// 	if unmarshalErr := json.Unmarshal(body, response); unmarshalErr != nil {
// 		return igateModel.AccountValidationResponse{
// 			RetCode:      http.StatusBadRequest,
// 			Description:  "unmarshall error",
// 			ResponseCode: unmarshalErr.Error(),
// 		}
// 	}

// 	referenceId := GenerateReferenceID(instructionID)
// 	loggers.BalanceInquiry(c.Path(), "igate", "Balance_Inquiry", instructionID, response.ResponseCode, response.Description, response.AccountNumber, response.AccountName, response.ProductCode, response.ProductName, referenceId, response.AvailableBalance, response.CurrentBalance)

// 	if response.ResponseCode == "00" { // Valid Account
// 		return igateModel.AccountValidationResponse{
// 			RetCode:      100,
// 			ResponseCode: response.ResponseCode,
// 			Description:  response.Description,
// 			ReferenceID:  GenerateReferenceID(instructionID),
// 		}
// 	} else { // Invalid Account
// 		return igateModel.AccountValidationResponse{
// 			RetCode:      101,
// 			Description:  response.Description,
// 			ResponseCode: response.ResponseCode,
// 		}
// 	}
// }

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

// // POSTMAN
// // Used for balance inquiries and validating the account, it also generates a reference ID.
// func AccountValidation(c *fiber.Ctx) error {
// 	requestAccountNumber := &igateModel.RequestAccountNumber{}
// 	if parsErr := c.BodyParser(requestAccountNumber); parsErr != nil {
// 		return c.JSON(fiber.Map{
// 			"message": "error parsing",
// 			"error":   parsErr.Error(),
// 		})
// 	}

// 	accountNumber, marshalErr := json.Marshal(requestAccountNumber)
// 	if marshalErr != nil {
// 		return c.JSON(fiber.Map{
// 			"message": "marshal error",
// 			"error":   marshalErr.Error(),
// 		})
// 	}

// 	// This will get the endpoint from DB
// 	ServiceEP := util.GetServiceEP("CheckAccount_igate", strings.ToLower(envRouting.Environment))
// 	// fmt.Println("SERVICE:", ServiceEP)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodPost, ServiceEP, bytes.NewBuffer(accountNumber))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Merchant-ID", "QVBJMDAwMDU=")

// 	// fmt.Println("REQUEST:", req)
// 	if err != nil {
// 		return c.JSON(fiber.Map{
// 			"message": "http request error",
// 			"error":   err.Error(),
// 		})
// 	}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return c.JSON(fiber.Map{
// 			"message": "client request error",
// 			"error":   err.Error(),
// 		})
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return c.JSON(fiber.Map{
// 			"message": "reading body error",
// 			"error":   err.Error(),
// 		})
// 	}

// 	// fmt.Println(string(body))
// 	response := &igateModel.AccountValidationResponse{}
// 	if unmarshalErr := json.Unmarshal(body, response); unmarshalErr != nil {
// 		return c.JSON(fiber.Map{
// 			"message": "unmarshal error",
// 			"error":   err.Error(),
// 		})
// 	}

// 	referenceId := GenerateReferenceID(requestAccountNumber.InstructionID)
// 	loggers.BalanceInquiry(c.Path(), "igate", "Balance_Inquiry", requestAccountNumber.InstructionID, response.ResponseCode, response.Description, response.AccountNumber, response.AccountName, response.ProductCode, response.ProductName, referenceId, response.AvailableBalance, response.CurrentBalance)
// 	if response.ResponseCode == "00" { // Valid Account
// 		if requestAccountNumber.InstructionID == "" {
// 			return c.JSON(fiber.Map{
// 				"retCode":      "101",
// 				"responseCode": "A0",
// 				"description":  "instruction id is missing",
// 			})
// 		} else {
// 			return c.JSON(fiber.Map{
// 				"retCode":      "100",
// 				"responseCode": response.ResponseCode,
// 				"description":  response.Description,
// 				"referenceId":  referenceId,
// 				"data":         response,
// 			})
// 		}
// 	} else { // Invalid Account
// 		return c.JSON(fiber.Map{
// 			"retCode":  "101",
// 			"response": response,
// 		})
// 	}
// }
