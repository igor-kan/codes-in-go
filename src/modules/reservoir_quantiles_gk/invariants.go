package mod_reservoir_quantiles_gk

import "sort"

func VerifyRankAccuracy(raw []float64, g *GKSummary, phi float64) error {
	est, err := g.Quantile(phi)
	if err != nil {
		return err
	}
	sorted := append([]float64(nil), raw...)
	sort.Float64s(sorted)
	lo := 0
	for lo < len(sorted) && sorted[lo] < est {
		lo++
	}
	desired := phi * float64(len(sorted))
	maxErr := (g.eps * float64(len(sorted))) + 2.0
	if float64(lo) < desired-maxErr || float64(lo) > desired+maxErr {
		return ErrQuantileInvariant
	}
	return nil
}
