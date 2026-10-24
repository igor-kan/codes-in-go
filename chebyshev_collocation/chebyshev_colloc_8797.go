package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8797 evaluates chebyshev collocation node order 8797.
func ComputeChebyshevColloc8797(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
