package model

import "time"

type (
	Request_trace struct {
		Txnid         string `json:"txnid"`
		Sourcetxntype string `json:"sourcetxntype"`
	}
	

	Trace_TransactionAlert struct {
		Id             string    `json:"id"`
		Txnid          string    `json:"txnid"`
		Networkalertid string    `json:"networkalertid"`
		Networkid      string    `json:"networkid"`
		Time           time.Time `json:"time"`
		Txntime        time.Time `json:"txntime"`
		Sourceid       string    `json:"sourceid"`
		Destid         string    `json:"destid"`
		Sourcebankid   string    `json:"sourcebankid"`
		Sourcebankname string    `json:"sourcebankname"`
		Destbankid     string    `json:"destbankid"`
		Destbankname   string    `json:"destbankname"`
		Value          int       `json:"value"`
	}

	Trace_AccountAlert struct {
		Id             string    `json:"id"`
		Networkalertid string    `json:"networkalertid"`
		Accountid      string    `json:"accountid"`
		Networkid      string    `json:"networkid"`
		Owningbankid   string    `json:"owningbankid"`
		Owningbankname string    `json:"owningbankname"`
		Time           time.Time `json:"time"`
	}

	Trace_Alert struct {
		Id                 string                   `json:"id"`
		Time               time.Time                `json:"time"`
		Networkid          string                   `json:"networkid"`
		Transactionalerts  []Trace_TransactionAlert `json:"transactionalerts"`
		Accountalerts      []Trace_AccountAlert     `json:"accountalerts"`
		Vizurl             string                   `json:"vizurl"`
		Sourcetxnid        string                   `json:"sourcetxnid"`
		Sourcetxntype      string                   `json:"sourcetxntype"`
		Length             int                      `json:"length"`
		Generations        int                      `json:"generations"`
		Totalvalue         int                      `json:"totalvalue"`
		Sourcevalue        int                      `json:"sourcevalue"`
		Uniqueaccounts     int                      `json:"uniqueaccounts"`
		MeandwellTime      string                   `json:"meandwellTime"`
		MediandwellTime    string                   `json:"mediandwellTime"`
		MeanmuleScore      float64                  `json:"meanmulescore"`
		Elapsedtime        string                   `json:"elapsedtime"`
		Numactionedmules   int                      `json:"numactionedmules"`
		Numlegitimate      int                      `json:"numlegitimate"`
		Numnotinvestigated int                      `json:"numnotinvestigated"`
		Parentalertid      string                   `json:"parentalertid"`
		Decisiondate       time.Time                `json:"decisiondate"`
		Mostrecentfeedback string                   `json:"mostrecentfeedback"`
	}
	Trace_RequestBodyArray struct {
		Items []RequestBody `json:"items"`
	}
)
