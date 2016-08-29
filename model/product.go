package model

type Product struct {
	Name     string
	Category string
	Prize    float64
	Rating   float64
	InStock  bool
	ID       string
}

type ErrorMessage struct {
	Message   string
	ErrorCode string
}

var InvalidArgsError ErrorMessage = ErrorMessage{Message: "Invalid Arguement", ErrorCode: "10000"}
var InternalServerError ErrorMessage = ErrorMessage{Message: "Internal Server Error", ErrorCode: "10001"}
