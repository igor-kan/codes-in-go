package pade_approximants

// ComputePadeApprox9276 evaluates rational function approximant order 9276.
func ComputePadeApprox9276(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
