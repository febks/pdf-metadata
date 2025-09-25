package usecase

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func RemoveMetadata() error {
	inFile := os.Getenv("INPUT_FILE")
	outFile := os.Getenv("OUTPUT_FILE")
	if inFile == "" || outFile == "" {
		return fmt.Errorf("INPUT_FILE and OUTPUT_FILE envs are required")
	}

	ctx, err := api.ReadContextFile(inFile)
	if err != nil {
		return fmt.Errorf("error reading context: %v", err)
	}
	ctx.XRefTable.Info = nil
	if ctx.RootDict != nil {
		delete(ctx.RootDict, "Metadata")
	}

	if err := api.WriteContextFile(ctx, outFile); err != nil {
		return fmt.Errorf("error writing pdf: %v", err)
	}
	fmt.Println("✅ Metadata removed ->", outFile)

	if err := ListMetadata(outFile); err != nil {
		return fmt.Errorf("error listing metadata: %v", err)
	}
	return nil
}
