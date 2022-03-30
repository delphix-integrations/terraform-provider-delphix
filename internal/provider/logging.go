package provider

import (
	"log"
	"os"
)

var (
	WarnLog  *log.Logger
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {

	ErrorLog = log.New(os.Stderr, "[DELPHIX] [ERROR] ", log.LstdFlags)
	WarnLog = log.New(os.Stdout, "[DELPHIX] [WARN] ", log.LstdFlags)
	InfoLog = log.New(os.Stdout, "[DELPHIX] [INFO] ", log.LstdFlags)

}
