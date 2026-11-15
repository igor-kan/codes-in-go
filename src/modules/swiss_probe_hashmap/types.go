package mod_swiss_probe_hashmap

const (
	CtrlEmpty   uint8 = 0x80
	CtrlDeleted uint8 = 0xFE
	GroupWidth  int   = 8
)

// SwissSlot holds a single key-value entry in flat table storage.
type SwissSlot struct {
	Key uint64
	Val uint64
}

// SwissMap implements an open-addressed hash table with separated 1-byte ctrl tags.
type SwissMap struct {
	Ctrl     []uint8
	Slots    []SwissSlot
	Count    int
	Tomb     int
	Capacity int
}
