package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// config := GetConfig()
	// fmt.Println("%+v", config)

	args := os.Args
	arg_command := strings.Title(args[1])
	arg_targets := args[2]

	fmt.Println("Command: ", arg_command)
	fmt.Println("Targets: ", arg_targets)

	command := MakeInstance(arg_command)
	command.Execute()

	// for _, element := range conf.Services {
	// 	fmt.Println("Name: ", element.Name)
	// 	fmt.Println("Host: ", element.Host)
	// }

	// g := git.Clone{Url: "aa", Depth: 1}
	// g.Execute()
}
