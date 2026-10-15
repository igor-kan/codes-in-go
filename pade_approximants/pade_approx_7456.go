package pade_approximants

// ComputePadeApprox7456 evaluates rational function approximant order 7456.
func ComputePadeApprox7456(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
