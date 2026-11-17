package mod_count_min_heavy_go

func verifyNoUnderestimate(cms *ConservativeCMS, exact map[string]uint64) error {
	for k, v := range exact {
		if cms.Estimate(k) < v {
			return ErrCountMinInvariant
		}
	}
	return nil
}
