package mod_gale_shapley_stable_go_d61

type StableMatcher interface {
	Match(prefsA, prefsB [][]int) []int
}

type TaggedComponent interface {
	ModuleSlug() string
}
