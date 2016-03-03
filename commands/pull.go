package commands

import "github.com/teh-cmc/fastforward/git"

// -----------------------------------------------------------------------------

// Pull implements the `ff pull` command.
//
// Pull implements the `git.Command` interface.
type Pull struct {
	branch string
}

// NewPull returns a new Pull.
func NewPull(branch string) *Pull { return &Pull{branch: branch} }

// -----------------------------------------------------------------------------

// Template always returns `nil`.
func (i Pull) Template() []byte { return nil }

// Command returns a command that pulls the FastForward branch.
func (i Pull) Command() []string { return []string{"pull", "--rebase", i.branch} }

// Pull implements the `fwd pull` command.
func (i Pull) Run() ([]byte, error) { return git.Run(i) }
