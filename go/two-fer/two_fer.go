// Package twofer holds a function indicating sharing something with someone else.
package twofer

// ShareWith emits a string indicating sharing an item with the name passed,
// or sharing it with just `you` if the passed name is an empty string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
