// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import "github.com/lemonyxk/gods/containers"

func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = (*Heap[T])(nil)
	var _ containers.JSONDeserializer = (*Heap[T])(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap[T]) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap[T]) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}
