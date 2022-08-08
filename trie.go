package trie

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

const rootNode = rune(0)

// Trie represents generic prefix-tree.
type Trie[T any] struct {
	root *node[T]
}

// New creates empty Trie.
func New[T any]() *Trie[T] {
	return &Trie[T]{root: makeNode[T](rootNode)}
}

// Add inserts new key/value pair.
func (t *Trie[T]) Add(key string, value T) {
	n := t.root

	for _, r := range key {
		if c, ok := n.GetChild(r); ok {
			n = c

			continue
		}

		n = n.AddChild(r)
	}

	n.SetValue(value)
}

func (t *Trie[T]) find(key string) (n *node[T], ok bool) {
	n = t.root

	for _, r := range key {
		if n, ok = n.GetChild(r); ok {
			continue
		}

		return
	}

	return n, true
}

// Del removes node by key.
func (t *Trie[T]) Del(key string) (ok bool) {
	var n, p *node[T]

	if n, ok = t.find(key); !ok {
		return
	}

	for ; n != nil; n = p {
		p = n.Parent()

		if n.HasValue() {
			n.DropValue()
		}

		if n.HasChilds() {
			break
		}

		p.DelChild(n.key)
	}

	return true
}

// Find does tree-lookup in order to find value associated with given key.
func (t *Trie[T]) Find(key string) (value T, ok bool) {
	var n *node[T]

	if n, ok = t.find(key); !ok {
		return
	}

	return n.Value(), n.HasValue()
}

// Suggest returns slice of existing keys with matching prefix.
func (t *Trie[T]) Suggest(prefix string) (rv []string, ok bool) {
	var n *node[T]

	if n, ok = t.find(prefix); !ok {
		return
	}

	add := func(v string) {
		rv = append(rv, v)
	}

	if n.HasValue() {
		add(prefix)
	}

	dfsKeys(n, prefix, add)

	return rv, len(rv) > 0
}

// String implements fmt.Stringer interface.
func (t *Trie[T]) String() string {
	var b bytes.Buffer

	writeNode(&b, t.root, 0)

	return b.String()
}

func writeNode[T any](
	w io.Writer,
	n *node[T],
	level int,
) {
	template := "key: '%c'"

	if level > 0 {
		template = strings.Repeat("\t", level) + template
	}

	switch {
	case n.HasValue():
		fmt.Fprintf(w, template+" value: '%v'", n.key, n.Value())
	case n.key == rootNode:
		fmt.Fprint(w, "root")
	default:
		fmt.Fprintf(w, template+":", n.key)
	}

	fmt.Fprintln(w)

	for _, c := range n.childs {
		writeNode(w, c, level+1)
	}
}

func dfsKeys[T any](
	n *node[T],
	prefix string,
	handler func(string),
) {
	for r, c := range n.childs {
		key := prefix + string(r)

		if c.HasValue() {
			handler(key)
		}

		dfsKeys(c, key, handler)
	}
}
