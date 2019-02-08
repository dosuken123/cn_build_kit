package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (s Service) Clone(args []string) {
	if s.Src.RepoURL == "" {
		return
	}

	cmdName := fmt.Sprintf("git clone %s %s", s.Src.RepoURL, s.GetSrcDir())

	if s.Src.CloneDepth > 0 {
		cmdName += fmt.Sprintf(" --depth %d", s.Src.CloneDepth)
	}

	s.ExecuteCommandWithLog("clone", cmdName)
}

func (s Service) Pull(args []string) {
	if s.Src.RepoURL == "" {
		return
	}

	cmdName := fmt.Sprintf("sudo --user %s %s", s.GetUserName(), "git checkout master")

	s.ExecuteCommandWithLog("pull", cmdName)

	cmdName = fmt.Sprintf("sudo --user %s %s", s.GetUserName(), "git pull origin master")

	s.ExecuteCommandWithLog("pull", cmdName)
}

func (s Service) Clean(args []string) {
	var targetDirs []string
	if len(args) > 0 && args[0] == "all" {
		targetDirs = append(targetDirs, s.GetServiceDir())
	} else {
		targetDirs = append(targetDirs, s.GetSrcDir(), s.GetCacheDir(), s.GetDataDir(), s.GetLogDir())
	}

	for _, targetDir := range targetDirs {
		os.RemoveAll(targetDir)
	}
}

func (s Service) AddScript(args []string) {
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
}

func (s Service) AddExample(args []string) {
	parsedArgs := parseArgSet(args)

	cmdName := fmt.Sprintf("sudo --user %s mkdir -p --mode %s %s", s.GetUserName(), "777", s.GetExampleDir())

	s.ExecuteCommandWithLog("add_example", cmdName)

	examplePath := filepath.Join(s.GetExampleDir(), parsedArgs["name"])

	cmdName = fmt.Sprintf("sudo --user %s touch %s", s.GetUserName(), examplePath)

	s.ExecuteCommandWithLog("add_example", cmdName)

	cmdName = fmt.Sprintf("sudo --user %s chmod %s %s", s.GetUserName(), "777", examplePath)

	s.ExecuteCommandWithLog("add_example", cmdName)
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
