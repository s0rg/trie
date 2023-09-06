package trie

type node[T any] struct {
	childs map[rune]*node[T]
	parent *node[T]
	value  T
	ok     bool
	key    rune
}

func makeNode[T any](r rune) *node[T] {
	return &node[T]{
		key:    r,
		childs: make(map[rune]*node[T]),
	}
}

func (n *node[T]) SetValue(val T) {
	n.value, n.ok = val, true
}

func (n *node[T]) DropValue() {
	n.ok = false
}

func (n *node[T]) HasValue() (ok bool) {
	return n.ok
}

func (n *node[T]) Value() (rv T) {
	return n.value
}

func (n *node[T]) AddChild(r rune) (rv *node[T]) {
	rv = makeNode[T](r)

	rv.parent = n
	n.childs[r] = rv

	return rv
}

func (n *node[T]) GetChild(r rune) (rv *node[T], ok bool) {
	rv, ok = n.childs[r]

	return
}

func (n *node[T]) DelChild(r rune) {
	if c, ok := n.childs[r]; ok {
		delete(n.childs, r)

		c.parent = nil
	}
}

func (n *node[T]) HasChildren() (ok bool) {
	return len(n.childs) > 0
}

func (n *node[T]) Parent() (rv *node[T]) {
	return n.parent
}
