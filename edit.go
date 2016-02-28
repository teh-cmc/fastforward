package forward

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// -----------------------------------------------------------------------------

const (
	templateIssue = `
# Please enter the name and description of your issue, separated by an empty
# line. Names longer than 80 characters will be truncated.
#
# Lines starting with '#' will be ignored, and an empty message aborts the
# creation of an issue.
`
)

// -----------------------------------------------------------------------------

// Edit starts the user's favorite editor and returns the output.
func Edit(prefix string) ([]byte, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return nil, fmt.Errorf("$EDITOR isn't set")
	}
	path, err := exec.LookPath(editor)
	if err != nil {
		return nil, err
	}

	f, err := ioutil.TempFile(os.TempDir(), fmt.Sprintf(".forward.%s-"))
	if err != nil {
		return nil, err
	}
	defer func() {
		// NOTE: file removal might fail, not much we can do though.
		_ = os.Remove(f.Name())
	}()

	cmd := exec.Command(path, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
