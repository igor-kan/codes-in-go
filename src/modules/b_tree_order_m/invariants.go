package mod_b_tree_order_m

func checkNode(n *bnode, deg int, isRoot bool) error {
	if !isRoot && len(n.items) < deg-1 {
		return ErrBTreeInvariant
	}
	if len(n.items) > 2*deg-1 {
		return ErrBTreeInvariant
	}
	for i := 1; i < len(n.items); i++ {
		if n.items[i-1].Key >= n.items[i].Key {
			return ErrBTreeInvariant
		}
	}
	if !n.leaf {
		if len(n.children) != len(n.items)+1 {
			return ErrBTreeInvariant
		}
		for _, ch := range n.children {
			if err := checkNode(ch, deg, false); err != nil {
				return err
			}
		}
	}
	return nil
}

func VerifyBTreeInvariants(t *BTree) error {
	if err := checkNode(t.root, t.deg, true); err != nil {
		return err
	}
	all := t.InOrder()
	if len(all) != t.size {
		return ErrBTreeInvariant
	}
	for i := 1; i < len(all); i++ {
		if all[i-1].Key >= all[i].Key {
			return ErrBTreeInvariant
		}
	}
	return nil
}
