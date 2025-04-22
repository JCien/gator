package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JCien/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//err = cfg.SetUser("Jesus")
	newState := state{}
	newState.cfg = &cfg

	myCommands := commands{
		commandList: make(map[string]func(*state, command) error),
	}
	
	myCommands.register("login", handlerLogin)

	argsMain := os.Args
	if len(argsMain) < 2 {
		log.Fatal("Not enough arguments provided")
	}

	commandArgs := argsMain[1:]

	commandHandler := &command{
		name: commandArgs[0],
		args: commandArgs,
	}

	err = myCommands.run(&newState, *commandHandler)
	if err != nil {
		fmt.Print(err)
	}

}
