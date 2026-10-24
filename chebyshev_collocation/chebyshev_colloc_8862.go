package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8862 evaluates chebyshev collocation node order 8862.
func ComputeChebyshevColloc8862(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
