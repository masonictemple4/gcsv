package utils

func Contains[K comparable](elems []K, v K) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
