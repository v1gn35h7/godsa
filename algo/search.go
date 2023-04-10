package algo

import (
	"strconv"
	"strings"
)


func binary_search(list []string, l, r, x int) int {
	if r < 0 {
		return -1
	}
	m := l + ((r - l) / 2)
	y, _ := strconv.Atoi(strings.TrimSpace(list[m]))
	if x == y {
		return m
	} else if x < y {
		return binary_search(list, l, m-1, x)
	} else if x > y {
		return binary_search(list, m+1, r, x)
	}
	return -1
}
