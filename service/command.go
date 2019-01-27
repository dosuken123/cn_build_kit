package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func (s Service) Clone(args []string) {
	if s.Src.RepoURL == "" {
		return
	}

	var cmd *exec.Cmd

	if s.Src.CloneDepth > 0 {
		cmd = exec.Command("git", "clone", s.Src.RepoURL, s.GetSrcDir(),
			"--depth", strconv.Itoa(s.Src.CloneDepth))
	} else {
		cmd = exec.Command("git", "clone", s.Src.RepoURL, s.GetSrcDir())
	}

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("stdout: %s\n", stdout)
}

func (s Service) Pull(args []string) {
	if s.Src.RepoURL == "" {
		return
	}

	var cmd *exec.Cmd
	var stdout []byte

	cmd = exec.Command("git", "checkout", "master")
	cmd.Dir = s.GetSrcDir()
	stdout, _ = cmd.Output()

	fmt.Printf("stdout: %s\n", stdout)

	cmd = exec.Command("git", "pull", "origin", "master")
	cmd.Dir = s.GetSrcDir()
	stdout, _ = cmd.Output()

	fmt.Printf("stdout: %s\n", stdout)
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
	parsedArgs := ParseArgSet(args)

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

	if val, ok := parsedArgs["script"]; ok {
		ioutil.WriteFile(scriptPath, []byte(val), os.ModePerm)
	}
}

func ParseArgSet(args []string) map[string]string {
	m := make(map[string]string)

	for _, arg := range args {
		s := strings.Split(arg, "=")
		key := strings.Replace(s[0], "--", "", -1)
		m[key] = s[1]
	}

	return m
}
