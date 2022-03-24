// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// All data structures must implement the container structure

package containers

import (
	"testing"

	"github.com/emirpasic/gods/utils"
)

// For testing purposes
type ContainerTest[P any] struct {
	values []P
}

func (container ContainerTest[P]) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest[P]) Size() int {
	return len(container.values)
}

func (container ContainerTest[P]) Clear() {
	container.values = []P{}
}

func (container ContainerTest[P]) Values() []P {
	return container.values
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest[int]{}
	container.values = []int{5, 1, 3, 2, 4}
	values := GetSortedValues[int](container, utils.IntComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := ContainerTest[string]{}
	container.values = []string{"g", "a", "d", "e", "f", "c", "b"}
	values := GetSortedValues[string](container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}
