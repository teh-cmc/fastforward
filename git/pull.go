package git

// -----------------------------------------------------------------------------

// Pull implements the `git pull` Command.
type Pull struct {
	branch string
}

// NewPull returns a new Pull Command.
func NewPull(branch string) *Pull { return &Pull{branch: branch} }

// -----------------------------------------------------------------------------

// Template always returns `nil`.
func (p Pull) Template() []byte { return nil }

// Command returns a command that pulls `p.branch`.
func (p Pull) Command() []string { return []string{"pull", "--rebase", p.branch} }
