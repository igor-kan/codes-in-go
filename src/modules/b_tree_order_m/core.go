package mod_b_tree_order_m

type BTree struct {
	deg  int
	root *bnode
	size int
}

func NewBTree(minDegree int) (*BTree, error) {
	if minDegree < 2 {
		return nil, ErrBTreeInvalid
	}
	return &BTree{
		deg:  minDegree,
		root: &bnode{leaf: true},
	}, nil
}

func (t *BTree) Len() int { return t.size }

func searchNode(n *bnode, key int64) (int64, bool) {
	i := 0
	for i < len(n.items) && key > n.items[i].Key {
		i++
	}
	if i < len(n.items) && n.items[i].Key == key {
		return n.items[i].Val, true
	}
	if n.leaf {
		return 0, false
	}
	return searchNode(n.children[i], key)
}

func (t *BTree) Get(key int64) (int64, bool) {
	return searchNode(t.root, key)
}

func (t *BTree) splitChild(parent *bnode, idx int) {
	d := t.deg
	y := parent.children[idx]
	z := &bnode{leaf: y.leaf}
	mid := y.items[d-1]
	z.items = append([]KVItem(nil), y.items[d:]...)
	y.items = y.items[:d-1]
	if !y.leaf {
		z.children = append([]*bnode(nil), y.children[d:]...)
		y.children = y.children[:d]
	}
	parent.children = append(parent.children, nil)
	copy(parent.children[idx+2:], parent.children[idx+1:])
	parent.children[idx+1] = z

	parent.items = append(parent.items, KVItem{})
	copy(parent.items[idx+1:], parent.items[idx:])
	parent.items[idx] = mid
}

func (t *BTree) insertNonFull(n *bnode, item KVItem) bool {
	i := len(n.items) - 1
	for j := 0; j < len(n.items); j++ {
		if n.items[j].Key == item.Key {
			n.items[j].Val = item.Val
			return false
		}
	}
	if n.leaf {
		n.items = append(n.items, KVItem{})
		for i >= 0 && n.items[i].Key > item.Key {
			n.items[i+1] = n.items[i]
			i--
		}
		n.items[i+1] = item
		return true
	}
	for i >= 0 && n.items[i].Key > item.Key {
		i--
	}
	i++
	if len(n.children[i].items) == 2*t.deg-1 {
		t.splitChild(n, i)
		if n.items[i].Key == item.Key {
			n.items[i].Val = item.Val
			return false
		}
		if item.Key > n.items[i].Key {
			i++
		}
	}
	return t.insertNonFull(n.children[i], item)
}

func (t *BTree) Put(key, val int64) {
	item := KVItem{Key: key, Val: val}
	maxKeys := 2*t.deg - 1
	if len(t.root.items) == maxKeys {
		s := &bnode{leaf: false, children: []*bnode{t.root}}
		t.root = s
		t.splitChild(s, 0)
	}
	if t.insertNonFull(t.root, item) {
		t.size++
	}
}

func walkNode(n *bnode, out *[]KVItem) {
	for i := 0; i < len(n.items); i++ {
		if !n.leaf {
			walkNode(n.children[i], out)
		}
		*out = append(*out, n.items[i])
	}
	if !n.leaf && len(n.children) > len(n.items) {
		walkNode(n.children[len(n.items)], out)
	}
}

func (t *BTree) InOrder() []KVItem {
	out := make([]KVItem, 0, t.size)
	walkNode(t.root, &out)
	return out
}
