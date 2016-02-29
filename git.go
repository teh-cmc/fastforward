package forward

import (
	"os"
	"os/exec"
)

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
	c := exec.Command(bin, cmd...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	in, err := c.StdinPipe()
	defer in.Close()
	if err != nil {
		return err
	}
	if err := c.Start(); err != nil {
		return err
	}
	_, err = in.Write(input)
	if err != nil {
		return err
	}
	in.Close()
	if err := c.Wait(); err != nil {
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
