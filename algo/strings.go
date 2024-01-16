package algo

import (
	"fmt"
	"strings"
)

func RevereStringWithSPace(inp string) string {
	res := make([]string, 0)
	runes := []rune(inp)
	n := len(inp)
	sstart := 0
	isSpace := false
	var resSTring strings.Builder

	for i := 0; i < n; i++ {

		if runes[i] == 32 && !isSpace {
			res = append(res, string(runes[sstart:i]))
			fmt.Println(runes[i], res)

			sstart = i
			isSpace = true
		} else if runes[i] != 32 && isSpace {
			res = append(res, string(runes[sstart:i]))
			sstart = i
			isSpace = false
		} else if i == n-1 {
			res = append(res, string(runes[sstart:i]))
		}
	}

	for j := len(res); j > 0; j-- {
		resSTring.Write([]byte(res[j-1]))
	}

	return string(resSTring.String())
}

func RevrWordsWithoutStack(inp string) string {
	runes := []byte(inp)
	n := len(inp)
	var res strings.Builder
	isSpace := false

	for i := n; i > 0; i-- {
		if runes[i-1] == 32 && !isSpace {
			res.Write(runes[i:])
			runes = runes[:i]
			isSpace = true
		} else if runes[i-1] != 32 && isSpace {
			res.Write(runes[i:])
			runes = runes[:i]
			isSpace = false
		}

		if i-1 == 0 {
			res.Write(runes)
		}
	}

	return res.String()
}
