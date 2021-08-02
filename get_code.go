package gauth

// GetCode get code
func GetCode(secret string) string {
	ga := NewGAuth()
	code, err := ga.GetCode(secret)
	if err != nil {
		return ""
	}

	return code
}
