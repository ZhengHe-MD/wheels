package collection

// StringSliceToSet converts slice of string to a set.
func StringSliceToSet(slice []string) map[string]interface{} {
	set := make(map[string]interface{}, len(slice))
	for _, ele := range slice {
		set[ele] = struct{}{}
	}
	return set
}

// FilterStringSlice apply filter to slice of string.
func FilterStringSlice(slice []string, filter func(ele string) bool) (ret []string) {
	for _, ele := range slice {
		if filter(ele) {
			ret = append(ret, ele)
		}
	}
	return ret
}
