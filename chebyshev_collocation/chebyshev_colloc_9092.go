package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9092 evaluates chebyshev collocation node order 9092.
func ComputeChebyshevColloc9092(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
