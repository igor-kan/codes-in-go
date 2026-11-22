package mod_sliding_log_limiter_d57

type SlidingAdmitter interface {
	Allow(now int64) bool
}

type TaggedComponent interface {
	ModuleSlug() string
}
