package util

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"instapay/db"
	"instapay/model/bah"
	"io/ioutil"
)

func GetServiceEP(service, environment string) string {
	serviceRoute := &bah.ServiceRoute{}
	db.DB.Raw("SELECT * FROM rbi_instapay.get_service(?,?)", service, environment).Scan(serviceRoute)
	return serviceRoute.ServiceUrl
}

func LoadCertificate(filename string) *x509.Certificate {
	// LOAD CERTIFICATE
	certFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading certificate file:", err.Error())
	}

	// PARSE CERTIFICATE
	block, _ := pem.Decode(certFile)
	if block == nil {
		fmt.Println("Error decoding PEM Block:", block)
	}

	cert, certErr := x509.ParseCertificate(block.Bytes)
	if certErr != nil {
		fmt.Println("Error parsing certificate:", certErr.Error())
	}
	return cert
}
