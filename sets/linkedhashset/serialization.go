// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"encoding/json"

	"github.com/lemonyxk/gods/containers"
)

func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = (*Set[T])(nil)
	var _ containers.JSONDeserializer = (*Set[T])(nil)
}

// ToJSON outputs the JSON representation of the set.
func (set *Set[T]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Set[T]) FromJSON(data []byte) error {
	elements := []T{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}
