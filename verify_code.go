package gauth

// VerifyCode verify code
func VerifyCode(secret string, code string) bool {
	ga := NewGAuth()
	ret, err := ga.VerifyCode(secret, code)
	if err != nil {
		return false
	}

	return ret
}
