package ds

import "fmt"

type sTree struct {
	n    int
	data []int
}

func NewSegTree(list []int) *sTree {

	tree := &sTree{
		len(list),
		make([]int, len(list)*len(list)),
	}

	tree.build(list)

	return tree
}

func (tree *sTree) build(list []int) {
	n := len(list)

	for i := 0; i < n; i++ {
		tree.data[n+i] = list[i]
	}

	for i := n - 1; i > 0; i-- {
		tree.data[i] = tree.data[i<<1] + tree.data[i<<1|1]
	}

	fmt.Println(tree.data)

}

func (tree *sTree) Search(l, r int) int {
	l = tree.n + l
	r = tree.n + r
	res := 0

	for l < r {
		if (l & 1) > 0 {
			res += tree.data[l]
			l += 1
		}

		if (r & 1) > 0 {
			r -= 1
			res += tree.data[r]

		}

		l = l >> 1
		r = r >> 1
	}

	fmt.Println(res)
	return res
}

func (tree *sTree) Update(p, v int) {
	tree.data[p+tree.n] = v
	p = p + tree.n

	for i := p; i > 1; i >>= 1 {
		tree.data[i>>1] = tree.data[i] + tree.data[i^1]
	}

}
