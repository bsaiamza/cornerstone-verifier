package log

import (
	"io"
	"log"
	"os"
)

var (
	// Info ...
	Info *log.Logger
	// Error ...
	Error *log.Logger
	// Warning ...
	Warning *log.Logger

	// ServerInfo ...
	ServerInfo *log.Logger
	// ServerError ...
	ServerError *log.Logger
	// ServerWarning ...
	ServerWarning *log.Logger
)

func init() {
	verifierLogFile, err := os.OpenFile("cornerstone_verifier.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Print("Error: Failed to create cornerstone_verifier log file:", err)
	}
	serverLogFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Print("Error: Failed to create server log file:", err)
	}

	vmw := io.MultiWriter(os.Stdout, verifierLogFile)
	smw := io.MultiWriter(os.Stdout, serverLogFile)

	Info = log.New(vmw, "[INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(vmw, "[ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(vmw, "[WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)

	ServerInfo = log.New(smw, "[INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	ServerError = log.New(smw, "[ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	ServerWarning = log.New(smw, "[WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)
}
