package commands

import "github.com/teh-cmc/fastforward/git"

// -----------------------------------------------------------------------------

// Push implements the `ff push` command.
//
// Push implements the `git.Command` interface.
type Push struct {
	branch string
}

// NewPush returns a new Push.
func NewPush(branch string) *Push { return &Push{branch: branch} }

// -----------------------------------------------------------------------------

// Template always returns `nil`.
func (i Push) Template() []byte { return nil }

// Command returns a command that pushes the FastForward branch.
func (i Push) Command() []string { return []string{"push", "origin", i.branch} }

// Push implements the `fwd push` command.
func (i Push) Run() ([]byte, error) { return git.Run(i) }
