package mod_lfu_cache_o1_buckets

import "container/list"

func (c *LFUCache) touch(el *list.Element) {
	node := el.Value.(*LFUNode)
	oldFreq := node.Freq
	bucket := c.Buckets[oldFreq]
	bucket.Remove(el)
	if bucket.Len() == 0 {
		delete(c.Buckets, oldFreq)
		if c.MinFreq == oldFreq {
			c.MinFreq++
		}
	}
	node.Freq++
	nextB := c.Buckets[node.Freq]
	if nextB == nil {
		nextB = list.New()
		c.Buckets[node.Freq] = nextB
	}
	newEl := nextB.PushFront(node)
	c.Nodes[node.Key] = newEl
}

// Get retrieves cached value and promotes frequency in O(1).
func (c *LFUCache) Get(key int) (int, bool) {
	el, ok := c.Nodes[key]
	if !ok {
		return 0, false
	}
	c.touch(el)
	return el.Value.(*LFUNode).Val, true
}

// Put inserts or updates key -> val, evicting LRU item in MinFreq bucket if full.
func (c *LFUCache) Put(key, val int) {
	if el, ok := c.Nodes[key]; ok {
		el.Value.(*LFUNode).Val = val
		c.touch(el)
		return
	}
	if len(c.Nodes) == c.Cap {
		minB := c.Buckets[c.MinFreq]
		victim := minB.Back()
		vNode := victim.Value.(*LFUNode)
		minB.Remove(victim)
		if minB.Len() == 0 {
			delete(c.Buckets, c.MinFreq)
		}
		delete(c.Nodes, vNode.Key)
	}
	c.MinFreq = 1
	b1 := c.Buckets[1]
	if b1 == nil {
		b1 = list.New()
		c.Buckets[1] = b1
	}
	node := &LFUNode{Key: key, Val: val, Freq: 1}
	c.Nodes[key] = b1.PushFront(node)
}
