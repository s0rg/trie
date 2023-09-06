package gtrie

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Trie represents generic prefix-tree.
type Trie[K comparable, T any] struct {
	root *node[K, T]
}

// New creates empty Trie.
func New[K comparable, T any]() *Trie[K, T] {
	var emptyKey K
	return &Trie[K, T]{root: makeNode[K, T](emptyKey)}
}

// Add inserts new key/value pair.
func (t *Trie[K, T]) Add(key []K, value T) {
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

func (t *Trie[K, T]) find(key []K) (n *node[K, T], ok bool) {
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
func (t *Trie[K, T]) Del(key []K) (ok bool) {
	var n, p *node[K, T]

	if n, ok = t.find(key); !ok {
		return
	}

	for ; n != nil; n = p {
		p = n.Parent()

		if n.HasValue() {
			n.DropValue()
		}

		if n.HasChildren() {
			break
		}

		p.DelChild(n.key)
	}

	return true
}

// Find does tree-lookup in order to find value associated with given key.
func (t *Trie[K, T]) Find(key []K) (value T, ok bool) {
	var n *node[K, T]

	if n, ok = t.find(key); !ok {
		return
	}

	return n.Value(), n.HasValue()
}

// Suggest returns slice of existing keys with matching prefix.
func (t *Trie[K, T]) Suggest(prefix []K) (rv [][]K, ok bool) {
	add := func(v []K, _ T) {
		rv = append(rv, v)
	}

	t.Iter(prefix, add)

	return rv, len(rv) > 0
}

// Iter iterates over trie by prefix using dfs.
// Pass prefix="" to iterate over whole trie.
func (t *Trie[K, T]) Iter(prefix []K, walker func(key []K, value T)) {
	var n *node[K, T]
	n, ok := t.find(prefix)
	if !ok {
		return
	}

	if n.HasValue() {
		walker(prefix, n.value)
	}

	dfsKeys(n, prefix, walker)
}

// String implements fmt.Stringer interface.
func (t *Trie[K, T]) String() string {
	var b bytes.Buffer

	writeNode(&b, t.root, 0)

	return b.String()
}

func writeNode[K comparable, T any](
	w io.Writer,
	n *node[K, T],
	level int,
) {
	var empty K
	template := "key: '%c'"

	if level > 0 {
		template = strings.Repeat("\t", level) + template
	}

	switch {
	case n.HasValue():
		_, _ = fmt.Fprintf(w, template+" value: '%v'", n.key, n.Value())
	case n.key == empty:
		_, _ = fmt.Fprint(w, "root")
	default:
		_, _ = fmt.Fprintf(w, template+":", n.key)
	}

	_, _ = fmt.Fprintln(w)

	for _, c := range n.children {
		writeNode(w, c, level+1)
	}
}

func dfsKeys[K comparable, T any](
	n *node[K, T],
	prefix []K,
	handler func(key []K, value T),
) {
	for r, c := range n.children {
		key := append(prefix, r)

		if c.HasValue() {
			handler(key, c.value)
		}

		dfsKeys(c, key, handler)
	}
}
