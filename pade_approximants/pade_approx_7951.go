package pade_approximants

// ComputePadeApprox7951 evaluates rational function approximant order 7951.
func ComputePadeApprox7951(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
