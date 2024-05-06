package builtins

import (
	"fmt"
	"io"
)

func UnAlias(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing alias name")
	}

	for _, alias := range args {
		_, ok := aliases[alias]

		if ok {
			delete(aliases, alias)
			fmt.Fprintf(w, "Removed alias: %s \n", alias)
		} else {
			fmt.Fprintf(w, "Alias '%s' not found\n", alias)

		}
	}

	return nil
}
