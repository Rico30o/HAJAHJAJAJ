package loggers

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	Separator     *log.Logger
)

var InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, err
}

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

// feedback
func Feedbacklogs(class, folder, Feedbackid string) {
	currentTime := time.Now()
	folderName := "./logs/feedback/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/alerts_feedback %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Feedback - - - -   ")
	InfoLogger.Printf("FEEDBACK: %+v\n", Feedbackid)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Feedback - - - -   ")
	log.Printf("FEEDBACK: %+v\n", Feedbackid)
	log.Printf("=========================================================================================")

}

// alerttransaction
func Alertaccount(class, folder, ID string, Networkalertid string, Accountid string, Networkid string, Owningbankid string, Owningbankname string) {
	currentTime := time.Now()
	folderName := "./logs/alert_account/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/alerts_account %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Alert-Transaction - - - -   ")
	InfoLogger.Printf("FEEDBACK: %+v\n", ID)
	InfoLogger.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	InfoLogger.Printf("ACCOUNTID: %+v\n", Accountid)
	InfoLogger.Printf("NETWORKID: %+v\n", Networkid)
	InfoLogger.Printf("OWNINGBANKID: %+v\n", Owningbankid)
	InfoLogger.Printf("OWNINGBANKNAME: %+v\n", Owningbankname)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Alert-Transaction - - - -   ")
	log.Printf("FEEDBACK: %+v\n", ID)
	log.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	log.Printf("ACCOUNTID: %+v\n", Accountid)
	log.Printf("NETWORKID: %+v\n", Networkid)
	log.Printf("OWNINGBANKID: %+v\n", Owningbankid)
	log.Printf("OWNINGBANKNAME: %+v\n", Owningbankname)
	log.Printf("=========================================================================================")

}

// alerttransaction
func Alerttransaction(class, folder, ID string, Count int64, Txnid string, Networkalertid string, Networkid string, Time time.Time, Txntime time.Time, Sourceid string, Destid string, Sourcebankid string, Sourcebankname string, Destbankid string, Destbankname string, Value int, Remitinfo string, Generation int, Currency string, Service string, Dwelltime string, Tracetype string, Mulescore float64, Parentalertid string, Decisiondate time.Time, Mostrecentfeedback string) {
	currentTime := time.Now()
	folderName := "./logs/alert_transaction/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/alert_transaction %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Alert-Transaction - - - -   ")
	InfoLogger.Printf("ID: %+v\n", ID)
	InfoLogger.Printf("COUNT: %+v\n", Count)
	InfoLogger.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	InfoLogger.Printf("NETWORKID: %+v\n", Networkid)
	InfoLogger.Printf("Time: %+v\n", Time)
	InfoLogger.Printf("TXNTIME: %+v\n", Txntime)
	InfoLogger.Printf("SOURCEID: %+v\n", Sourceid)
	InfoLogger.Printf("DESTID: %+v\n", Destid)
	InfoLogger.Printf("SOURCEBANKID: %+v\n", Sourcebankid)
	InfoLogger.Printf("SOURCEBANKNAME: %+v\n", Sourcebankname)
	InfoLogger.Printf("DESTBANKID: %+v\n", Destbankid)
	InfoLogger.Printf("DESTBANKNAME: %+v\n", Destbankname)
	InfoLogger.Printf("VALUE: %+v\n", Value)
	InfoLogger.Printf("REMITINFO: %+v\n", Remitinfo)
	InfoLogger.Printf("GENERATION: %+v\n", Generation)
	InfoLogger.Printf("CURRENCY: %+v\n", Currency)
	InfoLogger.Printf("SERVICE: %+v\n", Service)
	InfoLogger.Printf("DWELLTIME: %+v\n", Dwelltime)
	InfoLogger.Printf("TRACETYPE: %+v\n", Tracetype)
	InfoLogger.Printf("MULESCORE: %+v\n", Mulescore)
	InfoLogger.Printf("PARENTALERTID: %+v\n", Parentalertid)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)
	InfoLogger.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - -  Alert-Transaction - - - -   ")
	log.Printf("ID: %+v\n", ID)
	log.Printf("COUNT: %+v\n", Count)
	log.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	log.Printf("NETWORKID: %+v\n", Networkid)
	log.Printf("TIME: %+v\n", Time)
	log.Printf("TXNTIME: %+v\n", Txntime)
	log.Printf("SOURCEID: %+v\n", Sourceid)
	log.Printf("DESTID: %+v\n", Destid)
	log.Printf("SOURCEBANKID: %+v\n", Sourcebankid)
	log.Printf("SOURCEBANKNAME: %+v\n", Sourcebankname)
	log.Printf("DESTBANKID: %+v\n", Destbankid)
	log.Printf("DESTBANKNAME: %+v\n", Destbankname)
	log.Printf("VALUE: %+v\n", Value)
	log.Printf("REMITINFO: %+v\n", Remitinfo)
	log.Printf("GENERATION: %+v\n", Generation)
	log.Printf("CURRENCY: %+v\n", Currency)
	log.Printf("SERVICE: %+v\n", Service)
	log.Printf("DWELLTIME: %+v\n", Dwelltime)
	log.Printf("TRACETYPE: %+v\n", Tracetype)
	log.Printf("MULESCORE: %+v\n", Mulescore)
	log.Printf("PARENTALERTID: %+v\n", Parentalertid)
	log.Printf("DECISIONDATE: %+v\n", Parentalertid)
	log.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	log.Printf("=========================================================================================")

}

// tracenetwork
func Tracenetwork(class, folder, Vizurl string, Sourcetxnid string, Sourcetxntype string, Length int, Generations int, Totalvalue int, Sourcevalue int, Uniqueaccounts int, Meandwelltime string, Mediandwelltime string, Meanmulescore float64, Elapsedtime string, Numnotinvestigated int, Parentalertid string, Decisiondate time.Time, Mostrecentfeedback string) {
	currentTime := time.Now()
	folderName := "./logs/trace_network/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/trace_network %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Trace-Network - - - -   ")
	InfoLogger.Printf("VIZURL: %+v\n", Vizurl)
	InfoLogger.Printf("SOURCETXNID: %+v\n", Sourcetxnid)
	InfoLogger.Printf("SOURCETXNTYPE: %+v\n", Sourcetxntype)
	InfoLogger.Printf("LENGTH: %+v\n", Length)
	InfoLogger.Printf("GENERATIONS: %+v\n", Generations)
	InfoLogger.Printf("TOTALVALUE: %+v\n", Totalvalue)
	InfoLogger.Printf("SOURCEVALUE: %+v\n", Sourcevalue)
	InfoLogger.Printf("UNIQUEACCOUNTS: %+v\n", Uniqueaccounts)
	InfoLogger.Printf("MEANDWELLTIME: %+v\n", Meandwelltime)
	InfoLogger.Printf("MEDIANDWELLTIME: %+v\n", Mediandwelltime)
	InfoLogger.Printf("MEANMULESCORE: %+v\n", Meanmulescore)
	InfoLogger.Printf("ELAPSEDTIME: %+v\n", Elapsedtime)
	InfoLogger.Printf("NUMNOTINVESTIGATED: %+v\n", Numnotinvestigated)
	InfoLogger.Printf("PARENTALERTID: %+v\n", Parentalertid)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)
	InfoLogger.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Trace-Network - - - -   ")
	log.Printf("VIZURL: %+v\n", Vizurl)
	log.Printf("SOURCETXNID: %+v\n", Sourcetxnid)
	log.Printf("SOURCETXNTYPE: %+v\n", Sourcetxntype)
	log.Printf("LENGTH: %+v\n", Length)
	log.Printf("GENERATIONS: %+v\n", Generations)
	log.Printf("TOTALVALUE: %+v\n", Totalvalue)
	log.Printf("SOURCEVALUE: %+v\n", Sourcevalue)
	log.Printf("UNIQUEACCOUNTS: %+v\n", Uniqueaccounts)
	log.Printf("MEANDWELLTIME: %+v\n", Meandwelltime)
	log.Printf("MEDIANDWELLTIME: %+v\n", Mediandwelltime)
	log.Printf("MEANMULESCORE: %+v\n", Meanmulescore)
	log.Printf("ELAPSEDTIME: %+v\n", Elapsedtime)
	log.Printf("NUMNOTINVESTIGATED: %+v\n", Numnotinvestigated)
	log.Printf("PARENTALERTID: %+v\n", Parentalertid)
	log.Printf("DECISIONDATE: %+v\n", Decisiondate)
	log.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	log.Printf("=========================================================================================")

}

func Tracevisuallogs(class, folder, Traceid string) {
	currentTime := time.Now()
	folderName := "./logs/trace_visualisation/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/trace_visualisation %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Trace-ID - - - -   ")
	InfoLogger.Printf("TRACEID: %+v\n", Traceid)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Trace-ID - - - -   ")
	log.Printf("TRACEID: %+v\n", Traceid)
	log.Printf("=========================================================================================")

}
