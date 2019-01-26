package command

import (
	"fmt"
	"os/exec"
)

type Pull struct {
	SrcDir string
	Remote string
	Branch string
}

func (p Pull) Execute() {
	var cmd *exec.Cmd
	var stdout []byte

	cmd = exec.Command("git", "checkout", p.Branch)
	cmd.Dir = p.SrcDir
	stdout, _ = cmd.Output()

	fmt.Printf("stdout: %s\n", stdout)

	cmd = exec.Command("git", "pull", p.Remote, p.Branch)
	cmd.Dir = p.SrcDir
	stdout, _ = cmd.Output()

	fmt.Printf("stdout: %s\n", stdout)
}
