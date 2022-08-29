package util

type StatusCode int

const (
	Success StatusCode = iota
	SystemError
	UserNameError
	PasswordError
)
