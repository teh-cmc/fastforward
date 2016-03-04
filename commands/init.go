package commands

import (
	"fmt"

	"github.com/teh-cmc/fastforward/git"
)

// -----------------------------------------------------------------------------

// Init implements the `ff init` command.
type Init struct{}

// NewInit returns a new Init.
func NewInit() *Init { return &Init{} }

// -----------------------------------------------------------------------------

// AllowAutoPulling always returns `false`.
func (i Init) AllowAutoPulling() bool { return false }

// AllowAutoPushing always returns `true`.
func (i Init) AllowAutoPushing() bool { return true }

// Init implements the `ff init` command.
func (i Init) Run(branch string) ([]byte, error) {
	var output []byte
	var err error

	output, err = git.Run(git.NewBranch(git.BranchTypeCurrent, branch), branch)
	if err != nil {
		return output, err
	}
	current := string(output)
	output, err = git.Run(git.NewBranch(git.BranchTypeNew, branch), branch)
	if err != nil {
		return output, err
	}
	output, err = git.Run(git.NewBranch(git.BranchTypeSwitch, current), current)
	if err != nil {
		return output, err
	}
	output = []byte(fmt.Sprintf("branch '%s' successfully initialized.\n", branch))
	return output, nil
}
