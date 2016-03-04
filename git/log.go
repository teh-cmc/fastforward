package git

import "fmt"

// -----------------------------------------------------------------------------

// Log implements the `git log` Command.
type Log struct {
	branch  string
	pattern string
}

// NewLog returns a new Log Command.
func NewLog(branch, pattern string) *Log {
	return &Log{branch: branch, pattern: pattern}
}

// -----------------------------------------------------------------------------

// AllowAutoCheckout always returns `false`.
func (l Log) AllowAutoCheckout() bool { return false }

// Input always returns `nil`.
func (l Log) Input() []byte { return nil }

// Command returns a `git log` command.
func (l Log) Command() []string {
	return []string{
		"log", l.branch,
		"--format='%aN\xff%aE\xff%at\xff%s\xff%b'",
		fmt.Sprintf("--grep='%s'", l.pattern),
	}
}

// Transform does nothing.
func (l Log) Transform(output []byte) []byte { return output }
