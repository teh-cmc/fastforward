package commands

import "github.com/teh-cmc/fastforward/git"

// -----------------------------------------------------------------------------

// Pull implements the `ff pull` command.
type Pull struct{}

// NewPull returns a new Pull.
func NewPull() *Pull { return &Pull{} }

// -----------------------------------------------------------------------------

// AllowAutoPulling always returns `false`.
func (p Pull) AllowAutoPulling() bool { return false }

// AllowAutoPushing always returns `false`.
func (p Pull) AllowAutoPushing() bool { return false }

// Pull implements the `ff pull` command.
func (p Pull) Run(branch string) ([]byte, error) {
	return git.Run(git.NewPull(branch), branch)
}
