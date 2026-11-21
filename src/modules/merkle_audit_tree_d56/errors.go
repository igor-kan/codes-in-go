package mod_merkle_audit_tree_d56

import "errors"

var (
	ErrMerkleInvalid   = errors.New("merkle_audit_tree_d56: invalid argument")
	ErrMerkleBounds    = errors.New("merkle_audit_tree_d56: index out of bounds")
	ErrMerkleEmpty     = errors.New("merkle_audit_tree_d56: empty input")
	ErrMerkleInvariant = errors.New("merkle_audit_tree_d56: invariant violated")
)
