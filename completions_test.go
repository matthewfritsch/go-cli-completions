package go_cli_completion

import (
	"testing"
)

func TestGitCompletion(t *testing.T) {
	//     name := "Gladys"
	//     want := regexp.MustCompile(`\b`+name+`\b`)
	//     msg, err := Hello("Gladys")
	//     if !want.MatchString(msg) || err != nil {
	//         t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	//     }

	completions, errtype := GenerateCompletions(CompletionType(GIT_BRANCH_A))
	switch errtype {
	case CMD_EXEC_FAIL:
		t.Fatal("Command execution failed.")
	case CMD_PIPE_FAIL:
		t.Fatal("Command setup failed.")
	default:
		t.Logf("SUCCESS! Completions is:\n%s\n", completions)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestLocalFileCompletion(t *testing.T) {
//     msg, err := Hello("")
//     if msg != "" || err == nil {
//         t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
//     }
// }
