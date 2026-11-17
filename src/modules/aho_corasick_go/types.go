package mod_aho_corasick_go

type AcNode struct {
	next  [256]int
	fail  int
	hits  int
	depth int
}
