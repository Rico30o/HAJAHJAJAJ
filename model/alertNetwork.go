package model

import (
	"time"
)

type (
	NetworkBody struct {
		Since              string `json:"since"`
		Limit              int    `json:"limit"`
		PaginationToken    string `json:"pagination_token"`
		Filter             string `json:"filter"`
		Include_all_alerts bool   `json:"include_all_alerts"`
	}

	Networkadditional struct {
		TotalRecords            int    `json:"totalRecords"`
		Displayedrecords        int    `json:"displayedRecords"`
		Nextpaginationtoken     string `json:"nextpaginationtoken"`
		Previouspaginationtoken string `json:"previouspaginationtoken"`
	}
	AdditionalTransaction struct {
		Transaction_Response []Transaction_Response `json:"transaction_response"`
	}
	AdditonalAccount struct {
		Transaction_Response []Transaction_Response `json:"response_alert"`
	}
	NetworkResponse struct {
		Vizurl             string    `json:"vizurl"`
		Sourcetxnid        string    `json:"sourcetxnid"`
		Sourcetxntype      string    `json:"sourcetxntype"`
		Length             int       `json:"length"`
		Generations        int       `json:"generations"`
		Totalvalue         int       ` json:"totalvalue"`
		Sourcevalue        int       ` json:"sourcevalue"`
		Uniqueaccounts     int       `json:"uniqueaccounts"`
		Meandwelltime      string    `json:"meandwelltime"`
		Mediandwelltime    string    `json:"mediandwelltime"`
		Meanmulescore      float64   ` json:"meanmulescore"`
		Elapsedtime        string    `json:"elapsedtime"`
		Numactionedmules   int       `json:"numactionedmules"`
		Numlegitimate      int       `json:"numlegitimate"`
		Numnotinvestigated int       `json:"numnotinvestigated"`
		Parentalertid      string    `json:"parentalertid"`
		Decisiondate       time.Time `json:"decisiondate"`
		Mostrecentfeedback string    `json:"mostrecentfeedback"`
	}

	TransactionAlert struct {
		ID                 string    `json:"id"`
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
		Value              int       `json:"value"`
		Remitinfo          string    `json:"remitinfo"`
		Generation         int       `json:"generation"`
		Currency           string    `json:"currency"`
		Service            string    `json:"service"`
		Dwelltime          string    `json:"dwelltime"`
		Tracetype          string    `json:"tracetype"`
		Mulescore          float64   `json:"mulescore"`
		Parentalertid      string    `json:"parentalertid"`
		Decisiondate       time.Time `json:"decisiondate"`
		Mostrecentfeedback string    `json:"mostrecentfeedback"`
	}

	Alertnetwork struct {
		ID                           string    `json:"id"`
		Networkalertid               string    `json:"networkalertid"`
		Accountid                    string    `json:"accountid"`
		Networkid                    string    `json:"networkid"`
		Owningbankid                 string    `json:"owningbankid"`
		Owningbankname               string    `json:"owningbankname"`
		Time                         time.Time `json:"time"`
		Name                         string    `json:"name"`
		Mulescore                    float64   `json:"mulescore"`
		Sourcetransactionvalue       int       `json:"sourcetransactionValue"`
		Endpointflag                 bool      `json:"endpointflag"`
		Numoutboundrelationships     int       `json:"numoutboundrelationships"`
		Numinboundrelationships      int       `json:"numinboundrelationships"`
		Numscheduledmandates         int       `json:"numscheduledmandates"`
		Firstappearance              time.Time `json:"firstappearance"`
		Mostrecentappearance         time.Time `json:"mostrecentappearance"`
		Firsttransactiontime         time.Time `json:"firsttransactiontime"`
		Mostrecenttransactiontime    time.Time `json:"mostrecenttransactiontime"`
		Receivessalary               bool      `json:"receivessalary"`
		Dwelltime                    string    `json:"dwelltime"`
		Numnetworks                  int       `json:"numNetworks"`
		Numtracednetworks            int       `json:"numtracednetworks"`
		Generation                   int       `json:"generation"`
		Tracetype                    string    `json:"tracetype"`
		TotalsuspiciousvalueInbound  int       `json:"totalsuspiciousvalueinbound"`
		Totalsuspiciousvalueoutbound int       `json:"totalsuspiciousvalueoutbound"`
		Totalvalueionbound           int       `json:"totalvalueinbound"`
		Totalvalueoutbound           int       `json:"totalvalueoutbound"`
		Generations                  []int     `json:"generations"`
		Mostrecentfeedback           string    `json:"mostrecentfeedback"`
		Parentalertid                string    `json:"parentalertid"`
		Decisiondate                 time.Time `json:"decisionDate"`
	}
)
