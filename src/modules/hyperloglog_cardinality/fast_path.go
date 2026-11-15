package mod_hyperloglog_cardinality

// AddBatch adds a contiguous slice of 64-bit items into h.
func AddBatch(h *HLL, items []uint64) {
	for _, x := range items {
		h.Add(x)
	}
}
