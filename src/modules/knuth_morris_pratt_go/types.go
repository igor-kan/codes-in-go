package mod_knuth_morris_pratt_go

type KmpMatch struct {
	Offsets []int
	Prefix  []int
}

type ModuleMeta struct {
	Slug    string
	Version int
}
