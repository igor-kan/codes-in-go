package mod_radix_compact_trie_d56

type PrefixNode struct {
	children [26]*PrefixNode
	prefixCt int
	isWord   bool
}

type PrefixTrie struct {
	root *PrefixNode
}

func NewPrefixTrie() *PrefixTrie {
	return &PrefixTrie{root: &PrefixNode{}}
}

func (pt *PrefixTrie) Insert(w string) {
	cur := pt.root
	cur.prefixCt++
	for i := 0; i < len(w); i++ {
		c := int(w[i] - 'a')
		if c < 0 || c >= 26 {
			return
		}
		if cur.children[c] == nil {
			cur.children[c] = &PrefixNode{}
		}
		cur = cur.children[c]
		cur.prefixCt++
	}
	cur.isWord = true
}

func (pt *PrefixTrie) CountPrefix(p string) int {
	cur := pt.root
	for i := 0; i < len(p); i++ {
		c := int(p[i] - 'a')
		if c < 0 || c >= 26 || cur.children[c] == nil {
			return 0
		}
		cur = cur.children[c]
	}
	return cur.prefixCt
}
