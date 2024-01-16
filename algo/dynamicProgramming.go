package algo

import (
	"fmt"
	"math"
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

func BellNumber(n int) int {
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][i-1]
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
		}
	}

	fmt.Println(dp)
	return dp[n][n]
}

// gfg.org/coin-change-dp-7/
// 1D approach
func CoinChangeDp(coins []int, sum int) int {
	n := len(coins)
	dp := make([]int, sum+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := coins[i]; j <= sum; j++ {
			fmt.Println(j, i, coins[i], dp[j], j-coins[i], dp[j-coins[i]], dp[j]+dp[j-coins[i]])
			dp[j] += dp[j-coins[i]]
		}
	}

	fmt.Println(dp)
	return dp[sum]
}

// #google
// gfg.org/cutting-a-rod-dp-13/
func RodCutting(prices []int, n int) int {

	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		max_val := math.MinInt
		for j := 0; j < i; j++ {
			max_val = int(math.Max(float64(max_val), float64(prices[j]+dp[i-j-1])))
		}
		dp[i] = max_val
	}

	return dp[n]
}

func WordBreak(input string, dic []string) bool {
	hm := make(map[string]bool)

	for _, v := range dic {
		hm[v] = true
	}

	dp := make([]bool, len(input)+1)
	dp[0] = true

	for i := 1; i <= len(input); i++ {
		for j := 0; j < i; j++ {
			subs := string(input[j:i])
			if hm[subs] {
				dp[i] = true && dp[i-len(subs)]
			}
		}
	}

	return dp[len(input)]
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
