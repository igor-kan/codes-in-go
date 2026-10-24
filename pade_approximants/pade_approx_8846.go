package pade_approximants

// ComputePadeApprox8846 evaluates rational function approximant order 8846.
func ComputePadeApprox8846(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
