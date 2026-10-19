package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8162 evaluates chebyshev collocation node order 8162.
func ComputeChebyshevColloc8162(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
