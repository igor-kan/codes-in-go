package mod_eertree_palindrome_go_d61

import "errors"

var (
	ErrEertreeInvalid   = errors.New("eertree_palindrome_go_d61: invalid argument")
	ErrEertreeBounds    = errors.New("eertree_palindrome_go_d61: index out of bounds")
	ErrEertreeEmpty     = errors.New("eertree_palindrome_go_d61: empty input")
	ErrEertreeInvariant = errors.New("eertree_palindrome_go_d61: invariant violated")
)
