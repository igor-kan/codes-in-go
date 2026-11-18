package mod_weighted_round_robin_gcd

type PeerWeight struct {
	ID      int
	Weight  int
	Current int
}

type ModuleMeta struct {
	Slug    string
	Version int
}
