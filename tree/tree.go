package tree

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
// TODO: Create this function
func (t *Tree) New(n int) *Tree {
	return &Tree{nil, 10, nil}
}
