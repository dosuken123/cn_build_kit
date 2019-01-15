package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dosuken123/cn_build_kit/command"
)

func main() {
	// config := GetConfig()
	// fmt.Println("%+v", config)

	args := os.Args

	if len(args) < 3 {
		log.Fatal("Arguments are not enough", nil)
	}

	argCommand := args[1]
	argTarget := args[2]

	fmt.Println("Command: ", argCommand)
	fmt.Println("Targetss: ", argTarget)

	targets, error := ResolveTargetName(argTarget)

	if error != nil {
		log.Fatal(error)
	}

	for _, target := range targets {
		fmt.Println("Target name is ", target.Name)
		switch argCommand {
		case "clone":
			g := command.Clone{}
			g.Execute()
		default:
			log.Fatal("Command Not Found", nil)
		}
	}

	// for _, element := range conf.Services {
	// 	fmt.Println("Name: ", element.Name)
	// 	fmt.Println("Host: ", element.Host)
	// }

	// g := git.Clone{Url: "aa", Depth: 1}
	// g.Execute()
}
