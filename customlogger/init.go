package customlogger

import (
	"log"
	"os"
)

var (
	InfoLogger      *log.Logger
	DuplicateLogger *log.Logger
	OrphanLogger    *log.Logger
	ErrorLogger     *log.Logger
)

func init() {
	if _, err := os.Stat("./logs/"); os.IsNotExist(err) {
		if err := os.Mkdir("./logs/", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	infoFile, err := os.OpenFile("./logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	duplicateFile, err := os.OpenFile("./logs/duplicates.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orphanFile, err := os.OpenFile("./logs/orphan-comments.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	errorFile, err := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	DuplicateLogger = log.New(duplicateFile, "DUPLICATE: ", log.Lmsgprefix)
	OrphanLogger = log.New(orphanFile, "OPRPHAN: ", log.Lmsgprefix)
	ErrorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
