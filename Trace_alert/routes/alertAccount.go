package routes

import (
	"fmt"
	"instapay/db"
	"instapay/middleware/loggers"
	"instapay/model"
	errorhandling "instapay/model/errorHandling"
	"time"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_alertaccount int

const allowedRates_alertaccount = 5

func checkRateLimits_alertaccount() bool {
	requestCounts_alertaccount++
	if requestCounts_alertaccount > allowedRates_alertaccount {
		requestCounts_alertaccount = 0
		return true
	}
	return false
}
func Alertsaccount(c *fiber.Ctx) error {
	var userRequest model.RequestBodyalert

	if checkRateLimits_alertaccount() {
		//429
		return errorhandling.Rate_Limit_Exceeded(c, "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS")
	}

	if err := c.BodyParser(&userRequest); err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	// Initialize an empty slice to store the retrieved alerts
	var transactions []model.Alert

	Query := `SELECT * FROM public.trace_alerts WHERE 1 = 1`
	CountQuery := "SELECT COUNT(*) FROM public.trace_alerts WHERE 1=1"
	CountFilteredQuery := "SELECT COUNT(*) FROM public.trace_alerts WHERE 1=1"

	// Add conditions based on the request body
	if userRequest.Since != "" {
		Query += " AND DATE_TRUNC('day', decisiondate) = $1"
		CountFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = $1"
	}

	if userRequest.Filter != "" {
		Query += " AND " + userRequest.Filter
		CountFilteredQuery += " AND " + userRequest.Filter
	}

	// var totalCount int64
	var filteredCount int64

	var count int64
	resultCount := db.DB.Raw(CountQuery).Count(&count)
	if resultCount.Error != nil {
		//400
		return errorhandling.Bad_Request(c, "Error counting transactions")
	}

	// Execute the query and retrieve the count for filtered transactions
	resultFilteredCount := db.DB.Raw(CountFilteredQuery, userRequest.Since).Count(&filteredCount)
	if resultFilteredCount.Error != nil {
		//400
		return errorhandling.Bad_Request(c, "Error counting filtered transactions")
	}

	////limit count and display data

	if userRequest.Limit > 0 {
		if int64(userRequest.Limit) > filteredCount {
			//400
			return errorhandling.Bad_Request(c, "Limit exceeds the total number of filtered transactions")
		}
		Query += fmt.Sprintf(" LIMIT %d", userRequest.Limit)
	}

	// Execute the query and retrieve the results
	err := db.DB.Debug().Raw(Query, userRequest.Since).Scan(&transactions).Error

	if err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	if c.Method() != fiber.MethodPost {
		//405
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	// Check if no data was found for the given date
	if len(transactions) == 0 {
		//400
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")

	}

	var trans []model.AlertResponse

	//--------------------------- logs ---------------------------
	ID := transactions[0].ID
	Networkalertid := transactions[0].Networkalertid
	Accountid := transactions[0].Accountid
	Networkid := transactions[0].Networkid
	Owningbankid := transactions[0].Owningbankid
	Owningbankname := transactions[0].Owningbankname

	logMessage := fmt.Sprintf("%s:%s %s %s %s %s", ID, Networkalertid, Accountid, Networkid, Owningbankid, Owningbankname)
	loggers.Alertaccount(c.Path(), "folderName", logMessage, Networkalertid, Accountid, Networkid, Owningbankid, Owningbankname)

	//---------------------------

	Referenceid := Referenceid(ID)
	Generateid := Generateid(ID)
	type Alert struct {
		Alerts               fiber.Map             `json:"alerts"`
		Transactions         []model.Alert         `json:"transactions"`
		TransactionResponses []model.AlertResponse `json:"trans"`
	}

	response := Alert{
		Alerts: fiber.Map{
			"totalRecords":            count,
			"displayedRecords":        filteredCount,
			"nextPaginationToken":     transactions[0].Nextpaginationtoken,
			"previousPaginationToken": transactions[0].Previouspaginationtoken,
			"referenceID":             Referenceid,
			"generateID":              Generateid,
		},

		Transactions:         transactions,
		TransactionResponses: trans,
	}

	return c.JSON(response)
}

func Referenceid(instructionID string) string {
	var ipsReference string
	for ctr := 1; ctr <= 10; ctr++ {
		ipsReference = instructionID[len(instructionID)-ctr:]
	}
	return fmt.Sprintf("IFT%v-%v", ipsReference, time.Now().UnixMilli())
}

func Generateid(instructionID string) string {
	var ipsReference string
	for ctr := 1; ctr <= 10; ctr++ {
		ipsReference = instructionID[len(instructionID)-ctr:]
	}
	return fmt.Sprintf("%v", ipsReference)
}

//404
//400
//200
//429
//405

//401
//403
