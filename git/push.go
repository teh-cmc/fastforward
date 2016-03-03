package git

// -----------------------------------------------------------------------------

// Push implements the `git push` Command.
type Push struct {
	branch string
}

// NewPush returns a new Push Command.
func NewPush(branch string) *Push { return &Push{branch: branch} }

// -----------------------------------------------------------------------------

// Template always returns `nil`.
func (p Push) Template() []byte { return nil }

// Command returns a command that pushes `p.branch`.
func (p Push) Command() []string { return []string{"push", "origin", p.branch} }
