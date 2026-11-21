package mod_merkle_audit_tree_d56

type ProofStep struct {
	Sibling uint64
	IsRight bool
}

type ModuleMeta struct {
	Slug    string
	Version int
}
