package azuredevops

func firstNonEmpty(values ...*string) *string {
	for _, val := range values {
		if val != nil {
			if len(*val) > 0 {
				return val
			}
		}
	}
	return nil
}
