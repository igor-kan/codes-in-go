package mod_reservoir_sample_alg_r

type SampleSnapshot struct {
	Seen  uint64
	Items []int64
}

type ModuleMeta struct {
	Slug    string
	Version int
}
