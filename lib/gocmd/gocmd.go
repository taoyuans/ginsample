package gocmd

import (
	"fmt"
	"os/exec"
)

func GetCommandResult(command string, args []string) (string, error) {
	if command == "" || len(command) <= 0 {
		return "", nil
	}
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	} else {
		return string(out), nil
	}
}

func FmtCommandResult(command string, args []string) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out => :\n%s\n", string(out))
	}
	fmt.Printf("combined out => :\n%s\n", string(out))
}
