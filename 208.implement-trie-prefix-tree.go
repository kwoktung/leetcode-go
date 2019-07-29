package leetcode_go

type Trie struct {
	Prefix byte
	Children map[byte]*Trie
	End bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	if len(word) == 0 { return }
	b, remain := word[0], word[1:]
	if child, ok := t.Children[b]; ok {
		if len(remain) > 0 {
			child.Insert(remain)
		} else {
			child.End = true
		}
	} else {
		child := &Trie{
			Prefix: b,
			End: len(remain) == 0,
		}
		if len(remain) > 0 {
			child.Insert(remain)
		}
		if t.Children == nil {
			t.Children = map[byte]*Trie{}
		}
		t.Children[b] = child
	}
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	current := t
	for current != nil && len(word) > 0  {
		current = current.Children[word[0]]
		word = word[1:]
	}
	return current != nil && current.End
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	current := t
	for current != nil && len(prefix) > 0  {
		current = current.Children[prefix[0]]
		prefix = prefix[1:]
	}
	return current != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */