package mod_trie_wildcard_d59

type WtNode struct {
	ch   [26]*WtNode
	word bool
}

type WildTrie struct {
	root *WtNode
}

func NewWildTrie() *WildTrie { return &WildTrie{root: &WtNode{}} }

func (w *WildTrie) Add(word string) {
	cur := w.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if cur.ch[c] == nil {
			cur.ch[c] = &WtNode{}
		}
		cur = cur.ch[c]
	}
	cur.word = true
}

func (w *WildTrie) Search(pat string) bool {
	var match func(n *WtNode, i int) bool
	match = func(n *WtNode, i int) bool {
		if n == nil {
			return false
		}
		if i == len(pat) {
			return n.word
		}
		if pat[i] == '.' {
			for _, nxt := range n.ch {
				if nxt != nil && match(nxt, i+1) {
					return true
				}
			}
			return false
		}
		return match(n.ch[pat[i]-'a'], i+1)
	}
	return match(w.root, 0)
}
