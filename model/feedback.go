package model

import (
	"time"
)

// /////enhancement feedback

type RequestBody struct {
	Alertid      string    `json:"alertid"`
	Alerttype    string    `json:"alerttype"`
	Entityid     string    `json:"entityid"`
	Decisiondate time.Time `json:"decisiondate"`
	Feedback     string    `json:"feedback"`
	Status       bool      `json:"status"`
}

type Response struct {
	Feedbackid string `json:"feedbackid"`
}

type ErrorDetail struct {
	Source      string      `json:"Source"`
	ReasonCode  string      `json:"ReasonCode"`
	Description string      `json:"Description"`
	Recoverable bool        `json:"Recoverable"`
	Details     interface{} `json:"Details,omitempty"`
}

type ErrorResponses struct {
	Errors struct {
		Error []ErrorDetail `json:"Error"`
	} `json:"Errors"`
}

type RequestBodyArray struct {
	Items []RequestBody `json:"items"`
}
