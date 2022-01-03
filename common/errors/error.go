package errors

import (
	"fmt"
)

var (
	// ErrInvalidArgument ...
	ErrInvalidArgument = fmt.Errorf("information is invalid")

	// ErrPasswordIsNotCorrect ...
	ErrPasswordIsNotCorrect = fmt.Errorf("password is incorrect")

	// ErrInvalidOTP ...
	ErrInvalidOTP = fmt.Errorf("otp is invalid")

	// ErrInvalidToken ...
	ErrInvalidToken = fmt.Errorf("token is invalid")

	// ErrUsernameAlreadyExists ...
	ErrUsernameAlreadyExists = fmt.Errorf("username has been already registered")

	// ErrUsernameNotFound ...
	ErrUsernameNotFound = fmt.Errorf("username not found")

	// ErrPhoneNotFound ...
	ErrPhoneNotFound = fmt.Errorf("phone number not found")

	// ErrEmailNotFound ...
	ErrEmailNotFound = fmt.Errorf("email not found")

	// ErrPhoneAlreadyExists ...
	ErrPhoneAlreadyExists = fmt.Errorf("the phone number has been already registered")

	// ErrEmailAlreadyExists ...
	ErrEmailAlreadyExists = fmt.Errorf("the email has been already registered")
)
