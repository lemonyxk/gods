// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package avltree implements an AVL balanced binary tree.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/AVL_tree
package avltree

import (
	"fmt"

	"github.com/emirpasic/gods/utils"
	"github.com/lemonyxk/gods/trees"
)

func assertTreeImplementation[T comparable, P any]() {
	var _ trees.Tree[T, P] = new(Tree[T, P])
}

// Tree holds elements of the AVL tree.
type Tree[T comparable, P any] struct {
	Root       *Node[T, P]      // Root node
	Comparator utils.Comparator // Key comparator
	size       int              // Total number of keys in the tree
}

// Node is a single element within the tree
type Node[T comparable, P any] struct {
	Key      T
	Value    P
	Parent   *Node[T, P]    // Parent node
	Children [2]*Node[T, P] // Children nodes
	b        int8
}

// NewWith instantiates an AVL tree with the custom comparator.
func NewWith[T comparable, P any](comparator utils.Comparator) *Tree[T, P] {
	return &Tree[T, P]{Comparator: comparator}
}

// NewWithIntComparator instantiates an AVL tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator[T comparable, P any]() *Tree[T, P] {
	return &Tree[T, P]{Comparator: utils.IntComparator}
}

// NewWithStringComparator instantiates an AVL tree with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator[T comparable, P any]() *Tree[T, P] {
	return &Tree[T, P]{Comparator: utils.StringComparator}
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[T, P]) Put(key T, value P) {
	t.put(key, value, nil, &t.Root)
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[T, P]) Get(key T) (value P, found bool) {
	n := t.Root
	for n != nil {
		cmp := t.Comparator(key, n.Key)
		switch {
		case cmp == 0:
			return n.Value, true
		case cmp < 0:
			n = n.Children[0]
		case cmp > 0:
			n = n.Children[1]
		}
	}
	var p P
	return p, false
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[T, P]) Remove(key T) {
	t.remove(key, &t.Root)
}

// Empty returns true if tree does not contain any nodes.
func (t *Tree[T, P]) Empty() bool {
	return t.size == 0
}

// Size returns the number of elements stored in the tree.
func (t *Tree[T, P]) Size() int {
	return t.size
}

// Keys returns all keys in-order
func (t *Tree[T, P]) Keys() []T {
	keys := make([]T, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (t *Tree[T, P]) Values() []P {
	values := make([]P, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

// Left returns the minimum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree[T, P]) Left() *Node[T, P] {
	return t.bottom(0)
}

// Right returns the maximum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree[T, P]) Right() *Node[T, P] {
	return t.bottom(1)
}

// Floor Finds floor node of the input key, return the floor node or nil if no ceiling is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[T, P]) Floor(key T) (floor *Node[T, P], found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			n = n.Children[0]
		case c > 0:
			floor, found = n, true
			n = n.Children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree is smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[T, P]) Ceiling(key T) (floor *Node[T, P], found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			floor, found = n, true
			n = n.Children[0]
		case c > 0:
			n = n.Children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Clear removes all nodes from the tree.
func (t *Tree[T, P]) Clear() {
	t.Root = nil
	t.size = 0
}

// String returns a string representation of container
func (t *Tree[T, P]) String() string {
	str := "AVLTree\n"
	if !t.Empty() {
		output(t.Root, "", true, &str)
	}
	return str
}

func (n *Node[T, P]) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func (t *Tree[T, P]) put(key T, value P, p *Node[T, P], qp **Node[T, P]) bool {
	q := *qp
	if q == nil {
		t.size++
		*qp = &Node[T, P]{Key: key, Value: value, Parent: p}
		return true
	}

	c := t.Comparator(key, q.Key)
	if c == 0 {
		q.Key = key
		q.Value = value
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	var fix bool
	fix = t.put(key, value, q, &q.Children[a])
	if fix {
		return putFix(int8(c), qp)
	}
	return false
}

func (t *Tree[T, P]) remove(key T, qp **Node[T, P]) bool {
	q := *qp
	if q == nil {
		return false
	}

	c := t.Comparator(key, q.Key)
	if c == 0 {
		t.size--
		if q.Children[1] == nil {
			if q.Children[0] != nil {
				q.Children[0].Parent = q.Parent
			}
			*qp = q.Children[0]
			return true
		}
		fix := removeMin(&q.Children[1], &q.Key, &q.Value)
		if fix {
			return removeFix(-1, qp)
		}
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	fix := t.remove(key, &q.Children[a])
	if fix {
		return removeFix(int8(-c), qp)
	}
	return false
}

func removeMin[T comparable, P any](qp **Node[T, P], minKey *T, minVal *P) bool {
	q := *qp
	if q.Children[0] == nil {
		*minKey = q.Key
		*minVal = q.Value
		if q.Children[1] != nil {
			q.Children[1].Parent = q.Parent
		}
		*qp = q.Children[1]
		return true
	}
	fix := removeMin(&q.Children[0], minKey, minVal)
	if fix {
		return removeFix(1, qp)
	}
	return false
}

func putFix[T comparable, P any](c int8, t **Node[T, P]) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.Children[(c+1)/2].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return false
}

func removeFix[T comparable, P any](c int8, t **Node[T, P]) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return false
	}

	if s.b == -c {
		s.b = 0
		return true
	}

	a := (c + 1) / 2
	if s.Children[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s
		return false
	}

	if s.Children[a].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return true
}

func singlerot[T comparable, P any](c int8, s *Node[T, P]) *Node[T, P] {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

func doublerot[T comparable, P any](c int8, s *Node[T, P]) *Node[T, P] {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = rotate(-c, s.Children[a])
	p := rotate(c, s)

	switch {
	default:
		s.b = 0
		r.b = 0
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	}

	p.b = 0
	return p
}

func rotate[T comparable, P any](c int8, s *Node[T, P]) *Node[T, P] {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}
	r.Children[a^1] = s
	r.Parent = s.Parent
	s.Parent = r
	return r
}

func (t *Tree[T, P]) bottom(d int) *Node[T, P] {
	n := t.Root
	if n == nil {
		return nil
	}

	for c := n.Children[d]; c != nil; c = n.Children[d] {
		n = c
	}
	return n
}

// Prev returns the previous element in an inorder
// walk of the AVL tree.
func (n *Node[T, P]) Prev() *Node[T, P] {
	return n.walk1(0)
}

// Next returns the next element in an inorder
// walk of the AVL tree.
func (n *Node[T, P]) Next() *Node[T, P] {
	return n.walk1(1)
}

func (n *Node[T, P]) walk1(a int) *Node[T, P] {
	if n == nil {
		return nil
	}

	if n.Children[a] != nil {
		n = n.Children[a]
		for n.Children[a^1] != nil {
			n = n.Children[a^1]
		}
		return n
	}

	p := n.Parent
	for p != nil && p.Children[a] == n {
		n = p
		p = p.Parent
	}
	return p
}

func output[T comparable, P any](node *Node[T, P], prefix string, isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Children[0], newPrefix, true, str)
	}
}
