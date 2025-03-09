package main

import (
	"context"
	"log"
	"os"
	"templparser/internal/views"
)

func main() {
	f, err := os.Create("hello.html")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	err = views.Function("Conor").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}
}
