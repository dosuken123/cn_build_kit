package command

import (
	"fmt"

	"github.com/libgit2/git2go"
)

type Clone struct {
	Url   string
	Dir   string
	Depth int
}

func (c Clone) Execute() {
	fmt.Println("Cloning!!", c.Url)
	fmt.Println("Cloning!!", c.Dir)

	repo, err := git.Clone(c.Url, c.Dir, &git.CloneOptions{})

	fmt.Println("repo: %+v", repo)

	if err != nil {
		panic(err)
	}
}
