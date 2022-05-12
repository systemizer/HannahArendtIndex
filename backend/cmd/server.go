package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/systemizer/ArendtArchives/backend/server"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Server",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
}
