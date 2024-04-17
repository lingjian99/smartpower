package generics

// Difference 差集 [1,2,3,4], [1,2,5,6] s1\s2 ==>> [3,4]
func Difference[K comparable](s1 []K, s2 []K) []K {
	m := make(map[K]bool)

	for _, item := range s2 {
		m[item] = true
	}

	var diff []K
	for _, item := range s1 {
		if !m[item] {
			diff = append(diff, item)
		}
	}

	return diff
}
