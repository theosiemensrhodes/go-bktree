// Copyright 2014 Mahmud Ridwan. All rights reserved.

// Package bktree provides an implementation of BK-tree (http://en.wikipedia.org/wiki/BK-tree).
//
// A BK-tree is a metric tree suggested by Walter Austin Burkhard and Robert M. Keller specifically adapted to discrete metric spaces.
package bktree

import (
	"os"

	"github.com/gogo/protobuf/proto"
)

// The Metric type is a function used by BK-tree instances to measure the distance between two given strings.
type Metric func(a, b []byte) int

// BKTree represents a BK-tree with a given metric function.
type BKTree struct {
	Metric Metric // Metric function, required
	root   *Node
	dirty  bool
}

// New returns an initialized BK-tree.
func New(m Metric) *BKTree {
	return &BKTree{
		Metric: m,
		dirty:  false,
	}
}

// Reads data from file and deserialize into tree
func (t *BKTree) ReadFromFile(dbFile string) (err error) {
	data, err := os.ReadFile(dbFile)
	if err != nil {
		return
	}
	root := &Node{}
	err = proto.Unmarshal(data, root)
	if err != nil {
		return
	}

	t.root = root

	return
}

// Serializes data and saves into file
// If tree is empty no operation will be made and 'saved' parameter returns false.
func (t *BKTree) SaveToFile(filePath string) (saved bool, err error) {
	saved = false
	if t.root != nil {
		var data []byte
		data, err = proto.Marshal(t.root)
		if err != nil {
			return
		}
		err = os.WriteFile(filePath, data, 0644)
		if err != nil {
			return
		}
		t.dirty = false
		saved = true
	}
	return

}

// Add inserts a new word to the BK-tree.
func (t *BKTree) Add(data []byte) {
	if t.root == nil {
		t.root = &Node{data, make(map[int64]*Node)}
	} else {
		t.root.Add(data, t.Metric)
	}
	t.dirty = true
}

// Find returns all the words in the BK-tree with a distance of n from w.
func (t *BKTree) Find(data []byte, n int64) [][]byte {
	r := [][]byte{}
	if t.root != nil {
		r = t.root.Find(data, n, t.Metric, r)
	}
	return r
}

func (e *Node) Add(data []byte, m Metric) {
	d := int64(m(e.Data, data))
	if c, ok := e.Children[d]; !ok {
		e.Children[d] = &Node{data, make(map[int64]*Node)}
	} else {
		c.Add(data, m)
	}
}

func (e *Node) Find(data []byte, n int64, m Metric, r [][]byte) [][]byte {
	l := int64(m(e.Data, data))
	if l <= n {
		r = append(r, e.Data)
	}
	for i := l - n; i <= l+n; i++ {
		if i < 0 {
			continue // Skip negative distances
		}
		if c, ok := e.Children[i]; ok {
			r = c.Find(data, n, m, r)
		}
	}
	return r
}
