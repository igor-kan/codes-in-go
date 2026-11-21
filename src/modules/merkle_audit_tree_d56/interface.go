package mod_merkle_audit_tree_d56

type Auditor interface {
	Root() uint64
}

type TaggedComponent interface {
	ModuleSlug() string
}
