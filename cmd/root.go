package cmd

import (
	"errors"
	"os"

	"github.com/farbodsalimi/dokimi/internal/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

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

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dokimi",
		Short: "Dokimi provides some helper commands for testing in Go.",
	}

	rootCmd.AddCommand(NewReportCmd())
	rootCmd.AddCommand(NewCheckCoverageCmd())

	return rootCmd
}

func Execute(version string) error {
	rootCmd := NewRootCommand()
	rootCmd.Version = version

	cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
