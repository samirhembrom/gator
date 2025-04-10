package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/samirhembrom/blogaggregator/internal/config"
)

type state struct {
	Cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	Handler map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login handler expects a single argument, the username\n")
	}

	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return errors.New("error setting username\n")
	}
	fmt.Printf("Username: %s is set.\n", cmd.Args[0])

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Handler[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("No state provided")
	}
	if len(cmd.Name) == 0 {
		return errors.New("No command name provided")
	}
	if f, ok := c.Handler[cmd.Name]; ok {
		err := f(s, cmd)
		if err != nil {
			return err
		}

	} else {
		return errors.New("No function exist")
	}
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)
	s := &state{
		Cfg: &cfg,
	}

	cmds := &commands{
		Handler: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Print("Not enough args")
		os.Exit(1)
	}

	cmdName := args[1]

	cmdArgs := args[2:]

	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}
	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %v\n", cfg)
}
