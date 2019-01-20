package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dosuken123/cn_build_kit/config"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("Arguments are not enough", nil)
	}

	argTarget := args[1]
	argCommand := args[2]
	argArgs := args[3:]
	// fmt.Printf("Arguments: Command: %+v, Target: %+v\n", argCommand, argTarget)

	services, error := config.ResolveTargetName(argTarget)
	// fmt.Printf("Resolved services: %+v\n", services)

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
