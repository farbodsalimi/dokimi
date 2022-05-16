package cmd

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/farbodsalimi/dokimi/internal/configs"
)

var (
	// Root Command
	rootCmd = &cobra.Command{
		Use:   "dokimi",
		Short: "Dokimi is a Go testing framework.",
		Long:  `Dokimi is a Go testing framework.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(checkCoverageCmd)
}

func initConfig() {
	log.Println(configs.DokimiHomeDir)
	if _, err := os.Stat(configs.DokimiHomeDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configs.DokimiHomeDir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorln(err)
		os.Exit(1)
	}
}
