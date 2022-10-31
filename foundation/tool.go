package foundation

func InIntSlice(s int, a []int) bool {
	result := false
	for _, v := range a {
		if s == v {
			result = true
		}
	}
	return result
}

func InStringSlice(s string, a []string) bool {
	result := false
	for _, v := range a {
		if s == v {
			result = true
		}
	}
	return result
}
