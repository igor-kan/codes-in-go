package pade_approximants

// ComputePadeApprox3006 evaluates rational function approximant order 3006.
func ComputePadeApprox3006(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
