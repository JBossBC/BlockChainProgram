package util

type StatusCode int

const (
	success StatusCode = iota
	failed
	systemError
)
