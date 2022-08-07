package trie

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

const rootNode = rune(0)

var (
	ErrNotFound = errors.New("not found")
	ErrNoValue  = errors.New("empty value")
)

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

func (t *Trie[T]) find(key string) (n *node[T], err error) {
	n = t.root

	for _, r := range key {
		if c, ok := n.GetChild(r); ok {
			n = c

			continue
		}

		return nil, ErrNotFound
	}

	return n, nil
}

// Del removes node by key.
func (t *Trie[T]) Del(key string) (err error) {
	var n, p *node[T]

	if n, err = t.find(key); err != nil {
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

	return nil
}

// Find does tree-lookup in order to find value associated with given key.
func (t *Trie[T]) Find(key string) (value T, err error) {
	var n *node[T]

	if n, err = t.find(key); err != nil {
		return
	}

	if !n.HasValue() {
		err = ErrNoValue

		return
	}

	return n.Value(), nil
}

// String implements fmt.Stringer interface.
func (t *Trie[T]) String() string {
	var b bytes.Buffer

	writeNode(&b, t.root, 0)

	return b.String()
}

func writeNode[T any](w io.Writer, n *node[T], level int) {
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
