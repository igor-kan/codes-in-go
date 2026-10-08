package pade_approximants

// ComputePadeApprox2006 evaluates rational function approximant order 2006.
func ComputePadeApprox2006(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
