package pade_approximants

// ComputePadeApprox8831 evaluates rational function approximant order 8831.
func ComputePadeApprox8831(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
