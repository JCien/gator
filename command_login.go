package main

import (
	"log"
	"fmt")

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) <= 1 {
		log.Fatal("login requires one argument, username")
	}

	username := cmd.args[1]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("The user has been set")
	return nil
}
