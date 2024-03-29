package trie_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/s0rg/trie"
)

func TestTrieFind(t *testing.T) {
	t.Parallel()

	tr := trie.New[int]()

	type testCase struct {
		Path string
		Want int
	}

	cases := []testCase{
		{"ban", 1},
		{"banana", 2},
		{"boo", 3},
		{"bandana", 4},
		{"foo", 5},
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

	nonexistent := []string{"ba", "bo", "band", "fan"}

	for _, c := range nonexistent {
		if _, ok = tr.Find(c); ok {
			t.Fatalf("found non-existent: '%s'", c)
		}
	}
}

func TestTrieDel(t *testing.T) {
	t.Parallel()

	tr := trie.New[int]()

	const (
		kbar  = "bar"
		kbark = "bark"
		kfoo  = "foobar"
		vbark = 3
	)

	tr.Add(kbar, 1)
	tr.Add("baz", 2)
	tr.Add(kbark, vbark)
	tr.Add("boo", 4)
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

func TestTrieDelAll(t *testing.T) {
	t.Parallel()

	tr := trie.New[int]()

	tests := []struct {
		k string
		v int
	}{
		{"bar", 1},
		{"baz", 2},
		{"boo", 5},
		{"foobar", 5},
		{"bark", 3},
	}

	for _, t := range tests {
		tr.Add(t.k, t.v)
	}

	for _, test := range tests {
		if _, ok := tr.Find(test.k); !ok {
			t.Fatalf("'%s' not found", test.k)
		}

		if ok := tr.Del(test.k); !ok {
			t.Fatalf("cannot delete '%s'", test.k)
		}

		if _, ok := tr.Find(test.k); ok {
			t.Fatalf("'%s' found", test.k)
		}
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

	walker := func(key string, value int) {
		result[key] = value
	}

	tr := trie.New[int]()

	tr.Add("arc", 1)
	tr.Add("bak", 2)
	tr.Add("bar", 3)
	tr.Add("boo", 4)

	tr.Iter("b", walker)

	if result["bak"] != 2 {
		t.Fatal("key not found")
	}

	if result["bar"] != 3 {
		t.Fatal("key not found")
	}

	if result["boo"] != 4 {
		t.Fatal("key not found")
	}

	if _, ok := result["arc"]; ok {
		t.Fatal("key found, but don't need to exists")
	}
}

func TestTrieSuggest(t *testing.T) {
	t.Parallel()

	tr := trie.New[int]()

	tr.Add("arc", 1)
	tr.Add("bak", 2)
	tr.Add("bar", 3)
	tr.Add("boo", 4)

	var (
		res []string
		ok  bool
	)

	if _, ok = tr.Suggest("c"); ok {
		t.Fatal("suggested c")
	}

	if res, ok = tr.Suggest("a"); !ok {
		t.Fatal("not found a")
	}

	if len(res) != 1 {
		t.Fatal("suggest(a) != 1")
	}

	if res, ok = tr.Suggest("b"); !ok {
		t.Fatal("not found b")
	}

	if len(res) != 3 {
		t.Fatal("suggest(b) != 3")
	}

	if res, ok = tr.Suggest("ba"); !ok {
		t.Fatal("not found ba")
	}

	if len(res) != 2 {
		t.Fatal("suggest(ba) != 2")
	}

	if res, ok = tr.Suggest("bak"); !ok {
		t.Fatal("not found bak")
	}

	if len(res) != 1 {
		t.Fatal("suggest(bak) != 1")
	}
}

func TestTrieCommons(t *testing.T) {
	t.Parallel()

	tr := trie.New[int]()

	if res := tr.Common("foo", 3); len(res) > 0 {
		t.FailNow()
	}

	tr.Add("foo", 1)
	tr.Add("food", 2)
	tr.Add("car", 3)
	tr.Add("carpet", 4)
	tr.Add("cart", 5)
	tr.Add("cartridge", 9)
	tr.Add("probe", 10)
	tr.Add("problem", 11)
	tr.Add("probability", 12)

	want := []string{
		"car",
		"prob",
		"foo",
	}

	if res := tr.Common("", 3); slices.Compare(res, want) != 0 {
		t.Fail()
	}

	if res := tr.Common("ca", 3); len(res) != 1 || res[0] != "car" {
		t.Fail()
	}
}

func TestTrieCommons1(t *testing.T) {
	t.Parallel()

	tr := trie.New[int]()

	tr.Add("a1", 1)
	tr.Add("a2", 2)
	tr.Add("a3", 3)
	tr.Add("b1", 4)
	tr.Add("b2", 5)
	tr.Add("b3", 9)
	tr.Add("c1", 10)
	tr.Add("c2", 11)
	tr.Add("c3", 12)

	if res := tr.Common("", 1); len(res) != 3 {
		t.Fail()
	}
}

func FuzzTrie(f *testing.F) {
	f.Add("foo:F,bar:B")

	f.Fuzz(func(t *testing.T, input string) {
		input = strings.ToValidUTF8(input, "")
		tr := trie.New[string]()
		m := make(map[string]string)

		for _, p := range strings.Split(input, ",") {
			key, val, _ := strings.Cut(p, ":")
			m[key] = val
			tr.Add(key, val)
		}

		for k, v := range m {
			if got, ok := tr.Find(k); !ok || got != v {
				t.Errorf("key %q, want %q, got %q", k, v, got)
			}
		}
	})
}
