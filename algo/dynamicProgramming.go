package algo

import (
	"fmt"
	"math"
	"strings"
)

/*
* Simple Fibonacci with dynamic programming
 */
func Fibonacci(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	fmt.Println(dp)

	return dp[n]
}

/*
* 1/0 Knapsack problem
 */
func KnapSack01(n, maxw int16, w, p []int16) int16 {

	//var dp [4][7]int16
	dp := make([][]int16, 0)

	var i, j, x int16
	for x = 0; x <= n; x++ {
		r := make([]int16, 0)
		dp = append(dp, r)
	}

	for i = 0; i <= n; i++ {
		for j = 0; j <= maxw; j++ {
			if i == 0 || j == 0 {
				//dp[i][j] = 0
				dp[i] = append(dp[i], 0)

			} else if w[i-1] <= j {
				//dp[i][j] = int16(math.Max(float64(p[i-1]+dp[i-1][j-w[i-1]]), float64(dp[i-1][j])))
				max := int16(math.Max(float64(p[i-1]+dp[i-1][j-w[i-1]]), float64(dp[i-1][j])))
				dp[i] = append(dp[i], max)
			} else {
				//dp[i][j] = dp[i-1][j]
				dp[i] = append(dp[i], dp[i-1][j])
			}
		}
	}

	return dp[n][maxw]
}

func lengthOfLongestPalindromeSubstring(s string) int16 {

	chrs := []rune(s)
	dp := make([][]int16, 0)
	n := int16(len(s))

	var x int16
	for x = 0; x <= n; x++ {
		r := make([]int16, n+1)
		dp = append(dp, r)
	}

	for i := 0; i <= int(n); i++ {
		for j := 0; j <= i; j++ {
			if j > 2 {
				break
			}
			dp[i][j] = int16(j)
		}
	}

	fmt.Println(dp)

	k := int16(3)

	for i := int16(3); i <= n; i++ {
		for j := int16(3); j <= k; j++ {

			if isPalindrome(chrs[:j]) {

				dp[i][j] = j

			} else {
				var v int16
				if float64(dp[i][j-1]) == float64(dp[i-1][j-1])+float64(1) {
					v = dp[i][j-1]
				} else {
					v = int16(math.Max(float64(dp[i][j-1]), float64(dp[i-1][j-1])))
				}
				dp[i][j] = v
			}
		}
		k++
	}

	return dp[n][n]
}

func longestPalindrome(s string) string {
	chrs := []rune(s)
	dp := make([][]string, 0)
	n := int16(len(s))

	var x int16
	for x = 0; x <= n; x++ {
		r := make([]string, n+1)
		dp = append(dp, r)
	}

	for i := 0; i <= int(n); i++ {
		for j := 0; j <= i; j++ {
			if j > 2 {
				break
			}

			if j == 0 || i == 0 {
				dp[i][j] = " "
			} else {
				v := string(chrs[i-j : i])
				if isPalindrome([]rune(v)) {
					dp[i][j] = v
				} else {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	k := int16(3)

	for i := int16(3); i <= n; i++ {
		for j := int16(3); j <= k; j++ {

			if isPalindrome(chrs[:j]) {

				dp[i][j] = string(chrs[:j])

			} else {
				v := string(chrs[i-j : i])
				if isPalindrome([]rune(v)) && len(dp[i-1][j-1]) > len(v) {
					dp[i][j] = dp[i-1][j-1]
				} else {
					if isPalindrome([]rune(v)) && len(dp[i-1][j-1]) > len(v) {
						dp[i][j] = v
					} else {
						if len(dp[i-1][j-1]) > len(dp[i][j-1]) {
							dp[i][j] = dp[i-1][j-1]
						} else {
							dp[i][j] = dp[i][j-1]
						}

					}
				}
			}
			fmt.Println(dp[i])
		}

		k++
	}

	return dp[n][n]
}

func isPalindrome(str []rune) bool {
	var revStr []rune
	stack := make([]rune, 0)

	for i := 0; i < len(str); i++ {
		stack = append(stack, str[i])
	}

	for j := len(stack) - 1; j >= 0; j-- {
		revStr = append(revStr, stack[j])
	}

	return strings.TrimSpace(string(str)) == strings.TrimSpace(string(revStr))
}

func PlanidromPartintion(inpt string) int {

	n := len(inpt)

	if n == 0 || n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = i - (i - 1)
	}

	for i := 2; i <= n; i++ {

		if isPalindrome([]rune(inpt[:i])) {
			dp[i] = 0
		} else {
			for j := i - 1; j > 1; j-- {
				fmt.Println(string([]rune(inpt[(i - j):i])))
				if isPalindrome([]rune(inpt[(i - j):i])) {
					cuts := dp[i-j] + 1
					math.Min(float64(dp[i]), float64(cuts))
				}
			}
		}
		fmt.Println(i, inpt, dp[i], dp)
	}

	return dp[n]
}

func CoinsSum(list []int, sum int) int {
	fmt.Println(list)
	n := len(list)
	dp := make([]int, sum+1)

	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := list[i]; j <= sum; j++ {
			dp[j] += dp[j-list[i]]
			fmt.Println(dp[j], dp)
		}
	}

	return dp[sum]
}
