package mod_jump_consistent_hash_go

import "errors"

var (
	ErrJumpHashInvalid   = errors.New("jump_consistent_hash_go: invalid argument")
	ErrJumpHashBounds    = errors.New("jump_consistent_hash_go: index out of bounds")
	ErrJumpHashEmpty     = errors.New("jump_consistent_hash_go: empty input")
	ErrJumpHashInvariant = errors.New("jump_consistent_hash_go: invariant violated")
)
