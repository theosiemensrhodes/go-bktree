// Copyright 2014 Mahmud Ridwan. All rights reserved.

// Package bktree provides an implementation of BK-tree (http://en.wikipedia.org/wiki/BK-tree).
//
// A BK-tree is a metric tree suggested by Walter Austin Burkhard and Robert M. Keller specifically adapted to discrete metric spaces.
package bktree

// The Metric type is a function used by BK-tree instances to measure the distance between two given strings.
type Metric func(a, b string) int

// BKTree represents a BK-tree with a given metric function.
type BKTree struct {
	Metric Metric // Metric function, required
	root   *node
}

// New returns an initialized BK-tree.
func New(m Metric) *BKTree {
	return &BKTree{
		Metric: m,
	}
}

// Add inserts a new word to the BK-tree.
func (t *BKTree) Add(w string) {
	if t.root == nil {
		t.root = &node{w, make(map[int]*node)}
	} else {
		t.root.add(w, t.Metric)
	}
}

// Find returns all the words in the BK-tree with a distance of n from w.
func (t *BKTree) Find(w string, n int) []string {
	r := []string{}
	if t.root != nil {
		r = t.root.find(w, n, t.Metric, r)
	}
	return r
}

type node struct {
	word   string
	childs map[int]*node
}

func (e *node) add(w string, m Metric) {
	d := m(e.word, w)
	if c, ok := e.childs[d]; !ok {
		e.childs[d] = &node{w, make(map[int]*node)}
	} else {
		c.add(w, m)
	}
}

func (e *node) find(w string, n int, m Metric, r []string) []string {
	l := m(e.word, w)
	if l <= n {
		r = append(r, e.word)
	}
	for i := l - n; i <= l+n; i++ {
		if i < 0 {
			continue // Skip negative distances
		}
		if c, ok := e.childs[i]; ok {
			r = c.find(w, n, m, r)
		}
	}
	return r
}
