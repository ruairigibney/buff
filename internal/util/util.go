package util

// Int64Contains checks an int64 slice to see if it contains a specific int64
func Int64Contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
