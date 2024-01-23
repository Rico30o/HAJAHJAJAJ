package routes

import (
	"fmt"
	"instapay/db"
	"instapay/model"
	errorhandling "instapay/model/errorHandling"
	"time"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_alertnetwork int

const allowedRates_alertnetwork = 5

func checkRateLimits_alertnetwork() bool {
	requestCounts_alertnetwork++
	if requestCounts_alertnetwork > allowedRates_alertnetwork {
		requestCounts_alertnetwork = 0
		return true
	}
	return false
}

func Alertnetwork(c *fiber.Ctx) error {
	network := &model.NetworkBody{}

	if checkRateLimits_alertnetwork() {
		//429
		return errorhandling.Rate_Limit_Exceeded(c, "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS")
	}

	if c.Method() != fiber.MethodPost {

		//405
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")

	}

	if parsErr := c.BodyParser(&network); parsErr != nil {

		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	} else if network == nil || (network.Since == "" && network.Limit == 0 && network.PaginationToken == "" && network.Filter == "" && !network.Include_all_alerts) {

		//400
		return errorhandling.Bad_Request(c, "The request body is empty")
	}

	// Initialize models
	account := model.TransactionAlert{}
	transaction := &model.Alertnetwork{}
	networks := &model.NetworkResponse{}

	// Fetch data from the database
	if err := db.DB.Debug().Raw(`SELECT * FROM public.transactionAlerts`).Scan(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}
	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_alerts`).Scan(transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}
	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_alert`).Scan(networks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}

	// Initialize an empty slice to store the retrieved alerts
	var transactions []model.TransactionAlert

	// Construct the SQL query with date filtering and order by id
	query := `SELECT * FROM public.transactionAlerts WHERE 1 = 1`
	countQuery := "SELECT COUNT(*) FROM public.transactionAlerts WHERE 1=1"
	countFilteredQuery := "SELECT COUNT(*) FROM public.transactionAlerts WHERE 1=1"

	// Add conditions based on the request body
	if network.Since != "" {
		// Validate the date format
		if _, err := time.Parse("2006-01-02", network.Since); err != nil {
			//400
			return errorhandling.Bad_Request(c, "Please provide a valid date format (YYYY-MM-DD)")
		}

		query += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
		countFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
	}

	if network.Filter != "" {
		query += " AND " + network.Filter
		countFilteredQuery += " AND " + network.Filter
	}

	// Execute the query and retrieve the count for all transactions
	var totalCount int64
	if err := db.DB.Raw(countQuery).Count(&totalCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting all transactions",
		})
	}

	// Execute the query and retrieve the count for filtered transactions
	var filteredCount int64
	if network.Since != "" {
		if err := db.DB.Raw(countFilteredQuery, network.Since).Count(&filteredCount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	} else {
		if err := db.DB.Raw(countFilteredQuery).Count(&filteredCount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	}

	// Limit
	if network.Limit > 0 {
		if int64(network.Limit) > filteredCount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		query += fmt.Sprintf(" LIMIT %d", network.Limit)
	}

	// Execute the main query
	err := db.DB.Debug().Raw(query, network.Since).Scan(&transactions).Error

	if err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	// Check if no data was found for the given date
	if len(transactions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}
	//-----------------------------------------------------------------------------------
	// Initialize an empty slice to store the retrieved alerts
	var transactions1 []model.Alertnetwork

	// Construct the SQL query with date filtering and order by id
	query1 := `SELECT * FROM public.trace_alerts WHERE 1 = 1`
	countQuery1 := "SELECT COUNT(*) FROM public.trace_alerts WHERE 1=1"
	countFilteredQuery1 := "SELECT COUNT(*) FROM public.trace_alerts WHERE 1=1"

	// Add conditions based on the request body
	if network.Since != "" {
		// Validate the date format
		if _, err := time.Parse("2006-01-02", network.Since); err != nil {
			//400
			return errorhandling.Bad_Request(c, "Please provide a valid date format (YYYY-MM-DD)")
		}

		query1 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
		countFilteredQuery1 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
	}

	if network.Filter != "" {
		query1 += " AND " + network.Filter
		countFilteredQuery1 += " AND " + network.Filter
	}

	// Execute the query and retrieve the count for all transactions
	var totalCount1 int64
	if err := db.DB.Raw(countQuery1).Count(&totalCount1).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting all transactions",
		})
	}

	// Execute the query and retrieve the count for filtered transactions
	var filteredCount1 int64
	if network.Since != "" {
		if err := db.DB.Raw(countFilteredQuery1, network.Since).Count(&filteredCount1).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	} else {
		if err := db.DB.Raw(countFilteredQuery1).Count(&filteredCount1).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	}

	// Limit
	if network.Limit > 0 {
		if int64(network.Limit) > filteredCount1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		query1 += fmt.Sprintf(" LIMIT %d", network.Limit)
	}

	// Execute the main query
	err1 := db.DB.Debug().Raw(query1, network.Since).Scan(&transactions1).Error

	if err1 != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	// Check if no data was found for the given date
	if len(transactions1) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}

	//-----------------------------------------------------------------------------------
	// Initialize an empty slice to store the retrieved alerts
	var transactions2 []model.NetworkResponse

	// Construct the SQL query with date filtering and order by id
	query2 := `SELECT * FROM public.trace_alert WHERE 1 = 1`
	countQuery2 := "SELECT COUNT(*) FROM public.trace_alert WHERE 1=1"
	countFilteredQuery2 := "SELECT COUNT(*) FROM public.trace_alert WHERE 1=1"

	// Add conditions based on the request body
	if network.Since != "" {
		// Validate the date format
		if _, err := time.Parse("2006-01-02", network.Since); err != nil {
			//400
			return errorhandling.Bad_Request(c, "Please provide a valid date format (YYYY-MM-DD)")
		}

		query2 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
		countFilteredQuery2 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
	}

	if network.Filter != "" {
		query1 += " AND " + network.Filter
		countFilteredQuery2 += " AND " + network.Filter
	}

	// Execute the query and retrieve the count for all transactions
	var totalCount2 int64
	if err := db.DB.Raw(countQuery2).Count(&totalCount2).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting all transactions",
		})
	}

	// Execute the query and retrieve the count for filtered transactions
	var filteredCount2 int64
	if network.Since != "" {
		if err := db.DB.Raw(countFilteredQuery2, network.Since).Count(&filteredCount2).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	} else {
		if err := db.DB.Raw(countFilteredQuery2).Count(&filteredCount2).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	}

	// Limit
	if network.Limit > 0 {
		if int64(network.Limit) > filteredCount2 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		query1 += fmt.Sprintf(" LIMIT %d", network.Limit)
	}

	// Execute the main query
	err2 := db.DB.Debug().Raw(query2, network.Since).Scan(&transactions2).Error

	if err2 != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")

	}

	// Check if no data was found for the given date
	if len(transactions1) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}

	// Prepare response
	type AlertnetworkResponse struct {
		Alerts            fiber.Map              `json:"alerts"`
		TransactionAlerts model.Alertnetwork     `json:"transactionAlerts"`
		AccountAlerts     model.TransactionAlert `json:"accountAlerts"`
		Network           model.NetworkResponse  `json:"network"`
	}
	alerts := fiber.Map{
		"totalCount":    totalCount,
		"filteredCount": filteredCount,
		"id":            transactions[0].ID,
		"Time":          transactions1[0].Time,
		"networkID":     transactions1[0].Networkalertid,
	}

	response := AlertnetworkResponse{
		Alerts:            alerts,
		TransactionAlerts: transactions1[0],
		AccountAlerts:     transactions[0],
		Network:           transactions2[0],
	}

	return c.JSON(response)
}

//404/
//200/
//400/
//405/
//429/

//401
//403
