package mod_aho_corasick_go

type Automaton struct {
	nodes []AcNode
}

func CompilePatterns(patterns [][]byte) (*Automaton, error) {
	if err := validatePatterns(patterns); err != nil {
		return nil, err
	}
	nodes := []AcNode{{}}
	for _, pat := range patterns {
		cur := 0
		for _, b := range pat {
			if nodes[cur].next[b] == 0 {
				nodes[cur].next[b] = len(nodes)
				nodes = append(nodes, AcNode{})
			}
			cur = nodes[cur].next[b]
		}
		nodes[cur].hits++
	}
	q := make([]int, 0, len(nodes))
	for c := 0; c < 256; c++ {
		if nxt := nodes[0].next[c]; nxt != 0 {
			q = append(q, nxt)
		}
	}
	for head := 0; head < len(q); head++ {
		v := q[head]
		f := nodes[v].fail
		nodes[v].hits += nodes[f].hits
		for c := 0; c < 256; c++ {
			if ch := nodes[v].next[c]; ch != 0 {
				nodes[ch].fail = nodes[f].next[c]
				q = append(q, ch)
			} else {
				nodes[v].next[c] = nodes[f].next[c]
			}
		}
	}
	return &Automaton{nodes: nodes}, nil
}

func (a *Automaton) CountMatches(text []byte) int {
	return fastCountBytes(a.nodes, text)
}

func (a *Automaton) StateCount() int {
	return len(a.nodes)
}
