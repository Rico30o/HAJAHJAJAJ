package bah

type (
	LocalInstrumentList struct {
		LocalInstrument string `json:"localInstrument"`
		IsEnabled       bool   `json:"isEnabled"`
		Description     string `json:"description"`
	}

	ServiceRoute struct {
		ServiceUrl string `json:"url"`
	}

	StatusFields struct {
		IsSignedOn bool       `json:"isSignedOn"`
		Remarks    string     `json:"remarks"`
		SignedBy   string     `json:"signedBy"`
		Downtime   UpDowntime `json:"downtime"`
		Uptime     UpDowntime `json:"uptime"`
	}

	UpDowntime struct {
		Date string `json:"date"`
		Time string `json:"time"`
	}

	IPSStatus struct {
		SignedOn    bool   `json:"SignedOn"`
		Remarks     string `json:"remarks"`
		SignedBy    string `json:"signedBy"`
		SignonDate  string `json:"signonDate"`
		SignonTime  string `json:"signonTime"`
		SignoffDate string `json:"signoffDate"`
		SignoffTime string `json:"signoffTime"`
	}

	SetupDowntime struct {
		SetupDowntime int `json:"setupDowntime"`
	}
)
