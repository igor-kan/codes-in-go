package pade_approximants

// ComputePadeApprox9296 evaluates rational function approximant order 9296.
func ComputePadeApprox9296(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
