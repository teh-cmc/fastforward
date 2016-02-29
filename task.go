package forward

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

// -----------------------------------------------------------------------------

// TemplateTaskNew is the template for the `fwd task new` command.
const TemplateTaskNew = `
# Please enter the name and description (optional) of your new task, separated
# by an empty line. Names longer than 80 characters will be truncated.
#
# Lines starting with '#' will be ignored, and an empty message aborts the
# creation of the Forward project.
`

// -----------------------------------------------------------------------------

// TaskNew implements the `fwd task new` command.
//
// It initializes a new git repository at the specified `path` and creates the
// Forward directory hierarchy.
func TaskNew(c *cli.Context) {
	// get title and description
	t, d, err := TitleAndDescription(Edit("task-new", TemplateTaskNew))
	if err != nil {
		log.Fatal(err)
	}
	if t == "" { // NOTE: empty message -> abort
		return
	}

	// add task:new commit
	tid, err := LatestTaskID()
	if err != nil {
		log.Fatal(err)
	}
	t = fmt.Sprintf("[forward] task:new - (#%d) %v", tid+1, t)
	cmd := []string{"commit", "--allow-empty", "--file", "-"}
	if err := GitExec(cmd, []byte(t+"\n\n"+d)); err != nil {
		log.Fatal(err)
	}
}
