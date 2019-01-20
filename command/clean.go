package command

import "os"

type Clean struct {
	ServiceDir string
}

func (c Clean) Execute() {
	os.RemoveAll(c.ServiceDir)
}
