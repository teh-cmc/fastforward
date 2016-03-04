package git

// -----------------------------------------------------------------------------

// Pull implements the `git pull` command.
type Pull struct {
	branch string
}

// NewPull returns a new `Pull` command.
func NewPull(branch string) *Pull { return &Pull{branch: branch} }

// -----------------------------------------------------------------------------

// AllowAutoCheckout always returns `false`.
func (c Pull) AllowAutoCheckout() bool { return false }

// Input always returns `nil`.
func (p Pull) Input() []byte { return nil }

// Command returns a `git pull` command.
func (p Pull) Command() []string { return []string{"pull", "origin", p.branch} }

// Transform does nothing.
func (p Pull) Transform(output []byte) []byte { return output }
