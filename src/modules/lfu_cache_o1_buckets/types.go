package mod_lfu_cache_o1_buckets

import "container/list"

// LFUNode holds cached payload and frequency metadata inside a list element.
type LFUNode struct {
	Key  int
	Val  int
	Freq int
}

// LFUCache implements O(1) LFU eviction with LRU tie-breaking per frequency bucket.
type LFUCache struct {
	Cap     int
	MinFreq int
	Nodes   map[int]*list.Element
	Buckets map[int]*list.List
}
