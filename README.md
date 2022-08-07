[![License](https://img.shields.io/github/license/s0rg/trie)](https://github.com/s0rg/trie/blob/main/LICENSE)
[![Build](https://github.com/s0rg/trie/workflows/ci/badge.svg)](https://github.com/s0rg/trie/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/trie)](https://goreportcard.com/report/github.com/s0rg/trie)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/trie)](https://pkg.go.dev/github.com/s0rg/trie)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/trie)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/trie?sort=semver)](https://github.com/s0rg/trie/tags)


# trie
Generic prefix tree for golang

# example

```go
    import (
        "fmt"
        "log"

        "github.com/s0rg/trie"
    )

    t := trie.New[int]()

    t.Add("bar", 1)
    t.Add("baz", 2)
    t.Add("bat", 3)

    val, err := t.Find("bar")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(val) // will print 1
```
