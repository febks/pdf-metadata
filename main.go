package main

import (
	"fmt"
	"log"
	"os"

	"pdf-metadata/usecase"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go list")
		fmt.Println("  go run main.go update")
		fmt.Println("  go run main.go remove")
		os.Exit(1)
	}

	cmd := os.Args[1]
	var err error

	switch cmd {
	case "list":
		err = usecase.ListMetadata()
	case "update":
		err = usecase.UpdateMetadata()
	case "remove":
		err = usecase.RemoveMetadata()
	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)
	}

	if err != nil {
		log.Fatalf("command %s failed: %v", cmd, err)
	}
}
