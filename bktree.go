// Copyright 2014 Mahmud Ridwan. All rights reserved.

// Package bktree provides an implementation of BK-tree (http://en.wikipedia.org/wiki/BK-tree).
//
// A BK-tree is a metric tree suggested by Walter Austin Burkhard and Robert M. Keller specifically adapted to discrete metric spaces.
package bktree

import (
	"io/ioutil"
	"github.com/gogo/protobuf/proto"
)

// The Metric type is a function used by BK-tree instances to measure the distance between two given strings.
type Metric func(a, b string) int

// BKTree represents a BK-tree with a given metric function.
type BKTree struct {
	Metric Metric // Metric function, required
	root   *Node
}

// New returns an initialized BK-tree.
func New(m Metric) *BKTree {
	return &BKTree{
		Metric: m,
	}
}

// Read data structures from file
func (t *BKTree) ReadFromFile(dbFile string) (err error) {
	data, err := ioutil.ReadFile(dbFile)
	if err != nil {return}
	root := &Node{}
	err = proto.Unmarshal(data, root)
	if err != nil {return}

	t.root = root

	return
}

// Serialize into file
func (t *BKTree) SaveToFile(filePath string) error {
	data, err := proto.Marshal(t.root)
	if err != nil {return err}
	return ioutil.WriteFile(filePath, data, 0644)
}

// Add inserts a new word to the BK-tree.
func (t *BKTree) Add(w string) {
	if t.root == nil {
		t.root = &Node{w, make(map[int64]*Node)}
	} else {
		t.root.Add(w, t.Metric)
	}
}

// Find returns all the words in the BK-tree with a distance of n from w.
func (t *BKTree) Find(w string, n int64) []string {
	r := []string{}
	if t.root != nil {
		r = t.root.Find(w, n, -1, t.Metric, r)
	}
	return r
}


func (e *Node) Add(w string, m Metric) {
	d := int64(m(e.Word, w))
	if c, ok := e.Children[d]; !ok {
		e.Children[d] = &Node{w, make(map[int64]*Node)}
	} else {
		c.Add(w, m)
	}
}

func (e *Node) Find(w string, n, d int64, m Metric, r []string) []string {
	l := int64(m(e.Word, w))
	if l <= n {
		r = append(r, e.Word)
	}
	if d == -1 {
		d = l
	}
	for i := n - d; i <= n+d; i++ {
		if c, ok := e.Children[i]; ok {
			r = c.Find(w, n, d, m, r)
		}
	}
	return r
}

