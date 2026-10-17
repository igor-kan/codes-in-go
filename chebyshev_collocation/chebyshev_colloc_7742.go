package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7742 evaluates chebyshev collocation node order 7742.
func ComputeChebyshevColloc7742(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
