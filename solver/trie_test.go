package solver

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestTrie_Insert(t *testing.T) {
	type testCase struct {
		name    string
		input   []rune
		wantLen int
	}

	var tests = []testCase{
		{"test", []rune("test"), 1},
		{"duplicates", []rune("boolean"), 1},
		{"empty", []rune{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			root := NewTrie()

			root.Insert(tc.input)
			assert.Equal(t, tc.wantLen, len(root.children))

			var child = root
			var ok bool
			for i := 0; i < len(tc.input); i++ {
				assert.Equal(t, tc.wantLen, len(child.children))
				assert.False(t, child.isEnd)
				child, ok = child.children[tc.input[i]]
				assert.True(t, ok)
			}
			// handle checking for empty word
			if len(tc.input) > 0 {
				// no more children after we iterated over the whole word we inserted
				assert.Equal(t, 0, len(child.children))
				assert.True(t, child.isEnd)
			}
		})
	}

}

func TestTrie_Insert_Multiple(t *testing.T) {

	type testCase struct {
		name           string
		inputs         [][]rune
		overlapPos     int
		wantOverlapLen int
	}

	var tests = []testCase{
		{"simple", [][]rune{[]rune("test"), []rune("toy")}, 1, 2},
		{"three", [][]rune{[]rune("test"), []rune("toy"), []rune("three")}, 1, 3},
		{"empty", [][]rune{[]rune("test"), []rune("toy"), []rune("")}, 1, 2},
		{"mid branch", [][]rune{[]rune("test"), []rune("tesseracts")}, 3, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			root := NewTrie()
			for _, word := range tc.inputs {
				root.Insert(word)
			}
			assert.Equal(t, 1, len(root.children))
			for _, word := range tc.inputs {
				var child = root
				var ok bool
				for i := 0; i < len(word); i++ {
					if i == tc.overlapPos {
						assert.Equal(t, tc.wantOverlapLen, len(child.children))
					} else {
						assert.Equal(t, 1, len(child.children))
					}
					assert.False(t, child.isEnd)
					child, ok = child.children[word[i]]
					assert.True(t, ok)
				}
				// handle checking for empty word
				if len(word) > 0 {
					// no more children after we iterated over the whole word we inserted
					assert.Equal(t, 0, len(child.children))
					assert.True(t, child.isEnd)
				}
			}
		})
	}

}

func TestTrie_Insert_MultipleBranches(t *testing.T) {

	words := [][]rune{[]rune("abc"), []rune("abdefg"), []rune("abdexyz")}
	firstOverlap := 2
	secondOverlap := 4

	root := NewTrie()
	for _, word := range words {
		root.Insert(word)
	}
	assert.Equal(t, 1, len(root.children))
	for _, word := range words {
		var child = root
		var ok bool
		for i := 0; i < len(word); i++ {
			if i == firstOverlap || i == secondOverlap {
				assert.Equal(t, 2, len(child.children))
			} else {
				assert.Equal(t, 1, len(child.children))
			}
			assert.False(t, child.isEnd)
			child, ok = child.children[word[i]]
			assert.True(t, ok)
		}
		// no more children after we iterated over the whole word we inserted
		assert.Equal(t, 0, len(child.children))
		assert.True(t, child.isEnd)
	}
}

func TestTrie_IsPrefix(t *testing.T) {
	word := []rune("testing")
	prefixes := [][]rune{[]rune("t"), []rune("test"), []rune("testb"), []rune("testing")}
	want := []bool{true, true, false, true}

	root := NewTrie()
	root.Insert(word)

	for i := 0; i < len(prefixes); i++ {
		got := root.IsPrefix(prefixes[i])

		assert.Equal(t, want[i], got)
	}
}

func TestTrie_IsPrefix_Multiple(t *testing.T) {
	words := [][]rune{[]rune("testing"), []rune("tesseracts"), []rune("tesselation")}
	prefixes := [][]rune{[]rune("t"), []rune("test"), []rune("testb"), []rune("testing"), []rune("tess"), []rune("tessa"), []rune("tesse"), []rune("tesset")}
	want := []bool{true, true, false, true, true, false, true, false}

	root := NewTrie()
	for _, word := range words {
		root.Insert(word)
	}

	for i := 0; i < len(prefixes); i++ {
		got := root.IsPrefix(prefixes[i])

		assert.Equal(t, want[i], got)
	}
}

func TestTrie_IsWord(t *testing.T) {
	type testCase struct {
		name   string
		words  [][]rune
		search []rune
		want   bool
	}

	var tests = []testCase{
		{"valid short", [][]rune{[]rune("test"), []rune("tesseract")}, []rune("test"), true},
		{"valid long", [][]rune{[]rune("test"), []rune("tesseract")}, []rune("tesseract"), true},
		{"invalid", [][]rune{[]rune("test"), []rune("tesseract")}, []rune("bool"), false},
		{"invalid overrun", [][]rune{[]rune("test"), []rune("tesseract")}, []rune("testing"), false},
		{"invalid only prefix", [][]rune{[]rune("test"), []rune("tesseract")}, []rune("te"), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			root := NewTrie()
			for _, word := range tc.words {
				root.Insert(word)
			}

			got := root.IsWord(tc.search)

			assert.Equal(t, tc.want, got)
		})
	}
}
