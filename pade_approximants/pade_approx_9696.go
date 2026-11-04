package pade_approximants

// ComputePadeApprox9696 evaluates rational function approximant order 9696.
func ComputePadeApprox9696(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
