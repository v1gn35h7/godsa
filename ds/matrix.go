package ds

import (
	"fmt"
	"math/rand"
	"time"
)

func FindNumberOfIslands(mat [5][5]int, v int) {
	var visited [5][5]bool
	count := 0
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			if mat[i][j] == 1 && visited[i][j] == false {
				//fmt.Printf("%v %d %d --> ", visited[i][j], i, j)

				mdfs(mat, &visited, i, j)
				count++
				//fmt.Printf("\n")
				//fmt.Println(visited)
			}
		}
	}

	fmt.Println("Total islands: ", count)
}

func mdfs(mat [5][5]int, visited *[5][5]bool, row, col int) {
	visited[row][col] = true
	//fmt.Printf("%v", visited[row][col])
	//fmt.Printf(" %d %d %v --> ", row, col, visited[row][col])

	rowNb := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	colNb := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		if isSafeToDfs(mat, visited, row+rowNb[i], col+colNb[i]) {
			mdfs(mat, visited, row+rowNb[i], col+colNb[i])
		}
	}
}

func isSafeToDfs(mat [5][5]int, visited *[5][5]bool, row, col int) bool {

	if row >= 0 && col >= 0 && row < 5 && col < 5 {
		//fmt.Println(row, col)
		return visited[row][col] == false && mat[row][col] == 1
	} else {
		return false
	}
}

// New 2D array
func New2dArray(i, j int) [][]int {
	arr := make([][]int, i)

	for i, _ := range arr {
		arr[i] = make([]int, j)
	}

	for x, v := range arr {
		for y, _ := range v {
			// Seed the random number generator with current time
			rand.Seed(time.Now().UnixNano())

			// Generate a random integer between 0 and 100000
			randomInt := rand.Intn(100000)

			arr[x][y] = randomInt
		}
	}

	return arr
}

// print 2d array
func Print2dArray(arr [][]int) {
	i := len(arr)

	if i == 0 {
		fmt.Println("Zero sized array")
		return
	}

	for _, x := range arr {
		for _, y := range x {
			fmt.Printf("%d ", y)
		}
		fmt.Println("\n")
	}

}
