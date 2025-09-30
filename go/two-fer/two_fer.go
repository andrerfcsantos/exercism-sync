package twofer

import "fmt"

// ShareWith tells who you should share with
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
