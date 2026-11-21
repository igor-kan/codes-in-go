package mod_lfu_o1_cache_d56

type LfuCache struct {
	cap     int
	minFreq int
	vals    map[int]int
	freqs   map[int]int
	buckets map[int][]int
}

func NewLfuCache(cap int) *LfuCache {
	return &LfuCache{
		cap:     cap,
		vals:    make(map[int]int),
		freqs:   make(map[int]int),
		buckets: make(map[int][]int),
	}
}

func (c *LfuCache) touch(key int) {
	f := c.freqs[key]
	list := c.buckets[f]
	nextList := list[:0]
	for _, k := range list {
		if k != key {
			nextList = append(nextList, k)
		}
	}
	if len(nextList) == 0 {
		delete(c.buckets, f)
		if c.minFreq == f {
			c.minFreq++
		}
	} else {
		c.buckets[f] = nextList
	}
	c.freqs[key] = f + 1
	c.buckets[f+1] = append(c.buckets[f+1], key)
}

func (c *LfuCache) Get(key int) (int, bool) {
	v, ok := c.vals[key]
	if !ok {
		return 0, false
	}
	c.touch(key)
	return v, true
}

func (c *LfuCache) Put(key, value int) {
	if c.cap <= 0 {
		return
	}
	if _, ok := c.vals[key]; ok {
		c.vals[key] = value
		c.touch(key)
		return
	}
	if len(c.vals) >= c.cap {
		bkt := c.buckets[c.minFreq]
		evict := bkt[0]
		if len(bkt) == 1 {
			delete(c.buckets, c.minFreq)
		} else {
			c.buckets[c.minFreq] = bkt[1:]
		}
		delete(c.vals, evict)
		delete(c.freqs, evict)
	}
	c.vals[key] = value
	c.freqs[key] = 1
	c.buckets[1] = append(c.buckets[1], key)
	c.minFreq = 1
}
