// The package word.We to allows work with strings
package word

import "strings"

// UseCount return a map with words of string and the number of times it is repeated the word.
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return the total amount of the words
func Count(s string) int {
	// write the code for this func
	return len(s)
}
