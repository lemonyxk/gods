// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"github.com/lemonyxk/gods/containers"
	rbt "github.com/lemonyxk/gods/trees/redblacktree"
	"github.com/lemonyxk/gods/utils"
)

func assertEnumerableImplementation[T comparable, P any]() {
	var _ containers.EnumerableWithKey[T, P] = (*Map[T, P])(nil)
}

// Each calls the given function once for each element, passing that element's key and value.
func (m *Map[T, P]) Each(f func(key T, value P)) {
	iterator := m.Iterator()
	for iterator.Next() {
		f(iterator.Key(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a container
// containing the values returned by the given function as key/value pairs.
func (m *Map[T, P]) Map(f func(key1 T, value1 P) (T, P)) *Map[T, P] {
	newMap := &Map[T, P]{tree: rbt.NewWith[T, P](m.tree.Comparator)}
	iterator := m.Iterator()
	for iterator.Next() {
		key2, value2 := f(iterator.Key(), iterator.Value())
		newMap.Put(key2, value2)
	}
	return newMap
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (m *Map[T, P]) Select(f func(key T, value P) bool) *Map[T, P] {
	newMap := &Map[T, P]{tree: rbt.NewWith[T, P](m.tree.Comparator)}
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			newMap.Put(iterator.Key(), iterator.Value())
		}
	}
	return newMap
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (m *Map[T, P]) Any(f func(key T, value P) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (m *Map[T, P]) All(f func(key T, value P) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if !f(iterator.Key(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (key,value) for which the function is true or nil,nil otherwise if no element
// matches the criteria.
func (m *Map[T, P]) Find(f func(key T, value P) bool) (T, P) {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return iterator.Key(), iterator.Value()
		}
	}
	return utils.AnyEmpty[T](), utils.AnyEmpty[P]()
}
