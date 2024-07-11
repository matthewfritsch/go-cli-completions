package go_cli_completion

import (
	"io"
	"os/exec"
	"strings"
)

type CompletionType string

const (
	GIT_BRANCH_A string = "git branch -a"
	LS           string = "ls"
	LS_A         string = "ls -a"
)

const (
	SUCCESS uint8 = iota
	CMD_PIPE_FAIL
	CMD_EXEC_FAIL
)

func GenerateCompletions(ct CompletionType) ([]string, uint8) {
	command_str_split := strings.Split(string(ct), " ")
	command := command_str_split[0]
	args := command_str_split[1:]
	cmd := exec.Command(command, args...)
	out, err := cmd.StdoutPipe()

	if err != nil {
		return []string{}, CMD_PIPE_FAIL
	}

	if cmd.Run() != nil {
		return []string{}, CMD_EXEC_FAIL
	}

	stdout, err := io.ReadAll(out)
	return strings.Split(string(stdout), " "), SUCCESS
}

/*

- a structure for defining all the data around a bash completion
- functions for getting the completion
- functions for storing the completion in cache on a time basis

*/
