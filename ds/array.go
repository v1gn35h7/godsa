package ds

import "fmt"

func PrintArray(input []int) {
	for i := 0; i < len(input); i++ {
		fmt.Println(input[:i], input[i:])
	}
}
