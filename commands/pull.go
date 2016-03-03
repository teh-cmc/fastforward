package commands

import "github.com/teh-cmc/fastforward/git"

// -----------------------------------------------------------------------------

// Pull implements the `ff pull` command.
type Pull struct {
	branch string
}

// NewPull returns a new Pull.
func NewPull(branch string) *Pull { return &Pull{branch: branch} }

// -----------------------------------------------------------------------------

// Pull implements the `ff pull` command.
func (p Pull) Run() ([]byte, error) { return git.Run(git.NewPull(p.branch)) }
