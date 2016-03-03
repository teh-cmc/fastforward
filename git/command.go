package git

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

// -----------------------------------------------------------------------------

// Command exposes methods to run a `git` command.
type Command interface {
	AllowAutoCheckout() bool
	Input() []byte
	Command() []string
	Transform([]byte) []byte
}

// -----------------------------------------------------------------------------

// Run runs the given command `c`.
func Run(c Command, branch string) ([]byte, error) {
	bin, err := exec.LookPath("git")
	if err != nil {
		return nil, err
	}

	if c.AllowAutoCheckout() {
		current, err := Run(NewBranch(BranchTypeCurrent, branch), branch)
		if err != nil {
			return nil, err
		}
		if _, err = Run(NewBranch(BranchTypeSwitch, branch), branch); err != nil {
			return nil, err
		}
		defer func() {
			if _, err := Run(NewBranch(BranchTypeSwitch, string(current)), branch); err != nil {
				log.Fatal(err)
			}
		}()
	}

	cmd := exec.Command(bin, c.Command()...)
	output := bytes.NewBufferString("")
	cmd.Stderr = os.Stderr
	cmd.Stdout = output

	if input := c.Input(); input != nil {
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

	return c.Transform(output.Bytes()), nil
}
