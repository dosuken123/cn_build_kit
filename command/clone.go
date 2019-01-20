package command

import (
	"fmt"
	"os/exec"
	"strconv"
)

type Clone struct {
	Url   string
	Dir   string
	Depth int
}

func (c Clone) Execute() {
	fmt.Printf("Cloning URL: %+v, Dir: %+v, Depth: %d\n", c.Url, c.Dir, c.Depth)

	var cmd *exec.Cmd

	if c.Depth > 0 {
		cmd = exec.Command("git", "clone", c.Url, c.Dir, "--depth", strconv.Itoa(c.Depth))
	} else {
		cmd = exec.Command("git", "clone", c.Url, c.Dir)
	}

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("stdout: ", stdout)
}
