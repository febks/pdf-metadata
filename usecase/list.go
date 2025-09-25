package usecase

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func ListMetadata(inputFile string) error {
	if inputFile == "" {
		inputFile = os.Getenv("INPUT_FILE")
	}

	inFile := inputFile
	if inFile == "" {
		return fmt.Errorf("INPUT_FILE env is required")
	}

	ctx, err := api.ReadContextFile(inFile)
	if err != nil {
		return fmt.Errorf("error reading metadata: %v", err)
	}
	if ctx.XRefTable.Info == nil {
		fmt.Println("No metadata found")
		return nil
	}
	dict, err := ctx.DereferenceDict(*ctx.XRefTable.Info)
	if err != nil {
		return fmt.Errorf("error dereferencing info dict: %v", err)
	}

	fmt.Println("Metadata:")
	for k, v := range dict {
		fmt.Printf("  %s = %s\n", k, v.String())
	}
	return nil
}
