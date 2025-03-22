package main

import "fmt"

type AppError struct {
	Code string `json:"code"`
}

func (e *AppError) Error() string {
	fmt.Println(e.Code)
	return e.Code
}

func NewAppError(code string) *AppError {
	return &AppError{Code: code}
}
