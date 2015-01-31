package bktree

type Metric func(a, b string) int

type BKTree struct {
	Metric Metric

	root *node
}

func New(m Metric) *BKTree {
	return &BKTree{
		Metric: m,
	}
}

func (t *BKTree) Add(s string) {
	if t.root == nil {
		t.root = &node{s, make(map[int]*node)}
	} else {
		t.root.add(s, t.Metric)
	}
}

func (t *BKTree) Find(s string, n int) []string {
	r := []string{}
	if t.root != nil {
		r = t.root.find(s, n, -1, t.Metric, r)
	}
	return r
}

type node struct {
	word   string
	childs map[int]*node
}

func (e *node) add(s string, m Metric) {
	d := m(e.word, s)
	if c, ok := e.childs[d]; !ok {
		e.childs[d] = &node{s, make(map[int]*node)}
	} else {
		c.add(s, m)
	}
}

func (e *node) find(s string, n, d int, m Metric, r []string) []string {
	l := m(e.word, s)
	if l <= n {
		r = append(r, e.word)
	}
	if d == -1 {
		d = l
	}
	for i := n - d; i <= n+d; i++ {
		if c, ok := e.childs[i]; ok {
			r = c.find(s, n, d, m, r)
		}
	}
	return r
}
