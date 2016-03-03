package commands

import "github.com/teh-cmc/fastforward/git"

// -----------------------------------------------------------------------------

// Push implements the `ff push` command.
type Push struct {
	branch string
}

// NewPush returns a new Push.
func NewPush(branch string) *Push { return &Push{branch: branch} }

// -----------------------------------------------------------------------------

// Push implements the `ff push` command.
func (p Push) Run() ([]byte, error) { return git.Run(git.NewPush(p.branch)) }
