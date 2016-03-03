package git

// -----------------------------------------------------------------------------

// Checkout implements the `git checkout` command.
type Checkout struct {
	branch string
}

// NewCheckout returns a new Checkout Command.
func NewCheckout(branch string) *Checkout { return &Checkout{branch: branch} }

// -----------------------------------------------------------------------------

// Template always returns `nil`.
func (c Checkout) Template() []byte { return nil }

// Command returns a command that creates a new git branch.
func (c Checkout) Command() []string { return []string{"checkout", "-b", c.branch} }
