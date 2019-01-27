package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func (s Service) ExecuteCommand(commandName string, args []string, wg *sync.WaitGroup) {
	var err error

	err = s.ExecuteCustomCommand(commandName, args)

	if err != nil {
		err = s.ExecuteDefaultCommand(commandName, args)
	}

	if err != nil {
		fmt.Printf("[WARN] Command was not found. Service Name: %v, commandName: %v\n", s.Name, commandName)
	}

	wg.Done()
}

func (s Service) ExecuteCustomCommand(commandName string, args []string) error {
	scriptPath := filepath.Join(s.GetScriptDir(), commandName)

	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return errors.New("Default command does not exist")
	}

	out, err := exec.Command(scriptPath).Output()
	if err != nil {
		log.Fatal("Failed to execute custom command ", err)
	}
	fmt.Printf("output is %s\n", out)

	return nil
}

func (s Service) ExecuteDefaultCommand(commandName string, args []string) error {
	switch commandName {
	case "clone":
		s.Clone(args)
	case "clean":
		s.Clean(args)
	case "pull":
		s.Pull(args)
	case "add_script":
		s.AddScript(args)
	default:
		return errors.New("Default command does not exist")
	}
	return nil
}
