// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	rbt "github.com/lemonyxk/gods/trees/redblacktree"
	"github.com/lemonyxk/gods/utils"
)

// RedBlackTreeExtended to demonstrate how to extend a RedBlackTree to include new functions
type RedBlackTreeExtended[T comparable, P any] struct {
	*rbt.Tree[T, P]
}

// GetMin gets the min value and flag if found
func (tree *RedBlackTreeExtended[T, P]) GetMin() (value P, found bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	}
	return utils.AnyEmpty[P](), false
}

// GetMax gets the max value and flag if found
func (tree *RedBlackTreeExtended[T, P]) GetMax() (value interface{}, found bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	}
	return nil, false
}

// RemoveMin removes the min value and flag if found
func (tree *RedBlackTreeExtended[T, P]) RemoveMin() (value interface{}, deleted bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	}
	return nil, false
}

// RemoveMax removes the max value and flag if found
func (tree *RedBlackTreeExtended[T, P]) RemoveMax() (value interface{}, deleted bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	}
	return nil, false
}

func (tree *RedBlackTreeExtended[T, P]) getMinFromNode(node *rbt.Node[T, P]) (foundNode *rbt.Node[T, P], found bool) {
	if node == nil {
		return nil, false
	}
	if node.Left == nil {
		return node, true
	}
	return tree.getMinFromNode(node.Left)
}

func (tree *RedBlackTreeExtended[T, P]) getMaxFromNode(node *rbt.Node[T, P]) (foundNode *rbt.Node[T, P], found bool) {
	if node == nil {
		return nil, false
	}
	if node.Right == nil {
		return node, true
	}
	return tree.getMaxFromNode(node.Right)
}

func print(tree *RedBlackTreeExtended[int, string]) {
	max, _ := tree.GetMax()
	min, _ := tree.GetMin()
	fmt.Printf("Value for max key: %v \n", max)
	fmt.Printf("Value for min key: %v \n", min)
	fmt.Println(tree)
}

// RedBlackTreeExtendedExample main method on how to use the custom red-black tree above
func main() {
	tree := RedBlackTreeExtended[int, string]{rbt.NewWithIntComparator[int, string]()}

	tree.Put(1, "a") // 1->x (in order)
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(3, "c") // 1->x, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->x, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->x, 2->b, 3->c, 4->d, 5->e (in order)

	print(&tree)
	// Value for max key: e
	// Value for min key: a
	// RedBlackTree
	// │       ┌── 5
	// │   ┌── 4
	// │   │   └── 3
	// └── 2
	//     └── 1

	tree.RemoveMin() // 2->b, 3->c, 4->d, 5->e (in order)
	tree.RemoveMax() // 2->b, 3->c, 4->d (in order)
	tree.RemoveMin() // 3->c, 4->d (in order)
	tree.RemoveMax() // 3->c (in order)

	print(&tree)
	// Value for max key: c
	// Value for min key: c
	// RedBlackTree
	// └── 3
}
