package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateSecretKey generates a random secret key for the user
func GenerateSecretKey() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random secret key
	secretKey := make([]byte, 16)
	for i := range secretKey {
		secretKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(secretKey)
}

// GenerateOTP generates a 6-digit random OTP
func GenerateOTP() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit OTP
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// VerifyOTP verifies the entered OTP against the generated one
func VerifyOTP(userOTP, generatedOTP string) bool {
	return userOTP == generatedOTP
}

func main() {
	// Simulate user registration and saving secret key
	secretKey := GenerateSecretKey()
	fmt.Println("Secret Key:", secretKey)

	// Simulate user authentication process
	generatedOTP := GenerateOTP()
	fmt.Println("Generated OTP:", generatedOTP)

	// Simulate user input
	var userOTP string
	fmt.Print("Enter OTP: ")
	fmt.Scanln(&userOTP)

	// Verify OTP
	if VerifyOTP(userOTP, generatedOTP) {
		fmt.Println("OTP is valid! Access granted.")
	} else {
		fmt.Println("OTP is invalid! Access denied.")
	}
}
