package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	example3()
}

func example1() {
	h := sha1.New()
	io.WriteString(h, "His money is twice tained:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
}

func example2() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New() 
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("% x", h.Sum(nil))
}

func example3() {
	data := []byte("This page is left intentionally blank.")	
	fmt.Printf("% x", sha1.Sum(data))
}