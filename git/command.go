package git

import (
	"bytes"
	"os"
	"os/exec"
)

// -----------------------------------------------------------------------------

// Command exposes methods to run a git command.
type Command interface {
	Template() []byte
	Command() []string
}

// Run runs the given command `c`.
func Run(c Command) ([]byte, error) {
	bin, err := exec.LookPath("git")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(bin, c.Command()...)
	cmd.Stderr = os.Stderr

	output := bytes.NewBufferString("")
	cmd.Stdout = output

	if input := c.Template(); input != nil {
		in, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		defer func() { _ = in.Close() }()
		if err = cmd.Start(); err != nil {
			return nil, err
		}
		_, err = in.Write(input)
		if err != nil {
			return nil, err
		}
		_ = in.Close()
		if err = cmd.Wait(); err != nil {
			return nil, err
		}
	} else {
		if err := cmd.Run(); err != nil {
			return nil, err
		}
	}

	return output.Bytes(), nil
}
