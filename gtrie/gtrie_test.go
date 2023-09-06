package gtrie_test

import (
	"github.com/s0rg/trie"
	"github.com/s0rg/trie/gtrie"
	"strings"
	"testing"
)

func TestTrieFind(t *testing.T) {
	t.Parallel()

	tr := gtrie.New[byte, int]()

	type testCase struct {
		Path []byte
		Want int
	}

	cases := []testCase{
		{[]byte("ban"), 1},
		{[]byte("banana"), 2},
		{[]byte("boo"), 3},
		{[]byte("bandana"), 4},
		{[]byte("foo"), 5},
	}

	for _, c := range cases {
		tr.Add(c.Path, c.Want)
	}

	var (
		val int
		ok  bool
	)

	for _, c := range cases {
		if val, ok = tr.Find(c.Path); !ok {
			t.Fatalf("not found: '%s'", c.Path)
		}

		if val != c.Want {
			t.Fatalf("value not match got: %d want: %d", val, c.Want)
		}
	}

	nonexistent := [][]byte{[]byte("ba"), []byte("bo"), []byte("band"), []byte("fan")}

	for _, c := range nonexistent {
		if _, ok = tr.Find(c); ok {
			t.Fatalf("found non-existent: '%s'", c)
		}
	}
}

func TestTrieDel(t *testing.T) {
	t.Parallel()

	tr := gtrie.New[byte, int]()

	var (
		kbar  = []byte("bar")
		kbark = []byte("bark")
		kfoo  = []byte("foobar")
		vbark = 3
	)

	tr.Add(kbar, 1)
	tr.Add([]byte("baz"), 2)
	tr.Add(kbark, vbark)
	tr.Add([]byte("boo"), 4)
	tr.Add(kfoo, 5)

	var (
		val int
		ok  bool
	)

	if _, ok = tr.Find(kbar); !ok {
		t.Fatal("'bar' not found")
	}

	if _, ok = tr.Find(kbark); !ok {
		t.Fatalf("'bark' not found")
	}

	if ok = tr.Del(kbar); !ok {
		t.Fatal("cannot delete 'bar'")
	}

	if ok = tr.Del(kfoo); !ok {
		t.Fatal("cannot delete 'foobar'")
	}

	if ok = tr.Del(kfoo); ok {
		t.Fatal("double delete 'foobar'")
	}

	if _, ok = tr.Find(kbar); ok {
		t.Fatal("'bar' found")
	}

	if val, ok = tr.Find(kbark); !ok {
		t.Fatalf("'bark' not found")
	}

	if val != vbark {
		t.Fatalf("'bark' value mismatch want: %d got: %d", vbark, val)
	}
}

func TestTrieString(t *testing.T) {
	t.Parallel()

	tr := trie.New[int]()

	tr.Add("cup", 1)
	tr.Add("cub", 2)
	tr.Add("cab", 3)
	tr.Add("cop", 4)

	const wantKeys = 8

	res := tr.String()

	if got := strings.Count(res, "key"); got != wantKeys {
		t.Fatalf("unexpected keys count want: %d got: %d", wantKeys, got)
	}
}

func TestTrieWalk(t *testing.T) {
	t.Parallel()

	result := make(map[string]int, 10)

	walker := func(key []string, value int) {
		result[strings.Join(key, "/")] = value
	}

	tr := gtrie.New[string, int]()

	tr.Add([]string{"home", "teadove", "Documents"}, 1)
	tr.Add([]string{"home", "teadove", "Downloads"}, 2)
	tr.Add([]string{"home", "teadove", "Pictures"}, 3)
	tr.Add([]string{"home", "tainella", "Documents"}, 4)

	tr.Iter([]string{"home", "teadove"}, walker)

	if result["home/teadove/Documents"] != 1 {
		t.Fatal("key not found")
	}
	if result["home/teadove/Downloads"] != 2 {
		t.Fatal("key not found")
	}
	if result["home/teadove/Pictures"] != 3 {
		t.Fatal("key not found")
	}

	_, ok := result["home/tainella/Documents"]
	if ok {
		t.Fatal("key found, but don't need to exists")
	}
}

func TestTrieSuggest(t *testing.T) {
	t.Parallel()

	tr := gtrie.New[byte, int]()

	tr.Add([]byte("arc"), 1)
	tr.Add([]byte("bak"), 2)
	tr.Add([]byte("bar"), 3)
	tr.Add([]byte("boo"), 4)

	var (
		res [][]byte
		ok  bool
	)

	if _, ok = tr.Suggest([]byte("c")); ok {
		t.Fatal("suggested c")
	}

	if res, ok = tr.Suggest([]byte("a")); !ok {
		t.Fatal("not found a")
	}

	if len(res) != 1 {
		t.Fatal("suggest(a) != 1")
	}

	if res, ok = tr.Suggest([]byte("b")); !ok {
		t.Fatal("not found b")
	}

	if len(res) != 3 {
		t.Fatal("suggest(b) != 3")
	}

	if res, ok = tr.Suggest([]byte("ba")); !ok {
		t.Fatal("not found ba")
	}

	if len(res) != 2 {
		t.Fatal("suggest(ba) != 2")
	}

	if res, ok = tr.Suggest([]byte("bak")); !ok {
		t.Fatal("not found bak")
	}

	if len(res) != 1 {
		t.Fatal("suggest(bak) != 1")
	}
}

func FuzzTrie(f *testing.F) {
	f.Add("foo:F,bar:B")

	f.Fuzz(func(t *testing.T, input string) {
		input = strings.ToValidUTF8(input, "")
		tr := gtrie.New[byte, string]()
		m := make(map[string]string)

		for _, p := range strings.Split(input, ",") {
			key, val, _ := strings.Cut(p, ":")
			m[key] = val
			tr.Add([]byte(key), val)
		}

		for k, v := range m {
			if got, ok := tr.Find([]byte(k)); !ok || got != v {
				t.Errorf("key %q, want %q, got %q", k, v, got)
			}
		}
	})
}
