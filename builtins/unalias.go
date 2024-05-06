package builtins

import (
	"fmt"
)

func UnAlias(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing alias name")
	}

	for _, alias := range args {
		_, ok := aliases[alias]

		if ok {
			delete(aliases, alias)
			fmt.Printf("Removed alias: %s \n", alias)
		} else {
			fmt.Printf("Alias '%s' not found\n", alias)

		}
	}

	return nil
}
