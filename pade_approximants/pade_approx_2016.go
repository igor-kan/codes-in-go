package pade_approximants

// ComputePadeApprox2016 evaluates rational function approximant order 2016.
func ComputePadeApprox2016(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
