package pade_approximants

// ComputePadeApprox9261 evaluates rational function approximant order 9261.
func ComputePadeApprox9261(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
