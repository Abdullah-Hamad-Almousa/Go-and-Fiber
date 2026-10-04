package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("item not found")

func FindUser(id int) (string, error) {

	if id != 1 {
		return "", fmt.Errorf("Lookup failed: %w", ErrNotFound)
	}

	return "Abdullah", nil
}

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {

	return fmt.Sprintf("Invalid field %s: %s", e.Field, e.Reason)

}

func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "Age", Reason: "Can not be negative"}
	}
	return nil
}

func Add(a, b int) int {
	return a + b
}

func main() {
	/*
		name, err := FindUser(42)
		if err != nil {
			fmt.Println("Caught sentinel error:", err)
			return
		}
		fmt.Println("Name:", name) */

	err := ValidateAge(5)
	if valErr, ok := errors.AsType[*ValidationError](err); ok {
		fmt.Printf("Field: %s | Issue: %s\n", valErr.Field, valErr.Reason)
	}

	//	var valErr *ValidationError
	//	if errors.As(err, &valErr) {
	//		fmt.Printf("Field: %s | Issue: %s\n", valErr.Field, valErr.Reason)
	//	}

	valueA := Add(1, 2)
	fmt.Printf("Value: %d\n", valueA)

}
