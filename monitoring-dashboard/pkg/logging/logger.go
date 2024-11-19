package logging

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "MONITORING-DASHBOARD: ", log.LstdFlags)
}
