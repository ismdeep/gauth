package gauth

// CreateSecret create secret
func CreateSecret() string {
	ga := NewGAuth()
	secret, err := ga.CreateSecret(16)
	if err != nil {
		return ""
	}
	return secret
}
