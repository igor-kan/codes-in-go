package mod_aho_multi_scan_d59

type acNode struct {
	next [26]int
	fail int
	hits int
}

type AcAutomaton struct {
	nodes []acNode
}

func BuildAc(patterns []string) *AcAutomaton {
	nodes := []acNode{{}}
	for _, p := range patterns {
		cur := 0
		for i := 0; i < len(p); i++ {
			c := int(p[i] - 'a')
			if nodes[cur].next[c] == 0 {
				nodes[cur].next[c] = len(nodes)
				nodes = append(nodes, acNode{})
			}
			cur = nodes[cur].next[c]
		}
		nodes[cur].hits++
	}
	var q []int
	for c := 0; c < 26; c++ {
		if nxt := nodes[0].next[c]; nxt != 0 {
			q = append(q, nxt)
		}
	}
	for qi := 0; qi < len(q); qi++ {
		u := q[qi]
		f := nodes[u].fail
		nodes[u].hits += nodes[f].hits
		for c := 0; c < 26; c++ {
			if v := nodes[u].next[c]; v != 0 {
				nodes[v].fail = nodes[f].next[c]
				q = append(q, v)
			} else {
				nodes[u].next[c] = nodes[f].next[c]
			}
		}
	}
	return &AcAutomaton{nodes: nodes}
}

func (a *AcAutomaton) CountHits(text string) int {
	cur, tot := 0, 0
	for i := 0; i < len(text); i++ {
		c := int(text[i] - 'a')
		if c < 0 || c >= 26 {
			cur = 0
			continue
		}
		cur = a.nodes[cur].next[c]
		tot += a.nodes[cur].hits
	}
	return tot
}
