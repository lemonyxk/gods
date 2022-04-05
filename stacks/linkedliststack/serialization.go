// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedliststack

import "github.com/lemonyxk/gods/containers"

func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = (*Stack[T])(nil)
	var _ containers.JSONDeserializer = (*Stack[T])(nil)
}

// ToJSON outputs the JSON representation of the stack.
func (stack *Stack[T]) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

// FromJSON populates the stack from the input JSON representation.
func (stack *Stack[T]) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
