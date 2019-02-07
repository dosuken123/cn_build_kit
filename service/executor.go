package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func (s Service) ExecuteCommand(commandName string, args []string, wg *sync.WaitGroup) {
	var err error

	if error := s.EnsureUser(); error != nil {
		log.Fatal(error)
	}

	s.CleanLog(commandName)

	err = s.executeCustomCommand(commandName, args)

	if err != nil {
		err = s.executeDefaultCommand(commandName, args)
	}

	if err != nil {
		fmt.Printf("[WARN] Command was not found. Service Name: %v, commandName: %v\n", s.Name, commandName)
	}

	wg.Done()
}

func (s Service) executeCustomCommand(commandName string, args []string) error {
	scriptPath := filepath.Join(s.GetScriptDir(), commandName)

	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return errors.New("Default command does not exist")
	}

	s.ExecuteCommandWithLog(commandName, scriptPath)

	return nil
}

func (s Service) executeDefaultCommand(commandName string, args []string) error {
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
