package pade_approximants

// ComputePadeApprox7616 evaluates rational function approximant order 7616.
func ComputePadeApprox7616(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
