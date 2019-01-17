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

	argTarget := args[1]
	argCommand := args[2]

	fmt.Println("Command: ", argCommand)
	fmt.Println("Targetss: ", argTarget)

	targets, error := ResolveTargetName(argTarget)

	fmt.Printf("targets: %+v\n", targets)

	if error != nil {
		log.Fatal(error)
	}

	for _, target := range targets {
		fmt.Println("Target name is ", target.Name)
		fmt.Printf("target: %+v\n", target)
		target.GenerateMeta()
		fmt.Printf("target: %+v\n", target)
		switch argCommand {
		case "clone":
			c := command.Clone{Url: target.Src.RepoURL, Dir: target.Meta.SrcPath, Depth: target.Src.CloneDepth}
			c.Execute()
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
