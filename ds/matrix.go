package ds

import "fmt"

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
