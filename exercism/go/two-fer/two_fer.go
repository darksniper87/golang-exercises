// Package twofer provides the implementation of 'Two Fer' exercise.
package twofer

import (
	"fmt"
)

// ShareWith returns the full Two-fer sentence depending on the given name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
