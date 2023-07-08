package cmd

import (
	"github.com/krobus00/asynqmon-multi-service/internal/bootstrap"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run asynq monitoring server for multiple services",
	Long:  `Run asynq monitoring server for multiple services`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartServerBootstrap()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
