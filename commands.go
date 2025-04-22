package main

import ("github.com/JCien/gator/internal/config"

	"errors")

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	//This method registers a new handler function for a command name
	c.commandList = map[string]func(*state, command) error {
		name: f,
	}
}

func (c *commands) run(s *state, cmd command) error {
	//This method runs a given command with the provided state if it exists
	if commandName, ok := c.commandList[cmd.name]; ok {
		err := commandName(s, cmd)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Command does not exist")
	}

	return nil
}
