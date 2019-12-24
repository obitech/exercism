// Package proverb contains a function to construct a proverb.
package proverb

import "fmt"

// Proverb create a proverb given a slice of input strings.
func Proverb(rhyme []string) []string {
	out := []string{}
	if len(rhyme) == 0 {
		return out
	}
	for i := 0; i < len(rhyme)-1; i++ {
		out = append(out, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	out = append(out, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return out
}
