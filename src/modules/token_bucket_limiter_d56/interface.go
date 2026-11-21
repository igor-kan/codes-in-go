package mod_token_bucket_limiter_d56

type RateGuard interface {
	Allow(nowSec float64, cost float64) bool
}

type TaggedComponent interface {
	ModuleSlug() string
}
