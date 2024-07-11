package go_cli_completion

import (
	"testing"
)

func CheckErrorType(t *testing.T, completions []string, errtype uint8) {
	switch errtype {
	case CMD_EXEC_FAIL:
		t.Fatal("Command execution failed.")
	case CMD_PIPE_FAIL:
		t.Fatal("Command setup failed.")
	default:
		words := completions[0]
		for idx, completion := range completions {
			if idx != 0 {
				words += ", " + completion
			}
		}
		t.Logf("SUCCESS! Completions is:\n%s\n", words)
	}
}

func TestGitCompletion(t *testing.T) {
	completions, errtype := GenerateCompletions(CompletionType(GIT_BRANCH_A), "")
	CheckErrorType(t, completions, errtype)
}

func TestGitCompletionWithInput(t *testing.T) {
	completions, errtype := GenerateCompletions(CompletionType(GIT_BRANCH_A), "ma")
	CheckErrorType(t, completions, errtype)
}

func TestLS(t *testing.T) {
	completions, errtype := GenerateCompletions(CompletionType(LS), "")
	CheckErrorType(t, completions, errtype)
}

func TestLS_A(t *testing.T) {
	completions, errtype := GenerateCompletions(CompletionType(LS_A), "")
	CheckErrorType(t, completions, errtype)
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestLocalFileCompletion(t *testing.T) {
//     msg, err := Hello("")
//     if msg != "" || err == nil {
//         t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
//     }
// }
