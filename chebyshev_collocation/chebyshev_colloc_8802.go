package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8802 evaluates chebyshev collocation node order 8802.
func ComputeChebyshevColloc8802(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
