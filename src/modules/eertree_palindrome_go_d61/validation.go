package mod_eertree_palindrome_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrEertreeInvalid
	}
	return nil
}
