package algo

import (
	"math"
)

/*
* Simple Fibonacci with dynamic programming
 */
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	fib := make(map[int]int)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[(i-2)] + fib[(i-1)]
	}

	return fib[n]
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
