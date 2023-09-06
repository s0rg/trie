package gtrie

type node[K comparable, T any] struct {
	children map[K]*node[K, T]
	parent   *node[K, T]
	value    T
	ok       bool
	key      K
}

func makeNode[K comparable, T any](r K) *node[K, T] {
	return &node[K, T]{
		key:      r,
		children: make(map[K]*node[K, T]),
	}
}

func (n *node[K, T]) SetValue(val T) {
	n.value, n.ok = val, true
}

func (n *node[K, T]) DropValue() {
	n.ok = false
}

func (n *node[K, T]) HasValue() (ok bool) {
	return n.ok
}

func (n *node[K, T]) Value() (rv T) {
	return n.value
}

func (n *node[K, T]) AddChild(r K) (rv *node[K, T]) {
	rv = makeNode[K, T](r)

	rv.parent = n
	n.children[r] = rv

	return rv
}

func (n *node[K, T]) GetChild(r K) (rv *node[K, T], ok bool) {
	rv, ok = n.children[r]

	return
}

func (n *node[K, T]) DelChild(r K) {
	if c, ok := n.children[r]; ok {
		delete(n.children, r)

		c.parent = nil
	}
}

func (n *node[K, T]) HasChildren() (ok bool) {
	return len(n.children) > 0
}

func (n *node[K, T]) Parent() (rv *node[K, T]) {
	return n.parent
}
