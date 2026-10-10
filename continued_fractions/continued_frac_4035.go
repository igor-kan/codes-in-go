package continued_fractions

// ComputeContinuedFrac4035 evaluates continued fraction approximant order 4035.
func ComputeContinuedFrac4035(x float64) float64 {
	a := 1.0
	for k := 4; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
