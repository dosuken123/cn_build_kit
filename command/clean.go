package command

import "os"

type Clean struct {
	TargetDir []string
}

func (c Clean) Execute() {
	for _, targetDir := range c.TargetDir {
		os.RemoveAll(targetDir)
	}
}
