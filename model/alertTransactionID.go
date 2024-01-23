package model

import "time"

// GetTransacInfo
type (
	TransacAlertRequest struct {
		TransactionAlertID string `json:"transactionalertid"`
	}

	ResponseTransac struct {
		ID                 string    ` json:"id"`
		Txnid              string    `json:"txnid"`
		Networkalertid     string    ` json:"networkalertid"`
		Networkid          string    `json:"networkid"`
		Time               time.Time `json:"time"`
		Txntime            time.Time `json:"txntime"`
		Sourceid           string    `json:"sourceid"`
		Destid             string    `json:"destid"`
		Sourcebankid       string    `json:"sourcebankid"`
		Sourcebankname     string    `json:"sourcebankname"`
		Destbankid         string    `json:"destbankid"`
		Destbankname       string    `json:"destbankname"`
		Value              int       `json:"value"`
		Remitinfo          string    `json:"remitinfo"`
		Generation         int       `json:"generation"`
		Currency           string    ` json:"currency"`
		Service            string    ` json:"service"`
		Dwelltime          string    ` json:"dwelltime"`
		Tracetype          string    ` json:"tracetype"`
		Mulescore          float64   `json:"mulescore"`
		Parentalertid      string    ` json:"parentalertid"`
		Decisiondate       time.Time `json:"decisiondate"`
		Mostrecentfeedback string    `json:"mostrecentfeedback"`
	}
)
