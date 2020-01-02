package util

// StrSlice converts an interface slice to a slice of strings.
func StrSlice(i []interface{}) []string {
	var s []string
	for _, v := range i {
		s = append(s, v.(string))
	}
	return s
}

// BoolPtr returns a pointer to the input bool.
func BoolPtr(b bool) *bool {
	return &b
}

// StrPtr returns a pointer to the input string.
func StrPtr(s string) *string {
	return &s
}
