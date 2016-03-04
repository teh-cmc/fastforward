package commands

import "github.com/teh-cmc/fastforward/git"

// -----------------------------------------------------------------------------

// Push implements the `ff push` command.
type Push struct{}

// NewPush returns a new Push.
func NewPush() *Push { return &Push{} }

// -----------------------------------------------------------------------------

// AllowAutoPulling always returns `true`.
func (p Push) AllowAutoPulling() bool { return true }

// AllowAutoPushing always returns `false`.
func (p Push) AllowAutoPushing() bool { return false }

// Push implements the `ff push` command.
func (p Push) Run(branch string) ([]byte, error) {
	return git.Run(git.NewPush(branch), branch)
}
