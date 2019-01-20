package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/dosuken123/cn_build_kit/command"
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
		log.Fatal(err)
	}
	fmt.Printf("output is %s\n", out)

	return nil
}

func (s Service) ExecuteDefaultCommand(commandName string, args []string) error {
	switch commandName {
	case "clone":
		if s.Src.RepoURL != "" {
			c := command.Clone{Url: s.Src.RepoURL, Dir: s.GetSrcDir(), Depth: s.Src.CloneDepth}
			c.Execute()
		}
	case "clean":
		var c command.Clean
		if len(args) > 0 && args[0] == "all" {
			c = command.Clean{TargetDir: []string{s.GetServiceDir()}}
		} else {
			c = command.Clean{TargetDir: []string{s.GetSrcDir(), s.GetCacheDir(), s.GetDataDir(), s.GetLogDir()}}
		}
		c.Execute()
	case "add_script":
		c := command.AddScript{FileDir: s.GetScriptDir(), FileName: args[0]}
		c.Execute()
	default:
		return errors.New("Default command does not exist")
	}
	return nil
}
