package forward

import (
	"os"
	"os/exec"
)

// -----------------------------------------------------------------------------

// GitOutput runs the given git `cmd` and returns its output.
func GitOutput(cmd []string) ([]byte, error) {
	bin, err := exec.LookPath("git")
	if err != nil {
		return nil, err
	}
	return gitOutput(bin, cmd)
}

func gitOutput(bin string, cmd []string) ([]byte, error) {
	c := exec.Command(bin, cmd...)
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	return c.Output()
}

// -----------------------------------------------------------------------------

// GitExec runs the given git `cmd`.
//
// If `input` is non-nil, it is fed to git via a stdin pipe.
func GitExec(cmd []string, input []byte) error {
	bin, err := exec.LookPath("git")
	if err != nil {
		return err
	}
	if input != nil {
		return gitExecPiped(bin, cmd, input)
	}
	return gitExec(bin, cmd)
}

func gitExecPiped(bin string, cmd []string, input []byte) error {
	var err error
	c := exec.Command(bin, cmd...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	in, err := c.StdinPipe()
	defer func() { _ = in.Close() }()
	if err != nil {
		return err
	}
	if err = c.Start(); err != nil {
		return err
	}
	_, err = in.Write(input)
	if err != nil {
		return err
	}
	_ = in.Close()
	if err = c.Wait(); err != nil {
		return err
	}
	return nil
}

func gitExec(bin string, cmd []string) error {
	c := exec.Command(bin, cmd...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
