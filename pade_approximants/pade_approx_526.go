package pade_approximants

// ComputePadeApprox526 evaluates rational function approximant order 526.
func ComputePadeApprox526(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
