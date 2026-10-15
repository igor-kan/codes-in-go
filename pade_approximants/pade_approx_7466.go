package pade_approximants

// ComputePadeApprox7466 evaluates rational function approximant order 7466.
func ComputePadeApprox7466(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
