package util

func Contain(collection []string, search string) bool {
	result := false
	for _, s := range collection {
		if s == search {
			result = true
			break
		}
	}
	return result
}
