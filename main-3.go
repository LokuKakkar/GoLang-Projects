package main

import (
	"fmt"
	"log"
	"os"
)

func main3() {
	fmt.Println("Hello World 3")
	log.Println("This is a log message")
	// os.Exit(1)
	f, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer f.Close()

	s := []byte("This is a log message")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error writing to file: %v", err)
	}

	log.SetOutput(f)
}
