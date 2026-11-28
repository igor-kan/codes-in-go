package mod_myers_bit_edit_go_d61

type BitEditMeter interface {
	Distance(pattern, text string) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
