package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FastIO struct{}

func (io FastIO) GetArrayInputs() [][]int {
	ioreader := bufio.NewReader(os.Stdin)
	t, _ := ioreader.ReadString('\n')
	testCases, _ := strconv.Atoi(strings.TrimSpace(t))
	input := make([][]int, 0)

	for i := 0; i < testCases; i++ {
		inp := io.ScanLine(ioreader)
		difs := strings.Split(inp, ",")
		inputTemp := make([]int, 0)

		//inp := io.scanLine(ioreader)
		//difs := strings.Split(inp, " ")
		for _, d := range difs {
			tmp, _ := strconv.Atoi(strings.TrimSpace(d))
			inputTemp = append(inputTemp, tmp)
		}

		// Add to inputs array
		input = append(input, inputTemp)
	}

	return input
}

func (io FastIO) ScanLine(ioreader *bufio.Reader) string {
	fmt.Println("Enter test case")
	s, _ := ioreader.ReadString('\n')
	return s
}
