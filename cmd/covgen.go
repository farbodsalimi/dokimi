package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/farbodsalimi/dokimi/internal/parsers"
)

var (
	reporter string
	in       string
	out      string
)

func init() {
	rootCmd.AddCommand(covgenCmd)

	covgenCmd.PersistentFlags().StringVarP(&reporter, "reporter", "r", "", "Reporter name e.g. istanbul, lcov, ...")
	covgenCmd.PersistentFlags().StringVarP(&in, "in", "i", "", "Path to input file")
	covgenCmd.PersistentFlags().StringVarP(&out, "out", "o", "", "Path to output file")
}

var covgenCmd = &cobra.Command{
	Use:   "covgen",
	Short: "Generates coverage files in different formats.",
	Run: func(cmd *cobra.Command, args []string) {
		reporter := cmd.PersistentFlags().Lookup("reporter").Value
		in := cmd.PersistentFlags().Lookup("in").Value
		out := cmd.PersistentFlags().Lookup("out").Value

		log.Infof("Reporter:\t%s", reporter)
		log.Infof("Input:\t%s", in)
		log.Infof("Output:\t%s", out)

		switch reporter.String() {
		case "istanbul":
			parsers.IstanbulGenerator(in.String(), out.String())
		default:
			log.Fatalf("Unknown reporter: %s", reporter)
		}
	},
}
