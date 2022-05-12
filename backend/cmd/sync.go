package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/systemizer/ArendtArchives/backend/sync"
)

func init() {
	rootCmd.AddCommand(syncCommand)
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync Data from LOC",
	Run:   startSync,
}

func startSync(cmd *cobra.Command, args []string) {
	s, err := sync.New()
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
