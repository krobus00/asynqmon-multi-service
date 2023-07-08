package cmd

import (
	"fmt"
	"os"

	"github.com/krobus00/asynqmon-multi-service/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "asynqmon-multi-service",
	Short: "Asynq monitoring for multiple services",
	Long:  `Asynq monitoring for multiple services`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Init() {
	if err := config.LoadConfig(); err != nil {
		logrus.Fatalln(err.Error())
	}
	logrus.Info(fmt.Sprintf("starting %s:%s...", config.ServiceName(), config.ServiceVersion()))
	logLevel, _ := logrus.ParseLevel(config.LogLevel())
	logrus.SetLevel(logLevel)
}
