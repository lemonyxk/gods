// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package btree

import (
	"encoding/json"

	"github.com/emirpasic/gods/containers"
	"github.com/emirpasic/gods/utils"
)

func assertSerializationImplementation[T comparable, P any]() {
	var _ containers.JSONSerializer = (*Tree[T, P])(nil)
	var _ containers.JSONDeserializer = (*Tree[T, P])(nil)
}

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree[T, P]) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree[T, P]) FromJSON(data []byte) error {
	elements := make(map[T]P)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}
