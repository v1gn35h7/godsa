package algo

import "strings"

func substr(needle, hay string) bool {
	return strings.Contains(hay, needle)
}
