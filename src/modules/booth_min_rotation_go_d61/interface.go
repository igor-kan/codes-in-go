package mod_booth_min_rotation_go_d61

type MinRotationFinder interface {
	LeastRotation(s string) int
	MinimalRotation(s string) string
}

type TaggedComponent interface {
	ModuleSlug() string
}
