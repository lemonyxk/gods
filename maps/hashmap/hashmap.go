// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"fmt"

	"github.com/lemonyxk/gods/maps"
)

func assertMapImplementation[T comparable, P any]() {
	var _ maps.Map[T, P] = (*Map[T, P])(nil)
}

// Map holds the elements in go's native map
type Map[T comparable, P any] struct {
	m map[T]P
}

// New instantiates a hash map.
func New[T comparable, P any]() *Map[T, P] {
	return &Map[T, P]{m: make(map[T]P)}
}

// Put inserts element into the map.
func (m *Map[T, P]) Put(key T, value P) {
	m.m[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[T, P]) Get(key T) (value P, found bool) {
	value, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map[T, P]) Remove(key T) {
	delete(m.m, key)
}

// Empty returns true if map does not contain any elements
func (m *Map[T, P]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[T, P]) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map[T, P]) Keys() []T {
	keys := make([]T, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values (random order).
func (m *Map[T, P]) Values() []P {
	values := make([]P, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[T, P]) Clear() {
	m.m = make(map[T]P)
}

// String returns a string representation of container
func (m *Map[T, P]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
