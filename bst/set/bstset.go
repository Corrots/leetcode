package set

import "github.com/corrots/leetcode/bst/bst"

type BSTSet struct {
	bst *bst.BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{bst: bst.NewBST()}
}

func (b *BSTSet) Add(i interface{}) {
	panic("implement me")
}

func (b *BSTSet) Remove(i interface{}) {
	panic("implement me")
}

func (b *BSTSet) Contains(i interface{}) {
	panic("implement me")
}

func (b *BSTSet) Len() int {
	return b.bst.Len()
}

func (b *BSTSet) IsEmpty() bool {
	return b.bst.IsEmpty()
}
