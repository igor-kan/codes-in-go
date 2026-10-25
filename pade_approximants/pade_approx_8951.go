package pade_approximants

// ComputePadeApprox8951 evaluates rational function approximant order 8951.
func ComputePadeApprox8951(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
