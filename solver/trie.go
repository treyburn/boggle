package solver

type Trie struct {
	isEnd    bool
	isRoot   bool
	children map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		isRoot:   true,
		children: make(map[rune]*Trie),
	}
}

func newChild() *Trie {
	return &Trie{
		isRoot:   false,
		children: make(map[rune]*Trie),
	}
}

func (t *Trie) Insert(word []rune) {
	if len(word) < 1 {
		if !t.isRoot {
			t.isEnd = true
		}
		return
	}

	child, ok := t.children[word[0]]
	if !ok {
		child = newChild()
		t.children[word[0]] = child
	}

	child.Insert(word[1:])
}

func (t *Trie) IsPrefix(word []rune) bool {
	if len(word) < 1 {
		return true
	}
	child, ok := t.children[word[0]]
	if !ok {
		return false
	}

	return child.IsPrefix(word[1:])
}

func (t *Trie) IsWord(word []rune) bool {
	if len(word) < 1 {
		if t.isEnd {
			return true
		}
		return false
	}
	child, ok := t.children[word[0]]
	if !ok {
		return false
	}

	return child.IsWord(word[1:])
}
