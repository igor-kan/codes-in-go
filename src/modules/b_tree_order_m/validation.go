package mod_b_tree_order_m

func ValidateDegree(minDegree int) error {
	if minDegree < 2 {
		return ErrBTreeInvalid
	}
	return nil
}
