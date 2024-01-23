package model

import "time"

type (
	TransacIDmatches struct {
		Service         string    `json:"service"`
		Sourceid        string    `json:"sourceid"`
		Beneficiaryid   string    `json:"beneficiaryid"`
		Transactiontime time.Time `json:"transactiontime"`
		Amount          int64     `json:"amount"`
	}

	ResponseIDmatches struct {
		Matchid       string    `json:"matchid"`
		Sourceid      string    `json:"sourceid"`
		Beneficiaryid string    `json:"beneficiaryid"`
		Reftime       time.Time `json:"refTime"`
		Amount        int64     `json:"amount"`
	}
)
