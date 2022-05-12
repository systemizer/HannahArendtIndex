package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/systemizer/ArendtArchives/backend/ocr"
)

func init() {
	rootCmd.AddCommand(ocrCmd)
}

var ocrCmd = &cobra.Command{
	Use:   "ocr",
	Short: "Runs OCR on LOC images in DB",
	Run:   runOCR,
}

func runOCR(cmd *cobra.Command, args []string) {
	o, err := ocr.New()
	if err != nil {
		log.Fatal(err)
	}

	err = o.Start()
	if err != nil {
		log.Fatal(err)
	}
}
