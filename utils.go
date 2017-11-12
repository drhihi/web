package main

import (
	"crypto/rand"
	"fmt"

	"github.com/russross/blackfriday"
)

func generateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func convertMarkdownToHtml(markdown string) string {
	return string(blackfriday.MarkdownBasic([]byte(markdown)))
}