package pade_approximants

// ComputePadeApprox6046 evaluates rational function approximant order 6046.
func ComputePadeApprox6046(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
