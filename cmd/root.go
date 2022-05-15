package cmd

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/farbodsalimi/dokimi/internal/configs"
)

var (
	// Persistent Flags
	reporter string
	in       string
	out      string

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

	rootCmd.PersistentFlags().StringVarP(&reporter, "reporter", "r", "", "Reporter name e.g. istanbul, lcov, ...")
	rootCmd.PersistentFlags().StringVarP(&in, "in", "i", "", "Path to input file")
	rootCmd.PersistentFlags().StringVarP(&out, "out", "o", "", "Path to output file")

	viper.BindPFlag("reporter", rootCmd.PersistentFlags().Lookup("reporter"))
	viper.BindPFlag("in", rootCmd.PersistentFlags().Lookup("in"))
	viper.BindPFlag("out", rootCmd.PersistentFlags().Lookup("out"))

	rootCmd.AddCommand(covgenCmd)
	rootCmd.AddCommand(showCmd)
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
