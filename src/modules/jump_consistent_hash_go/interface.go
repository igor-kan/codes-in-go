package mod_jump_consistent_hash_go

type BucketMapper interface {
	Locate(key uint64, numBuckets int32) int32
}

type TaggedComponent interface {
	ModuleSlug() string
}
