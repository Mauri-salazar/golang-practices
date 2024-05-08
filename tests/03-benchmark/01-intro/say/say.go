// The say package allows us to say hello
package say

import "fmt"

// with the function Greet, we can to greet to a person
func Greet(s string) string {
	return fmt.Sprint("Welcome Dear ", s)
}
