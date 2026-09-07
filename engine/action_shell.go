package engine

import (
	"fmt"
	"os/exec"
)

// ShellAction runs Command through the shell. No timeout or stdin
// handling yet — deliberately bare-bones for now.
type ShellAction struct {
	Command string
}

func (a *ShellAction) Execute() error {
	out, err := exec.Command("sh", "-c", a.Command).CombinedOutput()
	fmt.Printf("shell action: %q -> %s", a.Command, out)
	if err != nil {
		return fmt.Errorf("shell action failed: %w", err)
	}
	return nil
}
