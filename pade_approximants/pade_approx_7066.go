package pade_approximants

// ComputePadeApprox7066 evaluates rational function approximant order 7066.
func ComputePadeApprox7066(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
