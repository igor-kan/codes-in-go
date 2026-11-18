package mod_weighted_round_robin_gcd

type WeightedBalancer interface {
	NextPeer() int
}

type TaggedComponent interface {
	ModuleSlug() string
}
