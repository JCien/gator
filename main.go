package main

import (
	"log"
	"os"

	"github.com/JCien/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//err = cfg.SetUser("Jesus")
	programState := &state{
		cfg: &cfg,
	}

	myCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	
	myCommands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = myCommands.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
