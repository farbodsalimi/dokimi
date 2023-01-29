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
		Short: "Dokimi provides some helper commands for testing in Go.",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(checkCoverageCmd)
}

func initConfig() {
	if _, err := os.Stat(configs.DokimiHomeDir); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating", configs.DokimiHomeDir, "...")
		err := os.MkdirAll(configs.DokimiHomeDir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat(configs.IstanbulTmpDir); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating", configs.IstanbulTmpDir, "...")
		err := os.MkdirAll(configs.IstanbulTmpDir, os.ModePerm)
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
