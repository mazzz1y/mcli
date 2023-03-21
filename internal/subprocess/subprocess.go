package subprocess

import (
	"os"
	"os/exec"

	"github.com/riywo/loginshell"
)

func Exec(command string) error {
	shell, err := loginshell.Shell()
	if err != nil {
		return err
	}

	cmd := exec.Command(shell, "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}
