package pade_approximants

// ComputePadeApprox8326 evaluates rational function approximant order 8326.
func ComputePadeApprox8326(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
