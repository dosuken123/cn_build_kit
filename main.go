package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/dosuken123/cn_build_kit/service"
)

func main() {
	// NOTE: In order to correctly set file permission with Mkdir commands,
	// overwrite umask value.
	syscall.Umask(0)

	args := os.Args

	if len(args) < 3 {
		log.Fatal("Arguments are not enough", nil)
	}

	argTarget := args[1]
	argCommand := args[2]
	argArgs := args[3:]

	services, err := service.ResolveTargetName(argTarget)

	if err != nil {
		log.Fatal(err)
	}

	chans := make([]chan bool, len(services))
	for i, srv := range services {
		chans[i] = make(chan bool)

		go func(s service.Service, command string, args []string, c chan bool) {
			err := srv.ExecuteCommand(command, args)
			if err != nil {
				log.Printf("[ERROR] Command execution failure. Service: %v, Command: %v, Error: %v\n", s.Name, command, err)
				os.Exit(1)
			}
			c <- true
		}(srv, argCommand, argArgs, chans[i])
	}

	for i := range services {
		<-chans[i]
	}
	fmt.Printf("DONE\n")
}
