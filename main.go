package myherotoolkit

import "os"

const (
	LOG_SUCCESS = "SUCCESS"
	LOG_FAILED  = "FAILED"
)

func WriteLog(logMessage string, logType string) error {
	text := []byte(logMessage)

	err := os.WriteFile("./mylog.log", text, 0644)
	if err != nil {
		return err
	}

	return nil
}
