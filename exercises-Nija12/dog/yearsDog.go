// convert the human years to dog years
package yearsDog

// multiply human year * dog year
func YearsDog(a int) int {
	yearsDog := 7

	//condition
	if a == 1 {
		return yearsDog
	} else {
		return a * yearsDog
	}
}
