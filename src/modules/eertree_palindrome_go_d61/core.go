package mod_eertree_palindrome_go_d61

func DistinctPalindromes(s string) int {
	n := len(s)
	next := []map[byte]int{{}, {}}
	link := []int{0, 0}
	length := []int{-1, 0}
	last := 1
	nodes := 2
	getLink := func(v, pos int) int {
		for {
			l := length[v]
			if pos-1-l >= 0 && s[pos-1-l] == s[pos] {
				return v
			}
			v = link[v]
		}
	}
	for pos := 0; pos < n; pos++ {
		cur := getLink(last, pos)
		c := s[pos]
		if nxt, ok := next[cur][c]; ok {
			last = nxt
			continue
		}
		newNode := nodes
		nodes++
		next = append(next, map[byte]int{})
		length = append(length, length[cur]+2)
		next[cur][c] = newNode
		if length[newNode] == 1 {
			link = append(link, 1)
		} else {
			cand := getLink(link[cur], pos)
			link = append(link, next[cand][c])
		}
		last = newNode
	}
	return nodes - 2
}

type EertreeCounter struct{}

func (e EertreeCounter) Count(s string) int {
	return DistinctPalindromes(s)
}

var _ PalindromeCounter = EertreeCounter{}
