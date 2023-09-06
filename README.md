[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/trie)](https://pkg.go.dev/github.com/s0rg/trie)
[![License](https://img.shields.io/github/license/s0rg/trie)](https://github.com/s0rg/trie/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/trie)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/trie?sort=semver)](https://github.com/s0rg/trie/tags)

[![CI](https://github.com/s0rg/trie/workflows/ci/badge.svg)](https://github.com/s0rg/trie/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/trie)](https://goreportcard.com/report/github.com/s0rg/trie)
[![Maintainability](https://api.codeclimate.com/v1/badges/b476ce7fd7bbaa30e5a6/maintainability)](https://codeclimate.com/github/s0rg/trie/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b476ce7fd7bbaa30e5a6/test_coverage)](https://codeclimate.com/github/s0rg/trie/test_coverage)


# Trie

Generic prefix tree for golang  
Has 2 versions:
- trie: string keyed Trie. Keys are strings, value are generic. Use case: text prefix search. 
- gtrie: generic keyed Trie. Key can be any comparable value: byte, string, int etc. Use case: OS path searches. 

# Example Trie

```go
    package main

    import (
        "fmt"
        "log"

        "github.com/s0rg/trie"
    )

    func main() {
        t := trie.New[int]()

        t.Add("bar", 1)
        t.Add("baz", 2)
        t.Add("bat", 3)

        val, ok := t.Find("bar")
        if !ok {
            log.Fatal("not found")
        }

        fmt.Println(val) // will print 1
    }
```

# Example GTrie

```go
    package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/s0rg/trie/gtrie"
)

func main() {
	walker := func(key []string, value int) {
		fmt.Printf("%s, %d", strings.Join(key, "/"), value)
	}

	tr := gtrie.New[string, int]()

	tr.Add([]string{"home", "teadove", "Documents"}, 1)
	tr.Add([]string{"home", "teadove", "Downloads"}, 2)
	tr.Add([]string{"home", "teadove", "Pictures"}, 3)
	tr.Add([]string{"home", "tainella", "Documents"}, 4)

	tr.Iter([]string{"home", "teadove"}, walker)
	// Will print
	// home/teadove/Documets 1
	// home/teadove/Downloads 2
	// home/teadove/Pictures 3
}
```

## Methods
### Trie
```go
func New() *Trie[T]
func (t *Trie[T]) Add(key string, value T)
func (t *Trie[T]) Del(key string) (ok bool)
func (t *Trie[T]) Find(key string) (value T, ok bool)
func (t *Trie[T]) String() string
func (t *Trie[T]) Suggest(prefix string) (rv []string, ok bool)
```
### GTrie
```go
func New() *GTrie[K, T]
func (t *Trie[K, T]) Add(key []K, value T)
func (t *Trie[K, T]) Del(key []K) (ok bool)
func (t *Trie[K, T]) Find(key []K) (value T, ok bool)
func (t *Trie[K, T]) String() string
func (t *Trie[K, T]) Suggest(prefix []K) (rv [][]K, ok bool)
```
