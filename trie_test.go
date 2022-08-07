package trie

import (
	"strings"
	"testing"
)

func TestTrieFind(t *testing.T) {
	t.Parallel()

	tr := New[int]()

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
		err error
	)

	for _, c := range cases {
		if val, err = tr.Find(c.Path); err != nil {
			t.Fatalf("not found: '%s'", c.Path)
		}

		if val != c.Want {
			t.Fatalf("value not match got: %d want: %d", val, c.Want)
		}
	}

	nonexistent := []string{"ba", "bo", "band", "fan"}

	for _, c := range nonexistent {
		if _, err = tr.Find(c); err == nil {
			t.Fatalf("found non-existent: '%s'", c)
		}
	}
}

func TestTrieDel(t *testing.T) {
	t.Parallel()

	tr := New[int]()

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
		err error
	)

	if _, err = tr.Find(kbar); err != nil {
		t.Fatal("'bar' not found")
	}

	if _, err = tr.Find(kbark); err != nil {
		t.Fatalf("'bark' not found")
	}

	if err = tr.Del(kbar); err != nil {
		t.Fatalf("cannot delete 'bar': %v", err)
	}

	if err = tr.Del(kfoo); err != nil {
		t.Fatalf("cannot delete 'foobar': %v", err)
	}

	if err = tr.Del(kfoo); err == nil {
		t.Fatal("double delete 'foobar'")
	}

	if _, err = tr.Find(kbar); err == nil {
		t.Fatal("'bar' found")
	}

	if val, err = tr.Find(kbark); err != nil {
		t.Fatalf("'bark' not found")
	}

	if val != vbark {
		t.Fatalf("'bark' value mismatch want: %d got: %d", vbark, val)
	}
}

func TestTrieString(t *testing.T) {
	t.Parallel()

	tr := New[int]()

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
