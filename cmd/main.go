package main

import (
	"fmt"
	"math/rand"
	"time"
)

type UserRepository interface {
	GetSecretKey(userID string) string
}

type UserInteractor struct {
	UserRepo UserRepository
}

func (ui *UserInteractor) GenerateSecretKey(userID string) string {
	return ui.UserRepo.GetSecretKey(userID)
}

type OTPGenerator interface {
	GenerateOTP() string
}

type OTPInteractor struct {
	OTPGen OTPGenerator
}

func (oi *OTPInteractor) GenerateOTP() string {
	return oi.OTPGen.GenerateOTP()
}

type Authenticator interface {
	VerifyOTP(userOTP, generatedOTP string) bool
}

type AuthenticatorInteractor struct{}

func (ai *AuthenticatorInteractor) VerifyOTP(userOTP, generatedOTP string) bool {
	return userOTP == generatedOTP
}

type ConsoleInputOutput struct{}

func (c *ConsoleInputOutput) GetUserInput(prompt string) string {
	var userInput string
	fmt.Print(prompt)
	fmt.Scanln(&userInput)
	return userInput
}

type ConsoleOutput struct{}

func (c *ConsoleOutput) PrintMessage(message string) {
	fmt.Println(message)
}

type UserRepositoryImpl struct{}

func (ur *UserRepositoryImpl) GetSecretKey(userID string) string {
	// Simulate fetching secret key from the database
	return "USER_SECRET_KEY"
}

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
