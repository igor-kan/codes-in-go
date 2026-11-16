package mod_reservoir_quantiles_gk

// QuantileSketch supports single-pass observation and approximate rank lookup.
type QuantileSketch interface {
	Observe(v float64)
	Quantile(phi float64) (float64, error)
}
