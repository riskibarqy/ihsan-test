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

// CleanPhoneNumber removes any non-numeric characters from the phone number
func CleanPhoneNumber(phone string) string {
	re := regexp.MustCompile(`\D`) // \D matches any non-digit character
	return re.ReplaceAllString(phone, "")
}
