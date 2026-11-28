package mod_booth_min_rotation_go_d61

func LeastRotation(s string) int {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return 0
	}
	doubled := make([]rune, 2*n)
	copy(doubled, runes)
	copy(doubled[n:], runes)
	i, j, k := 0, 1, 0
	for i < n && j < n && k < n {
		a := doubled[i+k]
		b := doubled[j+k]
		if a == b {
			k++
			continue
		}
		if a > b {
			i += k + 1
			if i <= j {
				i = j + 1
			}
		} else {
			j += k + 1
			if j <= i {
				j = i + 1
			}
		}
		k = 0
	}
	if i < j {
		return i
	}
	return j
}

func MinimalRotation(s string) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}
	start := LeastRotation(s)
	out := make([]rune, 0, n)
	out = append(out, runes[start:]...)
	out = append(out, runes[:start]...)
	return string(out)
}

func BoothInfo(s string) RotationResult {
	return RotationResult{Start: LeastRotation(s), Text: MinimalRotation(s)}
}

type BoothFinder struct{}

func (b BoothFinder) LeastRotation(s string) int {
	return LeastRotation(s)
}

func (b BoothFinder) MinimalRotation(s string) string {
	return MinimalRotation(s)
}

var _ MinRotationFinder = BoothFinder{}
