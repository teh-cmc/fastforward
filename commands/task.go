package commands

import (
	"github.com/teh-cmc/fastforward"
	"github.com/teh-cmc/fastforward/git"
)

// -----------------------------------------------------------------------------

// TaskNew implements the `ff task new` command.
type TaskNew struct{}

// NewTaskNew returns a new TaskNew.
//
// TaskNew implements the `forward.Commitable` interface.
func NewTaskNew() *TaskNew { return &TaskNew{} }

// -----------------------------------------------------------------------------

// AllowAutoPulling always returns `true`.
func (t TaskNew) AllowAutoPulling() bool { return true }

// AllowAutoPushing always returns `true`.
func (t TaskNew) AllowAutoPushing() bool { return true }

// TaskNew implements the `fwd init` command.
func (t TaskNew) Run(branch string) ([]byte, error) {
	cm, err := forward.EditMessage(t)
	if err != nil {
		return nil, err
	}
	return git.Run(git.NewCommit(branch, cm), branch)
}

// -----------------------------------------------------------------------------

// Command returns the command name for commit messages.
func (t TaskNew) Command() string { return "task:new" }

// Template returns the template for commit messages.
func (t TaskNew) Template() []byte {
	return []byte(`
# Please enter the name and description (optional) of your new task, separated
# by an empty line. Names longer than 80 characters will be truncated.
# Optional attributes of the form 'attr:a,b,c,d...' can also be specified.
#
# Lines starting with '#' will be ignored, and an empty message aborts the
# creation of the task.
#
# -- EXAMPLE --
#
# segfault on malformed inputs
#
# It would seem that feeding the input with an extra linefeed results in a
# segfault.
#
# tags:bug,input,segfault
# files:input.go,user.go
`)
}
