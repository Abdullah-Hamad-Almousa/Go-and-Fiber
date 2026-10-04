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

func main() {

	name, err := FindUser(42)
	if err != nil {
		fmt.Println("Caught sentinel error:", err)
		return
	}
	fmt.Println("Name:", name)

}
