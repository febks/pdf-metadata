package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func UpdateMetadata() error {
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

	infoDict := types.Dict{}
	// Ambil metadata dari ENV
	if title := os.Getenv("TITLE"); title != "" {
		infoDict["Title"] = types.StringLiteral(title)
	}
	if author := os.Getenv("AUTHOR"); author != "" {
		infoDict["Author"] = types.StringLiteral(author)
	}
	if creator := os.Getenv("CREATOR"); creator != "" {
		infoDict["Creator"] = types.StringLiteral(creator)
	}
	infoDict["ModDate"] = types.StringLiteral(types.DateString(time.Now()))

	ir, err := ctx.IndRefForNewObject(infoDict)
	if err != nil {
		return fmt.Errorf("error creating new Info object: %v", err)
	}
	ctx.XRefTable.Info = ir

	if err := api.WriteContextFile(ctx, outFile); err != nil {
		return fmt.Errorf("error writing pdf: %v", err)
	}
	fmt.Println("✅ Metadata updated ->", outFile)
	return nil
}
