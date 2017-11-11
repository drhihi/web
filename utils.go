package main

import (
	"crypto/rand"
	"fmt"
)

func generateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprint("%x", b)
}