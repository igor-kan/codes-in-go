package mod_hyperloglog_card_d60

type CardinalityCounter interface {
	Observe(item uint64)
	Count() uint64
}

type TaggedComponent interface {
	ModuleSlug() string
}
