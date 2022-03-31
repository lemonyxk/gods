// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treebidimap

import (
	"encoding/json"

	"github.com/lemonyxk/gods/containers"
	"github.com/lemonyxk/gods/utils"
)

func assertSerializationImplementation[T comparable, P comparable]() {
	var _ containers.JSONSerializer = (*Map[T, P])(nil)
	var _ containers.JSONDeserializer = (*Map[T, P])(nil)
}

// ToJSON outputs the JSON representation of the map.
func (m *Map[T, P]) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := m.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[T, P]) FromJSON(data []byte) error {
	elements := make(map[T]P)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.Put(key, value)
		}
	}
	return err
}
