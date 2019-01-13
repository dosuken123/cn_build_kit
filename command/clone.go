package command

import "fmt"

type Clone struct {
	Url   string
	Depth int
}

func (c Clone) Execute() {
	fmt.Println("Cloning!!", c.Url)
}
