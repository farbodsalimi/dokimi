package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dokimi",
		Short: "Dokimi is a Go testing framework.",
		Long:  `Dokimi is a Go testing framework.`,
	}
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(covgenCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorln(err)
		os.Exit(1)
	}
}
