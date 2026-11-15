package mod_swiss_probe_hashmap

// MatchGroupH2Packed scans 8 control bytes in parallel using SWAR bit tricks.
func MatchGroupH2Packed(word uint64, h2 uint8) uint64 {
	const lo uint64 = 0x0101010101010101
	const hi uint64 = 0x8080808080808080
	x := word ^ (lo * uint64(h2))
	return ((x - lo) & ^x & hi)
}
