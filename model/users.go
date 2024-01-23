package model

import (
	"encoding/xml"
)

type User struct {
	Name     string `xml:"name"`
	Lastname string `xml:"lastname"`
	Address  string `xml:"address"`
	Contact  string `xml:"contact"`
	Email    string `xml:"email"`
	Password string `xml:"password"`
}

type UsersInfo struct {
	CorporateID      string `xml:"corporate"`
	BranchID         string `xml:"branchid"`
	TransactionKey   string `xml:"transactionkey"`
	RequestRefNo     string `xml:"requestno"`
	TransactionType  string `xml:"transactiontype"`
	RequestTimeStamp string `xml:"requesttimestamp"`
	TerminalID       string `xml:"terminalid"`
	Address          string `xml:"address"`
}

type Deped struct {
	XMLName       xml.Name `xml:"DepedHead"`
	School        string   `xml:"DepedSchool"`
	PrincipalHead PrincipalHead
}

type PrincipalHead struct {
	XMLName     xml.Name `xml:"soapenv:Header"`
	Principal   string   `xml:"principal"`
	TeacherBody TeacherBody
}

type TeacherBody struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Teacher string   `xml:"teacher"`
	Student Student
}

type Student struct {
	XMLName      xml.Name `xml:"Deped:StudentInformation"`
	StudentName  string   `xml:"studentname"`
	StudentID    string   `xml:"studentid"`
	Section      string   `xml:"section"`
	MajorSubject string   `xml:"majorsubject"`
}

// type Employee struct {
// 	ID     int    `xml:"id"`
// 	Name   string `xml:"name"`
// 	Salary int    `xml:"salary"`
// }

// type User1 struct {
// 	ID       int    `xml:"primaryKey"`
// 	Name     string `xml:"unique"`
// 	Password string
// 	// Add other fields as needed
// }

type LoginResponse struct {
	XMLName xml.Name `json:"response"`
	Message string   `json:"message"`
}

// type IustomerInfo struct {
// 	ResponseCode        string `json:"responseCode"`
// 	Description         string `json:"description"`
// 	CifNo               string `json:"cifNo"`
// 	FilterType          string `json:"filterType"`
// 	MeetingDay          string `json:"meetingDay"`
// 	CustomerInformation CustomerInformation
// }

type CustomerInformation struct {
	GivenName        string `json:"givenName"`
	FirstName        string `json:"firstName"`
	MiddleName       string `json:"middleName"`
	LastName         string `json:"lastName"`
	FirstSpouseName  string `json:"firstSpouseName"`
	MiddleSpouseName string `json:"middleSpouseName"`
	LastSpouseName   string `json:"lastSpouseName"`
	PlaceOfBirth     string `json:"placeOfBirth"`
	DateOfBirth      string `json:"dateOfBirth"`
	Gender           string `json:"gender"`
	ContactDate      string `json:"contactDate"`
	District         string `json:"dqistrict"`
	City             string `json:"city"`
	Province         string `json:"province"`
	Country          string `json:"country"`
	PostCode         string `json:"postCode"`
	MemberClass      string `json:"memberClass"`
	MemberFlag       string `json:"memberFlag"`
	Suffix           string `json:"suffix"`
	CustomerDetail   CustomerDetail
}

type CustomerDetail struct {
	MotherMaidenName string `json:"motherMaidenName"`
	Sector           string `json:"sector"`
	Phone            string `json:"phone"`
	Mobile           string `json:"mobile"`
	Email            string `json:"email"`
	LegalId          string `json:"legalId"`
	LegalDocName     string `json:"legalDocName"`
	CoCode           string `json:"coCode"`
	UnitCode         string `json:"unitCode"`
	CenterCode       string `json:"centerCode"`
	PpiScore         string `json:"ppiScore"`
	PpiDate          string `json:"ppiDate"`
	CustomerKtp      CustomerKtp
}

type CustomerKtp struct {
	Address string `json:"address"`
	ExpDate string `json:"expDate"`
}

type IustomerInfo struct {
	ReferenceNumber string `json:"referenceNumber"`
	FilterType      string `json:"filterType"`
	AccountNo       string `json:"accountNo"`
}

var (
	// Port ...
	Port string
	// SSL
	SSL string
	// Environment
	Environment string
	// SSLCertificate ...
	SSLCertificate string
	// SSLKey ...
	SSLKey string
	// SSLSIgning
	SSLSigning string
	// SSLSigningKey
	SSLSigningKey string
	// RBI PUBKEY
	RBI_PublicKey string
	// Mastercard PUBKEY
	MC_PublicKey string
	// Mastercard Signing
	MC_SSLSigning string
	// For EncryptionDecryption
	SecretKey string
	// POSTGRESQL CONFIG
	PostgreSQLHost    string
	PostgresPort      string
	PostgresSSLMode   string
	PostgresTimeZone  string
	DatabaseName      string
	PostgreslUsername string
	PostgreslPassword string
	//iGATE
	IGateBaseUrl string
)
