package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7607 evaluates chebyshev collocation node order 7607.
func ComputeChebyshevColloc7607(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
