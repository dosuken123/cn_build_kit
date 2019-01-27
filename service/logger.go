package service

import (
	"log"
	"os"
	"path/filepath"
)

func (s Service) Log(commandName string, text string) {
	logFilePath := filepath.Join(s.GetLogDir(), commandName)
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	logger := log.New(f, "", 0)
	logger.Println(text)
}
