package model

import (
	"time"
)

type RequestBodyalert struct {
	Since           string `json:"since"`
	Limit           int    `json:"limit"`
	Paginationtoken string `json:"paginationtoken"`
	Filter          string `json:"filter"`
}

// Response represents the structure of the response

type AlertResponse struct {
	Parentalertid      string    `json:"parentalertid"`
	Decisiondate       time.Time `json:"decisiondate"`
	Mostrecentfeedback string    `json:"mostrecentfeedback"`
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Errors ErrorDetail `json:"Errors"`
}

// ErrorDetail represents the structure of error details
type ErrorDetailalert struct {
	Error []ErrorInfo `json:"Error"`
}

// ErrorInfo represents detailed information about an error
type ErrorInfo struct {
	Source      string `json:"Source"`
	ReasonCode  string `json:"ReasonCode"`
	Description string `json:"Description"`
	Recoverable bool   `json:"Recoverable"`
	Details     string `json:"Details"`
}

type AlertStruct struct {
	TotalRecords            int     `json:"totalRecords"`
	DisplayedRecords        int     `json:"displayedRecords"`
	NextPaginationToken     string  `json:"nextPaginationToken"`
	PreviousPaginationToken string  `json:"previousPaginationToken"`
	Alerts                  []Alert `json:"alerts"`
}

type Alert struct {
	ID                           string    `json:"id"`
	Networkalertid               string    `json:"networkalertid"`
	Accountid                    string    `json:"accountid"`
	Networkid                    string    `json:"networkid"`
	Owningbankid                 string    `json:"owningbankid"`
	Owningbankname               string    `json:"owningbankname"`
	Time                         time.Time `json:"time"`
	Name                         string    `json:"name"`
	Mulescore                    float64   `json:"muleScore"`
	Sourcetransactionvalue       int       `json:"sourcetransactionvalue"`
	Endpointflag                 bool      `json:"endpointflag"`
	Numoutboundrelationships     int       `json:"numoutboundrelationships"`
	NumInboundRelationships      int       `json:"numinboundrelationships"`
	Numscheduledmandates         int       `json:"numscheduledmandates"`
	Firstappearance              time.Time `json:"firstappearance"`
	Mostrecentappearance         time.Time `json:"mostrecentappearance"`
	Firsttransactiontime         time.Time `json:"firsttransactiontime"`
	Mostrecenttransactiontime    time.Time `json:"mostrecenttransactiontime"`
	Receivessalary               bool      `json:"receivessalary"`
	Dwelltime                    string    `json:"dwelltime"`
	Numnetworks                  int       `json:"numnetworks"`
	Numtracednetworks            int       `json:"numtracednetworks"`
	Generation                   int       `json:"generation"`
	Tracetype                    string    `json:"traceType"`
	Totalsuspiciousvalueinbound  int       `json:"totalsuspiciousvalueinbound"`
	Totalsuspiciousvalueoutbound int       `json:"totalsuspiciousvalueoutbound"`
	Totalvalueinbound            int       `json:"totalvalueinbound"`
	Totalvalueoutbound           int       `json:"totalvalueoutbound"`
	Generations                  []int     `json:"generations"`
	Mostrecentfeedback           string    `json:"mostrecentfeedback"`
	Parentalertid                string    `json:"parentalertid"`
	Decisiondate                 time.Time `json:"decisiondate"`
	Nextpaginationtoken          string    `json:"nextpaginationtoken"`
	Previouspaginationtoken      string    `json:"previouspaginationtoken"`
}
