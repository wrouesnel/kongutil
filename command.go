package kongutil

import (
	"reflect"
	"slices"
	"strings"

	"github.com/alecthomas/kong"
)

// GetMethod is a copy of the internal method from kong so we can check for run methods
func GetMethod(value reflect.Value, name string) reflect.Value {
	method := value.MethodByName(name)
	if !method.IsValid() {
		if value.CanAddr() {
			method = value.Addr().MethodByName(name)
		}
	}
	return method
}

// Command returns the nearest node to this commands command path
func Command(c *kong.Context) string {
	command := []string{}
mainLoop:
	for i := len(c.Path) - 1; i > -1; i-- {
		trace := c.Path[i]

		switch {
		case trace.Positional != nil:
			command = append(command, "<"+trace.Positional.Name+">")

		case trace.Argument != nil:
			command = append(command, "<"+trace.Argument.Name+">")

		case trace.Command != nil:
			// Inspect the command to see if it has a run idx on its struct.
			if method := GetMethod(trace.Command.Target, "Run"); method.IsValid() {
				break mainLoop
			}
			command = append(command, trace.Command.Name)
		}

		if trace.Parent != nil {
			if trace.Parent.Type != kong.CommandNode {
				break
			}
		}

	}

	slices.Reverse(command)

	return strings.Join(command, " ")
}
