package commands

import (
	"fmt"

	"github.com/teh-cmc/fastforward/git"
)

// -----------------------------------------------------------------------------

// Init implements the `ff init` command.
type Init struct {
	branch string
}

// NewInit returns a new Init.
func NewInit(branch string) *Init { return &Init{branch: branch} }

// -----------------------------------------------------------------------------

// Init implements the `ff init` command.
func (i Init) Run() ([]byte, error) {
	_, err := git.Run(git.NewCheckout(i.branch))
	output := fmt.Sprintf("FastForward branch '%s' successfully initialized.", i.branch)
	return []byte(output), err
}
