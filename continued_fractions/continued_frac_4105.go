package continued_fractions

// ComputeContinuedFrac4105 evaluates continued fraction approximant order 4105.
func ComputeContinuedFrac4105(x float64) float64 {
	a := 1.0
	for k := 2; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
