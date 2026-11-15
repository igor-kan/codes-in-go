package mod_hyperloglog_cardinality

// HLLConfig holds precision configuration for a HyperLogLog sketch.
type HLLConfig struct {
	Precision uint8
	Seed      uint64
}

// HLL stores 2^Precision byte registers tracking max leading-zero runs.
type HLL struct {
	Precision uint8
	Seed      uint64
	Registers []uint8
}
