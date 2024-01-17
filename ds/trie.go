package ds

type trie struct {
	data      string
	childrens []*trie
}

func (t *trie) Add(input string) {
	// addRcursive(input, t)
	chars := []byte(input)
	a := 97
	head := t

	for _, v := range chars {
		i := int(v) - a
		if head.childrens[i] == nil {
			head.childrens[i] = &trie{string(v), make([]*trie, 26)}
			head = head.childrens[i]
		}
	}
}

func (t *trie) Search(input string) bool {
	chars := []byte(input)
	a := 97
	head := t

	for _, v := range chars {
		i := int(v) - a
		if head.childrens[i] == nil {
			return false
		} else {
			head = head.childrens[int(v)-a]
		}
	}

	return true
}

func NewTrie() *trie {
	return &trie{data: "$", childrens: make([]*trie, 26)}
}
