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
