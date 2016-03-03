package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

// -----------------------------------------------------------------------------

// CommandType is an enum that represents available `ff` commands.
type CommandType string

const (
	// CommandTypePull represents the `ff pull` command.
	CommandTypePull CommandType = "pull"
	// CommandTypePush represents the `ff push` command.
	CommandTypePush CommandType = "push"
	// CommandTypeInit represents the `ff init` command.
	CommandTypeInit CommandType = "init"
	// CommandTypeTaskNew represents the `ff task new` command.
	CommandTypeTaskNew CommandType = "task:new"
	// CommandTypeTaskList represents the `ff task list` command.
	CommandTypeTaskList CommandType = "task:list"
)

// Command exposes methods to run a `ff` command.
type Command interface {
	AllowAutoPulling() bool
	AllowAutoPushing() bool
	Run(branch string) ([]byte, error)
}

// NewCommand returns a new Command of the given type `t`.
func NewCommand(t CommandType) Command {
	switch t {
	case CommandTypePull:
		return NewPull()
	case CommandTypePush:
		return NewPush()
	case CommandTypeInit:
		return NewInit()
	default:
		log.Fatalf("'%s': command not supported")
	}
	return nil
}

// -----------------------------------------------------------------------------

// Run runs the given command in the specified context `c`.
func Run(cmd Command, c *cli.Context) {
	branch := c.GlobalString("branch")
	offline := c.GlobalBool("offline")

	if !offline && cmd.AllowAutoPulling() {
		Run(NewPull(), c)
	}

	output, err := cmd.Run(branch)
	if err != nil {
		os.Exit(1)
	}
	if len(output) > 0 {
		fmt.Printf("[FastForward] %s\n", output)
	}

	if cmd.AllowAutoPushing() {
		Run(NewPush(), c)
	}
}
