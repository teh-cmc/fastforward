package git

// -----------------------------------------------------------------------------

// Push implements the `git push` Command.
type Push struct {
	branch string
}

// NewPush returns a new Push Command.
func NewPush(branch string) *Push { return &Push{branch: branch} }

// -----------------------------------------------------------------------------

// AllowAutoCheckout always returns true.
func (c Push) AllowAutoCheckout() bool { return true }

// Input always returns `nil`.
func (p Push) Input() []byte { return nil }

// Command returns a `git push` command.
func (p Push) Command() []string { return []string{"push", "origin", p.branch} }

// Transform does nothing.
func (p Push) Transform(output []byte) []byte { return output }
