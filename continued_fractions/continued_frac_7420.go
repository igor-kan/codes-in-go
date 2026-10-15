package continued_fractions

// ComputeContinuedFrac7420 evaluates continued fraction approximant order 7420.
func ComputeContinuedFrac7420(x float64) float64 {
	a := 1.0
	for k := 5; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
