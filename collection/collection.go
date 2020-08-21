package collection

// StringSliceToSet converts slice of string to a set.
func StringSliceToSet(slice []string) map[string]interface{} {
	set := make(map[string]interface{}, len(slice))
	for _, ele := range slice {
		set[ele] = struct{}{}
	}
	return set
}

// StringSliceFilter apply filter to slice of string.
func StringSliceFilter(slice []string, filter func(ele string) bool) (ret []string) {
	for _, ele := range slice {
		if filter(ele) {
			ret = append(ret, ele)
		}
	}
	return ret
}

// StringSliceForEach apply the do function to each element
func StringSliceForEach(slice []string, do func(i int, ele string, raw []string)) {
	for i, ele := range slice {
		do(i, ele, slice)
	}
}
