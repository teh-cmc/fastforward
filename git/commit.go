package git

import "github.com/teh-cmc/fastforward"

// -----------------------------------------------------------------------------

// Commit implements the `git commit` command.
type Commit struct {
	branch string
	msg    *forward.CommitMessage
}

// NewCommit returns a new `Commit` command.
func NewCommit(branch string, msg *forward.CommitMessage) *Commit {
	return &Commit{branch: branch, msg: msg}
}

// -----------------------------------------------------------------------------

// AllowAutoCheckout always returns true.
func (c Commit) AllowAutoCheckout() bool { return true }

// Input returns the associated commit message.
func (c Commit) Input() []byte { return c.msg.Bytes() }

// Command returns a `git commit` command.
func (c Commit) Command() []string { return []string{"commit", "--allow-empty", "--file", "-"} }

// Transform does nothing.
func (c Commit) Transform(output []byte) []byte { return output }
