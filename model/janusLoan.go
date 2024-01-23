package model

type (
	Request_loan struct {
		AvailCount        string `json:"availCount"`
		HaveAvailedBefore string `json:"aaveAvailedBefore"`
		DID               string `json:"dID"`
		LoanProductCode   string `json:"loanProductCode"`
		PaymentMode       string `json:"paymentMode"`
		Term              string `json:"term"`
		Amount            string `json:"amount"`
		PurposeOfLoan     string `json:"purposeOfLoan"`
		Longitude         string `json:"longitude"`
		Latitude          string `json:"latitude"`
		InstiCode         string `json:"instiCode"`
		RequestID         string `json:"requestID"`
		Timestamp         string `json:"timestamp"`
		Name              Name
	}

	Name struct {
		BeneficiaryGender     string `json:"beneficiaryGender"`
		BeneficiaryFirstName  string `json:"beneficiaryFirstName"`
		BeneficiaryMiddleName string `json:"beneficiaryMiddleName"`
		BeneficiaryLastName   string `json:"beneficiaryLastName"`
		BeneficiaryBirthday   string `json:"beneficiaryBirthday"`
		BeneficiaryAge        string `json:"beneficiaryAge"`
		CID                   string `json:"cID"`
		FirstName             string `json:"firstName"`
		LastName              string `json:"lastName"`
		MiddleName            string `json:"middleName"`
		CenterCode            string `json:"centerCode"`
		UnitCode              string `json:"unitCode"`
		BranchCode            string `json:"branchCode"`
		ContactNumber         string `json:"contactNumber"`
		BirthDate             string `json:"birthDate"`
	}

	Request struct {
		AvailCount        int    `json:"availcount,omitempty"`
		HaveAvailedBefore string `json:"haveavailedbefore,omitempty"`
		DID               string ` json:"did"`
		LoanProductCode   int    `json:"loanProductCode"`
		PaymentMode       int    `json:"paymentMode"`
		Term              int    `json:"term"`
		Amount            int    `json:"amount"`
		PurposeOfLoan     int    ` json:"purposeOfLoan"`
		Longitude         string `json:"longitude"`
		Latitude          string `json:"latitude"`
		InstiCode         int    `json:"instiCode"`
		RequestID         string `json:"requestId,omitempty"`
		Timestamp         string `json:"timestamp,omitempty"`
		CustomerInfo      CustomerInfo
	}

	CustomerInfo struct {
		BeneficiaryGender     string `json:"beneficiaryGender"`
		BeneficiaryFirstName  string `json:"beneficiaryFirstName"`
		BeneficiaryMiddleName string `json:"beneficiaryMiddleName"`
		BeneficiaryLastName   string `json:"beneficiaryLastName"`
		BeneficiaryBirthday   string `json:"beneficiaryBirthday"`
		BeneficiaryAge        string `json:"beneficiaryAge"`
		CID                   string `json:"cid"`
		FirstName             string `json:"firstName"`
		LastName              string `json:"lastName"`
		MiddleName            string `json:"middleName"`
		CenterCode            string `json:"centerCode"`
		UnitCode              string `json:"unitCode"`
		BranchCode            string `json:"branchCode"`
		ContactNumber         string `json:"contactNumber"`
		BirthDate             string `json:"birthDate"`
	}

	Responsemodel struct {
		RetCode string      `json:"retCode"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	ErrorModel struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
		Error     error  `josn:"error "`
	}
)
