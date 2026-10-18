package pade_approximants

// ComputePadeApprox7911 evaluates rational function approximant order 7911.
func ComputePadeApprox7911(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
