package gauth

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// GAuth struct of gauth
type GAuth struct {
	codeLen float64
	table   map[string]int
}

// NewGAuth new gauth
func NewGAuth() *GAuth {
	return &GAuth{
		codeLen: 6,
		table:   arrayFlip(table),
	}
}

// SetCodeLength Set the code length, should be >=6
func (receiver *GAuth) SetCodeLength(length float64) error {
	if length < 6 {
		return ErrSecretLengthLss
	}
	receiver.codeLen = length
	return nil
}

// CreateSecret create new secret
// 16 characters, randomly chosen from the allowed base32 characters.
func (receiver *GAuth) CreateSecret(lens ...int) (string, error) {
	var (
		length int
		secret []string
	)
	// init length
	switch len(lens) {
	case 0:
		length = 16
	case 1:
		length = lens[0]
	default:
		return "", ErrParam
	}
	for i := 0; i < length; i++ {
		val := new(big.Int).SetInt64(int64(128))
		val, _ = rand.Int(rand.Reader, val)
		n := val.Int64()
		secret = append(secret, table[int(n)%len(table)])
	}
	return strings.Join(secret, ""), nil
}

// VerifyCode Check if the code is correct. This will accept codes starting from $discrepancy*30sec ago to $discrepancy*30sec from now
func (receiver *GAuth) VerifyCode(secret, code string, discrepancy int64) (bool, error) {
	// now time
	curTimeSlice := time.Now().Unix() / 30
	for i := -discrepancy; i <= discrepancy; i++ {
		calculatedCode, err := receiver.GetCode(secret, curTimeSlice+i)
		if err != nil {
			return false, err
		}
		if calculatedCode == code {
			return true, nil
		}
	}
	return false, nil
}

// GetCode Calculate the code, with given secret and point in time
func (receiver *GAuth) GetCode(secret string, timeSlices ...int64) (string, error) {
	var timeSlice int64
	switch len(timeSlices) {
	case 0:
		timeSlice = time.Now().Unix() / 30
	case 1:
		timeSlice = timeSlices[0]
	default:
		return "", ErrParam
	}
	secret = strings.ToUpper(secret)
	secretKey, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}
	tim, err := hex.DecodeString(fmt.Sprintf("%016x", timeSlice))
	if err != nil {
		return "", err
	}
	hm := HmacSha1(secretKey, tim)
	offset := hm[len(hm)-1] & 0x0F
	hashPart := hm[offset : offset+4]
	value, err := strconv.ParseInt(hex.EncodeToString(hashPart), 16, 0)
	if err != nil {
		return "", err
	}
	value = value & 0x7FFFFFFF
	modulo := int64(math.Pow(10, receiver.codeLen))
	format := fmt.Sprintf("%%0%dd", int(receiver.codeLen))
	return fmt.Sprintf(format, value%modulo), nil
}
