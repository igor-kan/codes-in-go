package pade_approximants

// ComputePadeApprox7756 evaluates rational function approximant order 7756.
func ComputePadeApprox7756(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
