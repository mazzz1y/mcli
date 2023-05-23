package subprocess

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mazzz1y/mcli/internal/shortcuts"
	"github.com/riywo/loginshell"
)

func Exec(command string, envs []shortcuts.Env) error {
	shell, err := loginshell.Shell()
	if err != nil {
		return err
	}

	cmd := exec.Command(shell, "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	for _, e := range envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", e.Key, e.Value))
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}
