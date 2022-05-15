package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/farbodsalimi/dokimi/internal/parsers"
)

var covgenCmd = &cobra.Command{
	Use:   "covgen",
	Short: "Generates coverage files in different formats.",
	Run: func(cmd *cobra.Command, args []string) {
		reporter := viper.GetString("reporter")
		in := viper.GetString("in")
		out := viper.GetString("out")

		log.Infof("Reporter:\t%s", reporter)
		log.Infof("Input:\t%s", in)
		log.Infof("Output:\t%s", out)

		switch reporter {
		case "istanbul":
			parsers.IstanbulGenerator(in, out)
		default:
			log.Fatalf("Unknown reporter: %s", reporter)
		}
	},
}
