package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9207 evaluates chebyshev collocation node order 9207.
func ComputeChebyshevColloc9207(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
