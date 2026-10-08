package pade_approximants

// ComputePadeApprox1106 evaluates rational function approximant order 1106.
func ComputePadeApprox1106(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
