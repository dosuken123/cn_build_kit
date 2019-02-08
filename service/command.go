package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (s Service) Clone(args []string) error {
	if s.Src.RepoURL == "" {
		return nil
	}

	depth := ""

	if s.Src.CloneDepth > 0 {
		depth = fmt.Sprintf(" --depth %d", s.Src.CloneDepth)
	}

	cmdName := fmt.Sprintf("sudo --user %s git clone %s %s %s", s.GetUserName(), depth, s.Src.RepoURL, s.GetSrcDir())

	if err := s.ExecuteCommandWithLog("clone", cmdName); err != nil {
		return err
	}

	return nil
}

func (s Service) Pull(args []string) error {
	if s.Src.RepoURL == "" {
		return nil
	}

	cmdName := fmt.Sprintf("sudo --user %s %s", s.GetUserName(), "git checkout master")
	if err := s.ExecuteCommandWithLog("pull", cmdName); err != nil {
		return err
	}

	cmdName = fmt.Sprintf("sudo --user %s %s", s.GetUserName(), "git pull origin master")
	if err := s.ExecuteCommandWithLog("pull", cmdName); err != nil {
		return err
	}

	return nil
}

func (s Service) Clean(args []string) error {
	var targetDirs []string
	if len(args) > 0 && args[0] == "all" {
		targetDirs = append(targetDirs, s.GetServiceDir())
	} else {
		targetDirs = append(targetDirs, s.GetSrcDir(), s.GetCacheDir(), s.GetDataDir(), s.GetLogDir())
	}

	for _, targetDir := range targetDirs {
		os.RemoveAll(targetDir)
	}

	return nil
}

func (s Service) AddScript(args []string) error {
	parsedArgs := parseArgSet(args)

	if _, err := os.Stat(s.GetScriptDir()); os.IsNotExist(err) {
		os.MkdirAll(s.GetScriptDir(), os.ModePerm)
	}

	scriptPath := filepath.Join(s.GetScriptDir(), parsedArgs["name"])
	emptyFile, err := os.Create(scriptPath)
	if err != nil {
		log.Fatal(err)
	}
	emptyFile.Close()

	if err := os.Chmod(scriptPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Chown(scriptPath, s.GetUIDInt(), s.GetGIDInt()); err != nil {
		log.Fatal(err)
	}

	if val, ok := parsedArgs["script"]; ok {
		ioutil.WriteFile(scriptPath, []byte(val), os.ModePerm)
	}

	return nil
}

func (s Service) AddExample(args []string) error {
	parsedArgs := parseArgSet(args)

	cmdName := fmt.Sprintf("sudo --user %s mkdir -p --mode %s %s", s.GetUserName(), "777", s.GetExampleDir())
	if err := s.ExecuteCommandWithLog("add_example", cmdName); err != nil {
		return err
	}

	examplePath := filepath.Join(s.GetExampleDir(), parsedArgs["name"])

	cmdName = fmt.Sprintf("sudo --user %s touch %s", s.GetUserName(), examplePath)
	if err := s.ExecuteCommandWithLog("add_example", cmdName); err != nil {
		return err
	}

	cmdName = fmt.Sprintf("sudo --user %s chmod %s %s", s.GetUserName(), "777", examplePath)
	if err := s.ExecuteCommandWithLog("add_example", cmdName); err != nil {
		return err
	}

	return nil
}

func parseArgSet(args []string) map[string]string {
	m := make(map[string]string)

	for _, arg := range args {
		s := strings.Split(arg, "=")
		key := strings.Replace(s[0], "--", "", -1)
		m[key] = s[1]
	}

	return m
}
