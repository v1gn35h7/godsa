package algo

/*
* Simple Fibonacci with dynamic programming
* Ref: https://www.geeksforgeeks.org/overlapping-subproblems-property-in-dynamic-programming-dp-1/
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
