package pade_approximants

// ComputePadeApprox7276 evaluates rational function approximant order 7276.
func ComputePadeApprox7276(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
