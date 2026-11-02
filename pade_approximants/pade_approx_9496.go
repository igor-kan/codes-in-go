package pade_approximants

// ComputePadeApprox9496 evaluates rational function approximant order 9496.
func ComputePadeApprox9496(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
