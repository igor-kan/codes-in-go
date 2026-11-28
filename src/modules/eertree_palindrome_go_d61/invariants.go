package mod_eertree_palindrome_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrEertreeInvariant
	}
	return nil
}
