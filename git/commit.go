package git

// -----------------------------------------------------------------------------

// Commit implements the `git commit` command.
type Commit struct {
	branch string
	input  []byte
}

// NewCommit returns a new `Commit` command.
func NewCommit(branch string, input []byte) *Commit {
	return &Commit{branch: branch, input: input}
}

// -----------------------------------------------------------------------------

// AllowAutoCheckout always returns true.
func (c Commit) AllowAutoCheckout() bool { return true }

// Input returns the associated `c.input`.
func (c Commit) Input() []byte { return c.input }

// Command returns a `git commit` command.
func (c Commit) Command() []string { return []string{"commit", "--allow-empty", "--file", "-"} }

// Transform does nothing.
func (c Commit) Transform(output []byte) []byte { return output }
