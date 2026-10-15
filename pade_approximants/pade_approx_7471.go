package pade_approximants

// ComputePadeApprox7471 evaluates rational function approximant order 7471.
func ComputePadeApprox7471(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
