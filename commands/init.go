package commands

import (
	"fmt"

	"github.com/teh-cmc/fastforward/git"
)

// -----------------------------------------------------------------------------

// Init implements the `ff init` command.
//
// Init implements the `git.Command` interface.
type Init struct {
	branch string
}

// NewInit returns a new Init.
func NewInit(branch string) *Init { return &Init{branch: branch} }

// -----------------------------------------------------------------------------

// Template always returns `nil`.
func (i Init) Template() []byte { return nil }

// Command returns a command that creates a new git branch.
func (i Init) Command() []string { return []string{"checkout", "-b", i.branch} }

// Init implements the `fwd init` command.
func (i Init) Run() ([]byte, error) {
	_, err := git.Run(i)
	output := fmt.Sprintf("FastForward branch '%s' successfully initialized.", i.branch)
	return []byte(output), err
}
