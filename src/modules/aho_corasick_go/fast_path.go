package mod_aho_corasick_go

func fastCountBytes(nodes []AcNode, text []byte) int {
	state := 0
	total := 0
	for _, b := range text {
		state = nodes[state].next[b]
		total += nodes[state].hits
	}
	return total
}
