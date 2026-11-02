package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9522 evaluates chebyshev collocation node order 9522.
func ComputeChebyshevColloc9522(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
