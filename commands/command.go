package commands

import (
	"log"

	"github.com/codegangsta/cli"
)

// -----------------------------------------------------------------------------

// Command is an enum that represents every possible `ff` commands.
type Command int

const (
	// CommandPull represents the `ff pull` command.
	CommandPull Command = iota
	// CommandPush represents the `ff push` command.
	CommandPush Command = iota
	// CommandInit represents the `ff init` command.
	CommandInit Command = iota
)

// -----------------------------------------------------------------------------

// Run runs a `Command` with the given context `c`.
func Run(cmd Command, c *cli.Context) {
	branch := c.GlobalString("branch")
	offline := c.GlobalBool("offline")

	var output []byte
	var err error
	switch cmd {
	case CommandPull:
		output, err = NewPull(branch).Run()
	case CommandPush:
		if err := autoPull(branch, offline); err != nil {
			log.Fatal(err)
		}
		output, err = NewPush(branch).Run()
	case CommandInit:
		output, err = NewInit(branch).Run()
	}

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", output)

	if err := autoPush(branch, offline); err != nil {
		log.Fatal(err)
	}
}

func autoPull(branch string, offline bool) error {
	if offline {
		return nil
	}
	_, err := NewPull(branch).Run()
	return err
}

func autoPush(branch string, offline bool) error {
	if offline {
		return nil
	}
	_, err := NewPush(branch).Run()
	return err
}
