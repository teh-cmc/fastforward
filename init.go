package forward

import (
	"log"
	"path/filepath"

	"github.com/codegangsta/cli"
)

// -----------------------------------------------------------------------------

// TemplateInit is the template for the `fwd init` command.
const TemplateInit = `
# Please enter the name and description (optional) of your new Forward project,
# separated by an empty line. Names longer than 80 characters will be truncated.
#
# Lines starting with '#' will be ignored, and an empty message aborts the
# creation of the Forward project.
`

// -----------------------------------------------------------------------------

// Init implements the `fwd init` command.
//
// It initializes a new git repository at the specified `path` and creates the
// Forward directory hierarchy.
func Init(c *cli.Context) {
	// get directory
	dir := "." // NOTE: `dir` defaults to current directory
	if p := c.Args().Get(0); p != "" {
		dir = p
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Clean(dir)

	// get title and description
	t, d, err := TitleAndDescription(Edit("init", TemplateInit))
	if err != nil {
		log.Fatal(err)
	}
	if t == "" { // NOTE: empty message -> abort
		return
	}

	// init git repository
	if err := GitExec([]string{"init", dir}, nil); err != nil {
		log.Fatal(err)
	}

	// add init commit
	cmd := []string{"commit", "--allow-empty", "--file", "-"}
	if err := GitExec(cmd, []byte(t+"\n"+d)); err != nil {
		log.Fatal(err)
	}
}
