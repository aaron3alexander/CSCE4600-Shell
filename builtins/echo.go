package builtins

import (
	"fmt"
	"strings"
)

func Echo(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no arguments given")
	}
	output := strings.Join(args, " ")
	fmt.Println(output)
	return nil
}
