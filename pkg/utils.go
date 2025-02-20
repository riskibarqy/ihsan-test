package utils

import (
	"fmt"
	"regexp"
	"time"

	"math/rand"
)

// Now generate unix number of current time
func Now() int {
	return int(time.Now().Unix())
}

// GenerateNoRekening generates a 12-digit random bank account number
func GenerateNoRekening() string {
	prefix := "666"
	randomNumber := rand.Intn(900000000) + 100000000
	return fmt.Sprintf("%s%d", prefix, randomNumber)
}

// CleanPhoneNumber removes any non numeric characters from the phone number
func CleanPhoneNumber(phone string) string {
	re := regexp.MustCompile(`\D`) // \D matches any non-digit character
	return re.ReplaceAllString(phone, "")
}

// GenerateTransactionID creates a unique transaction ID
func GenerateTransactionID() string {
	randomCode := randomString(4) + randomDigits(3)
	return fmt.Sprintf("TXN-%d-%s", Now(), randomCode)
}

// randomString generates a random uppercase string
func randomString(n int) string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// randomDigits generates a random numeric string
func randomDigits(n int) string {
	digits := "0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}
