package go_cli_completion

import (
	"bytes"
	"os/exec"
	"strings"
)

type CompletionType string

const (
	GIT_BRANCH_A string = "git branch -a | sed 's/^\\*//'"
	LS           string = "ls"
	LS_A         string = "ls -a"
)

const (
	SUCCESS uint8 = iota
	CMD_PIPE_FAIL
	CMD_EXEC_FAIL
)

func GenerateCompletions(ct CompletionType, input string) ([]string, uint8) {
	cmd := exec.Command("bash", "-c", string(ct))

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if cmd.Run() != nil {
		return []string{}, CMD_EXEC_FAIL
	}

	// fmt.Printf("STDOUT: %s\n", stdout.String())
	// fmt.Printf("STDERR: %s\n", stderr.String())

	cleaned := ParsedCompletions(stdout.String())
	partials := PartialCompletions(cleaned, input)
	return partials, SUCCESS
}

func ParsedCompletions(stdout string) []string {
	lines := strings.Split(stdout, "\n")
	var trimmed []string
	for _, line := range lines {
		t := strings.TrimSpace(line)
		if t == "" {
			continue
		}
		trimmed = append(trimmed, t)
	}
	return trimmed
}

func PartialCompletions(cleaned []string, input string) []string {
	var matches []string
	for _, str := range cleaned {
		if len(str) < len(input) {
			continue
		}
		if str[0:len(input)] == input {
			matches = append(matches, str)
		}
	}
	return matches
}

/*

- a structure for defining all the data around a bash completion
- functions for getting the completion
- functions for storing the completion in cache on a time basis

*/
