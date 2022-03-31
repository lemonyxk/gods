// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treebidimap implements a bidirectional map backed by two red-black tree.
//
// This structure guarantees that the map will be in both ascending key and value order.
//
// Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package treebidimap

import (
	"fmt"
	"strings"

	"github.com/lemonyxk/gods/maps"
	"github.com/lemonyxk/gods/trees/redblacktree"
	"github.com/lemonyxk/gods/utils"
)

func assertMapImplementation[T comparable, P comparable]() {
	var _ maps.BidiMap[T, P] = (*Map[T, P])(nil)
}

// Map holds the elements in two red-black trees.
type Map[T comparable, P comparable] struct {
	forwardMap      redblacktree.Tree[T, P]
	inverseMap      redblacktree.Tree[P, T]
	keyComparator   utils.Comparator
	valueComparator utils.Comparator
}

// NewWith instantiates a bidirectional map.
func NewWith[T comparable, P comparable](keyComparator utils.Comparator, valueComparator utils.Comparator) *Map[T, P] {
	return &Map[T, P]{
		forwardMap:      *redblacktree.NewWith[T, P](keyComparator),
		inverseMap:      *redblacktree.NewWith[P, T](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewWithIntComparators instantiates a bidirectional map with the IntComparator for key and value, i.e. keys and values are of type int.
func NewWithIntComparators[T comparable, P comparable]() *Map[T, P] {
	return NewWith[T, P](utils.IntComparator, utils.IntComparator)
}

// NewWithStringComparators instantiates a bidirectional map with the StringComparator for key and value, i.e. keys and values are of type string.
func NewWithStringComparators[T comparable, P comparable]() *Map[T, P] {
	return NewWith[T, P](utils.StringComparator, utils.StringComparator)
}

// Put inserts element into the map.
func (m *Map[T, P]) Put(key T, value P) {
	if v, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(v)
	}
	if k, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(k)
	}

	m.forwardMap.Put(key, value)
	m.inverseMap.Put(value, key)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[T, P]) Get(key T) (value P, found bool) {
	if d, ok := m.forwardMap.Get(key); ok {
		return d, true
	}
	return utils.AnyEmpty[P](), false
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[T, P]) GetKey(value P) (key T, found bool) {
	if d, ok := m.inverseMap.Get(value); ok {
		return d, true
	}
	return utils.AnyEmpty[T](), false
}

// Remove removes the element from the map by key.
func (m *Map[T, P]) Remove(key T) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[T, P]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[T, P]) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (ordered).
func (m *Map[T, P]) Keys() []T {
	return m.forwardMap.Keys()
}

// Values returns all values (ordered).
func (m *Map[T, P]) Values() []P {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the map.
func (m *Map[T, P]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map[T, P]) String() string {
	str := "TreeBidiMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
