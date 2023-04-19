package algo

import (
	"fmt"
)

/*
* Takes a string and returns all subsets of
 */
func PrintAllSubsets(input string) map[string]bool {

	output := make(map[string]bool)
	subSetUtil(input, "", output)
	fmt.Println(output)

	return output

}

func subSetUtil(input, s string, output map[string]bool) {
	fmt.Println(input, s)
	ln := len(input)
	if ln == 0 {
		if len(s) > 0 {
			output[s] = true
		}
		return
	}

	output[input+s] = true

	subSetUtil(input[:(ln-1)], input[(ln-1):], output)
	subSetUtil(input[:(ln-1)], s, output)
}
