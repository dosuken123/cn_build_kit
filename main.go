package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dosuken123/cn_build_kit/service"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("Arguments are not enough", nil)
	}

	argTarget := args[1]
	argCommand := args[2]
	argArgs := args[3:]

	services, error := service.ResolveTargetName(argTarget)

	if error != nil {
		log.Fatal(error)
	}

	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		go service.ExecuteCommand(argCommand, argArgs, &wg)
	}
	wg.Wait()
	fmt.Printf("DONE\n")
}
