package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6077 evaluates chebyshev collocation node order 6077.
func ComputeChebyshevColloc6077(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
