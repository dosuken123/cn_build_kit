package config

import (
	"log"
	"sync"

	"github.com/dosuken123/cn_build_kit/command"
)

func (s Service) ExecuteCommand(commandName string, args []string, wg *sync.WaitGroup) {
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
		log.Fatal("Command Not Found", nil)
	}
	wg.Done()
}
