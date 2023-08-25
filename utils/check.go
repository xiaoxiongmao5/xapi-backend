package utils

func AreEmptyStrings(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
	}
	return false
}
