package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (s Service) ExecuteCommand(commandName string, args []string) error {
	if error := s.EnsureUser(); error != nil {
		log.Fatal(error)
	}

	s.CleanLog(commandName)

	if ok, err := s.executeCustomCommand(commandName, args); ok {
		return err
	}

	if ok, err := s.executeDefaultCommand(commandName, args); ok {
		return err
	}

	fmt.Printf("[WARN] Command was not found. Service Name: %v, commandName: %v\n", s.Name, commandName)
	return nil
}

func (s Service) executeCustomCommand(commandName string, args []string) (bool, error) {
	scriptPath := filepath.Join(s.GetScriptDir(), commandName)

	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return false, nil
	}

	cmdName := fmt.Sprintf("sudo --preserve-env --user %s %s", s.GetUserName(), scriptPath)
	return true, s.ExecuteCommandWithLog(commandName, cmdName)
}

func (s Service) executeDefaultCommand(commandName string, args []string) (bool, error) {
	switch commandName {
	case "clone":
		return true, s.Clone(args)
	case "clean":
		return true, s.Clean(args)
	case "pull":
		return true, s.Pull(args)
	case "add_script":
		return true, s.AddScript(args)
	case "add_example":
		return true, s.AddExample(args)
	}
	return false, nil
}
