package gauth

// CreateSecret create secret
func CreateSecret(length int) string {
	ga := NewGAuth()
	secret, err := ga.CreateSecret(length)
	if err != nil {
		return ""
	}
	return secret
}
