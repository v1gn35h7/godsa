package algo

import (
	"fmt"
)

/*
* Takes a string and returns all subsets of the strings
 */
func PrintAllSubsets(input string) map[string]bool {

	output := make(map[string]bool)
	subSetUtil(input, "", output)
	fmt.Println(output)

	return output
}

func subSetUtil(input, s string, output map[string]bool) {
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

// gfg.org/coin-change-dp-7/
func CoinChange(list []int, n, sum int) int {
	if sum == 0 {
		return 1
	}

	if sum < 0 {
		return 0
	}

	if n <= 0 {
		return 0
	}

	return CoinChange(list, n, sum-list[n-1]) + CoinChange(list, n-1, sum)

}
