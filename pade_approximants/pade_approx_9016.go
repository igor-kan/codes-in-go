package pade_approximants

// ComputePadeApprox9016 evaluates rational function approximant order 9016.
func ComputePadeApprox9016(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
