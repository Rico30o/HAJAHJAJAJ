package model

import "time"

//Alert Transaction //

type (
	Transaction_Body struct {
		Since           string `json:"since"`
		Limit           int    `json:"limit"`
		PaginationToken string `json:"pagination_token"`
		Filter          string `json:"filter"`
	}

	Alert_response struct {
		TotalRecords            int    ` json:"totalRecords"`
		DisplayedRecords        int    `json:"displayedRecords"`
		Nextpaginationtoken     string `json:"nextpaginationtoken"`
		Previouspaginationtoken string `json:"previouspaginationtoken"`
	}

	Alertdata struct {
		ID                   string    ` json:"id"`
		Time                 time.Time `json:"time"`
		Networkid            string    `json:"networkid"`
		Transaction_Response Transaction_Response
	}

	Transaction_Response struct {
		ID                 string    ` json:"id"`
		Txnid              string    `json:"txnid"`
		Networkalertid     string    `json:"networkalertid"`
		Networkid          string    `json:"networkid"`
		Time               time.Time `json:"time"`
		Txntime            time.Time `json:"txntime"`
		Sourceid           string    `json:"sourceid"`
		Destid             string    `json:"destid"`
		Sourcebankid       string    `json:"sourcebankid"`
		Sourcebankname     string    `json:"sourcebankname"`
		Destbankid         string    `json:"destbankid"`
		Destbankname       string    `json:"destbankname"`
		Value              int       ` json:"value"`
		Remitinfo          string    `json:"remitinfo"`
		Generation         int       ` json:"generation"`
		Currency           string    `json:"currency"`
		Service            string    `json:"service"`
		Dwelltime          string    `json:"dwelltime"`
		Tracetype          string    `json:"tracetype"`
		Mulescore          float64   `json:"mulescore"`
		Parentalertid      string    `json:"parentalertid"`
		Decisiondate       time.Time `json:"decisiondate"`
		Mostrecentfeedback string    `json:"mostrecentfeedback"`
		Status             bool      `json:"status"`
	}
	// errror//
	APIError struct {
		Source      string `json:"Source"`
		ReasonCode  string `json:"ReasonCode"`
		Description string `json:"Description"`
		Recoverable bool   `json:"Recoverable"`
		Details     string `json:"Details"`
	}
)
