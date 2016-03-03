package commands

import (
	"fmt"
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

	if cmd != CommandInit && !offline {
		if _, err := NewPull(branch).Run(); err != nil {
			log.Fatal(err)
		}
	}

	var output []byte
	var err error
	switch cmd {
	case CommandInit:
		output, err = runInit(branch, c)
	}

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", output)
	if !offline {
		if _, err := NewPush(branch).Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func runInit(branch string, c *cli.Context) ([]byte, error) {
	_, err := NewInit(branch).Run()
	output := fmt.Sprintf("FastForward branch '%s' successfully initialized.", branch)
	return []byte(output), err
}
