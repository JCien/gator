package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	//This method registers a new handler function for a command name
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	//This method runs a given command with the provided state if it exists
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
			return errors.New("Command does not exist")
	}

	return f(s, cmd)
}
