package continued_fractions

// ComputeContinuedFrac4085 evaluates continued fraction approximant order 4085.
func ComputeContinuedFrac4085(x float64) float64 {
	a := 1.0
	for k := 6; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
