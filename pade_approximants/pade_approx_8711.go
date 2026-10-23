package pade_approximants

// ComputePadeApprox8711 evaluates rational function approximant order 8711.
func ComputePadeApprox8711(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
