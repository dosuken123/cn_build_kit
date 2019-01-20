package config

import (
	"log"
	"sync"

	"github.com/dosuken123/cn_build_kit/command"
)

func (s Service) ExecuteCommand(commandName string, wg *sync.WaitGroup) {
	switch commandName {
	case "clone":
		if s.Src.RepoURL != "" {
			c := command.Clone{Url: s.Src.RepoURL, Dir: s.GetSrcDir(), Depth: s.Src.CloneDepth}
			c.Execute()
		}
	case "clean":
		c := command.Clean{ServiceDir: s.GetServiceDir()}
		c.Execute()
	default:
		log.Fatal("Command Not Found", nil)
	}
	wg.Done()
}
