package pade_approximants

// ComputePadeApprox9271 evaluates rational function approximant order 9271.
func ComputePadeApprox9271(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
