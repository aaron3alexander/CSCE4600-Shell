package builtins

import (
	"fmt"
	"io"
)

var commandHistory []string

func AddHistory(command string) {
	commandHistory = append(commandHistory, command)
}

func History(w io.Writer, args ...string) error {

	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	} else if len(args) == 1 {
		if args[0] == "-c" {
			ClearHistory()
			fmt.Fprintf(w, "history cleared\n")
			return nil
		}
	}

	if len(commandHistory) <= 1 {
		fmt.Fprintf(w, "history is empty\n")
		return nil
	}

	for i, command := range commandHistory[:len(commandHistory)-1] {
		fmt.Fprintf(w, "%d: %s\n", i+1, command)
	}
	return nil
}

// ClearHistory clears the command history.
func ClearHistory() {
	commandHistory = nil
}
