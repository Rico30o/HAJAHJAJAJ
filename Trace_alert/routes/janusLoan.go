package routes

import (
	"instapay/db"
	"instapay/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
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

type CustomerInfo struct {
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

func CreateResponse(c *fiber.Ctx) error {
	request := Request{}
	var data = make(map[string]interface{})

	if parseErr := c.BodyParser(&request); parseErr != nil {
		log.Printf("Error parsing request body: %s\n", parseErr)
		data["Message"] = "Error: Body Parsing"
		data["IsSuccess"] = false
		data["Error"] = parseErr

		return c.Status(fiber.StatusBadRequest).JSON(model.Responsemodel{
			RetCode: "400",
			Message: "Bad Request",
			Data: model.ErrorModel{
				Message:   "Error: Body Parsing",
				IsSuccess: false,
				Error:     parseErr,
			},
		})
	}

	log.Printf("Received request: %+v\n", request)

	// Include individual fields of CustomerInfo in the response data.
	data["AvailCount"] = request.AvailCount
	data["HaveAvailedBefore"] = request.HaveAvailedBefore
	data["DID"] = request.DID
	data["LoanProductCode"] = request.LoanProductCode
	data["PaymentMode"] = request.PaymentMode
	data["Term"] = request.Term
	data["Amount"] = request.Amount
	data["PurposeOfLoan"] = request.PurposeOfLoan
	data["Longitude"] = request.Longitude
	data["Latitude"] = request.Latitude
	data["InstiCode"] = request.InstiCode
	data["RequestID"] = request.RequestID
	data["Timestamp"] = request.Timestamp

	// Include individual fields of CustomerInfo in the response data.
	data["BeneficiaryGender"] = request.CustomerInfo.BeneficiaryGender
	data["BeneficiaryFirstName"] = request.CustomerInfo.BeneficiaryFirstName
	data["BeneficiaryMiddleName"] = request.CustomerInfo.BeneficiaryMiddleName
	data["BeneficiaryLastName"] = request.CustomerInfo.BeneficiaryLastName
	data["BeneficiaryBirthday"] = request.CustomerInfo.BeneficiaryBirthday
	data["BeneficiaryAge"] = request.CustomerInfo.BeneficiaryAge
	data["CID"] = request.CustomerInfo.CID
	data["FirstName"] = request.CustomerInfo.FirstName
	data["LastName"] = request.CustomerInfo.LastName
	data["MiddleName"] = request.CustomerInfo.MiddleName
	data["CenterCode"] = request.CustomerInfo.CenterCode
	data["UnitCode"] = request.CustomerInfo.UnitCode
	data["BranchCode"] = request.CustomerInfo.BranchCode
	data["ContactNumber"] = request.CustomerInfo.ContactNumber
	data["BirthDate"] = request.CustomerInfo.BirthDate

	if err := db.DB.Debug().Raw(`
	INSERT INTO request_loan(
		availcount, haveavailedbefore, did, loanproductcode, paymentmode, term, amount, purposeofloan, longitude, latitude, insticode, requestid, "timestamp", beneficiarygender, beneficiaryfirstname, beneficiarymiddlename, beneficiarylastname, beneficiarybirthday, beneficiaryage, cid, firstname, lastname, middlename, centercode, unitcode, branchcode, contactnumber, birthdate)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`,
		request.AvailCount, request.HaveAvailedBefore, request.DID, request.LoanProductCode, request.PaymentMode, request.Term, request.Amount, request.PurposeOfLoan, request.Longitude, request.Latitude, request.InstiCode, request.RequestID, request.Timestamp,
		request.CustomerInfo.BeneficiaryGender, request.CustomerInfo.BeneficiaryFirstName, request.CustomerInfo.BeneficiaryMiddleName, request.CustomerInfo.BeneficiaryLastName, request.CustomerInfo.BeneficiaryBirthday, request.CustomerInfo.BeneficiaryAge, request.CustomerInfo.CID, request.CustomerInfo.FirstName, request.CustomerInfo.LastName, request.CustomerInfo.MiddleName,
		request.CustomerInfo.CenterCode, request.CustomerInfo.UnitCode, request.CustomerInfo.BranchCode, request.CustomerInfo.ContactNumber, request.CustomerInfo.BirthDate).Scan(&data).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Responsemodel{
			RetCode: "400",
			Message: "Bad Request",
			Data: model.ErrorModel{
				Message:   "Error: Database Insert",
				IsSuccess: false,
				Error:     err,
			},
		})
	}

	return c.JSON(data)
}
