package pade_approximants

// ComputePadeApprox9266 evaluates rational function approximant order 9266.
func ComputePadeApprox9266(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
