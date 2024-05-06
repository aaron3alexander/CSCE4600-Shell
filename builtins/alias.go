package builtins

import (
	"fmt"
)

var aliases = make(map[string][]string)

func Alias(args ...string) error {
	if len(args) == 0 {
		for alias, command := range aliases {
			fmt.Printf("%s='%s'\n", alias, command)
		}
		if len(aliases) == 0 {
			return fmt.Errorf("no aliases")
		}
		return nil
	}
	if len(args) == 1 {
		command, ok := aliases[args[0]]
		if ok {
			fmt.Println(command)
			return nil
		} else {
			return fmt.Errorf("alias not found")
		}
	}
	aliases[args[0]] = args[1:]

	return nil
}

func AliasMap() map[string][]string {
	return aliases
}
