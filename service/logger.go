package service

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (s Service) CleanLog(commandName string) {
	logFile := filepath.Join(s.GetLogDir(), commandName)
	os.Remove(logFile)
}

func (s Service) Log(commandName string, text string) {
	logFile := openFile(s.GetLogDir(), commandName)
	defer logFile.Close()

	logger := initLogger(logFile)
	logger.Println(text)
}

func (s Service) ExecuteCommandWithLog(commandName string, script string) {
	logFile := openFile(s.GetLogDir(), commandName)
	defer logFile.Close()

	logger := initLogger(logFile)

	cmdArgs := strings.Fields(script)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	cmd.Env = os.Environ()

	for key, value := range s.GetVariables() {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmdReader := io.MultiReader(stdout, stderr)
	scanner := bufio.NewScanner(cmdReader)

	done := make(chan struct{})

	go func() {
		for scanner.Scan() {
			logger.Printf("%s\n", scanner.Text())
		}

		done <- struct{}{}
	}()

	cmd.Start()

	<-done

	cmd.Wait()
}

func print(reader io.Reader, loggerOut *log.Logger) {
	r := bufio.NewReader(reader)
	line, _, _ := r.ReadLine()
	loggerOut.Printf("%s\n", line)
}

func openFile(logDir string, commandName string) *os.File {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModePerm)
	}

	logFilePath := filepath.Join(logDir, commandName)
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

	if err != nil {
		log.Fatal("Failed to open log file ", err)
	}

	return f
}

func initLogger(file *os.File) *log.Logger {
	return log.New(file, "", 0)
}
