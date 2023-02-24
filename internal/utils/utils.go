package utils

func Contains[T comparable](slice []T, element T) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}

	return false
}
