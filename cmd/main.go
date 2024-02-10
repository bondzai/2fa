package main

import (
	"fmt"
	"math/rand"
	"time"
)

// UserRepository represents the user repository interface
type UserRepository interface {
	GetSecretKey(userID string) string
}

// UserInteractor represents the user interactor
type UserInteractor struct {
	UserRepo UserRepository
}

// GenerateSecretKey generates a random secret key for the user
func (ui *UserInteractor) GenerateSecretKey(userID string) string {
	return ui.UserRepo.GetSecretKey(userID)
}

// OTPGenerator represents the OTP generator interface
type OTPGenerator interface {
	GenerateOTP() string
}

// OTPInteractor represents the OTP interactor
type OTPInteractor struct {
	OTPGen OTPGenerator
}

// GenerateOTP generates a 6-digit random OTP
func (oi *OTPInteractor) GenerateOTP() string {
	return oi.OTPGen.GenerateOTP()
}

// Authenticator represents the authenticator interface
type Authenticator interface {
	VerifyOTP(userOTP, generatedOTP string) bool
}

// AuthenticatorInteractor represents the authenticator interactor
type AuthenticatorInteractor struct{}

// VerifyOTP verifies the entered OTP against the generated one
func (ai *AuthenticatorInteractor) VerifyOTP(userOTP, generatedOTP string) bool {
	return userOTP == generatedOTP
}

// ConsoleInputOutput represents the console input/output adapter
type ConsoleInputOutput struct{}

// GetUserInput gets user input from the console
func (c *ConsoleInputOutput) GetUserInput(prompt string) string {
	var userInput string
	fmt.Print(prompt)
	fmt.Scanln(&userInput)
	return userInput
}

// ConsoleOutput prints output to the console
type ConsoleOutput struct{}

// PrintMessage prints a message to the console
func (c *ConsoleOutput) PrintMessage(message string) {
	fmt.Println(message)
}

// UserRepositoryImpl represents the user repository implementation
type UserRepositoryImpl struct{}

// GetSecretKey retrieves the secret key for the user from the database
func (ur *UserRepositoryImpl) GetSecretKey(userID string) string {
	// Simulate fetching secret key from the database
	return "USER_SECRET_KEY"
}

// OTPGeneratorImpl represents the OTP generator implementation
type OTPGeneratorImpl struct{}

// GenerateOTP generates a 6-digit random OTP
func (og *OTPGeneratorImpl) GenerateOTP() string {
	const charset = "0123456789"
	rand.Seed(time.Now().UnixNano())
	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = charset[rand.Intn(len(charset))]
	}
	return string(otp)
}

func main() {
	userRepo := &UserRepositoryImpl{}
	userInteractor := &UserInteractor{UserRepo: userRepo}

	otpGen := &OTPGeneratorImpl{}
	otpInteractor := &OTPInteractor{OTPGen: otpGen}

	authenticatorInteractor := &AuthenticatorInteractor{}

	inputOutput := &ConsoleInputOutput{}
	output := &ConsoleOutput{}

	// Simulate user registration and saving secret key
	secretKey := userInteractor.GenerateSecretKey("USER_ID")
	output.PrintMessage("Secret Key: " + secretKey)

	// Simulate user authentication process
	generatedOTP := otpInteractor.GenerateOTP()
	output.PrintMessage("Generated OTP: " + generatedOTP)

	// Simulate user input
	userOTP := inputOutput.GetUserInput("Enter OTP: ")

	// Verify OTP
	if authenticatorInteractor.VerifyOTP(userOTP, generatedOTP) {
		output.PrintMessage("OTP is valid! Access granted.")
	} else {
		output.PrintMessage("OTP is invalid! Access denied.")
	}
}
