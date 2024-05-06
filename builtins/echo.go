package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no arguments given")
	}
	output := strings.Join(args, " ")
	fmt.Fprintln(w, output)
	return nil
}
