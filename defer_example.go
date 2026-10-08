package main

import (
	"fmt"
	"os"
)

// Goal: Guarantee cleanup actions (like closing/removing a file)
// run reliably before a function exits.
func RunDeferExample() {
	fileName := "temp.demo.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer func() {
		file.Close()
		os.Remove(fileName)
		fmt.Println("[defer] Cleaned up and removed temp file.")
	}()

	//This will write a new file
	file.WriteString("Hello from defer!")
	fmt.Println("[defer] Finished file operations.")

}
