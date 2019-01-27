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
	"sync"
)

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
	fmt.Printf("cmdArgs %+v\n", cmdArgs)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "MY_VAR=1")

	// for key, value := range s.GetVariables() {
	// 	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	// }

	fmt.Printf("aa %+v\n", cmd)

	var wg sync.WaitGroup

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmdReader := io.MultiReader(stdout, stderr)
	reader := bufio.NewReader(cmdReader)

	cmd.Start()

	wg.Add(2)
	go func() {
		defer wg.Done()
		line, _, _ := reader.ReadLine()
		logger.Printf("%s\n", line)
	}()

	go func() {
		defer wg.Done()
		line, _, _ := reader.ReadLine()
		logger.Printf("%s\n", line)
	}()

	wg.Wait()

	cmd.Wait()
}

func copyLogs(r io.Reader, logfn func(args ...interface{})) {
	buf := make([]byte, 80)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			logfn(buf[0:n])
		}
		if err != nil {
			break
		}
	}
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
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal("Failed to open log file ", err)
	}

	return f
}

func initLogger(file *os.File) *log.Logger {
	return log.New(file, "", 0)
}
