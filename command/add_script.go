package command

import (
	"log"
	"os"
	"path/filepath"
)

type AddScript struct {
	FileDir  string
	FileName string
}

func (a AddScript) Execute() {
	if _, err := os.Stat(a.FileDir); os.IsNotExist(err) {
		os.MkdirAll(a.FileDir, os.ModePerm)
	}

	scriptPath := filepath.Join(a.FileDir, a.FileName)
	emptyFile, err := os.Create(scriptPath)
	if err != nil {
		log.Fatal(err)
	}
	emptyFile.Close()

	if err := os.Chmod(scriptPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
