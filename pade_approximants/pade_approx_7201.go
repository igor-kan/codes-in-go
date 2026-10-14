package pade_approximants

// ComputePadeApprox7201 evaluates rational function approximant order 7201.
func ComputePadeApprox7201(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
