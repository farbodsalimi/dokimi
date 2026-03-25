package cmd

import (
	"fmt"

	istanbulReporter "github.com/farbodsalimi/dokimi/internal/reporters/istanbul"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewReportCmd() *cobra.Command {
	var (
		reporter string
		input    string
		output   string
		show     bool
	)

	reportCmd := &cobra.Command{
		Use:   "report",
		Short: "Writes reports for Go coverage profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infof("Reporter:\t%s", reporter)
			log.Infof("Input:\t%s", input)
			log.Infof("Output:\t%s", output)
			log.Infof("Show:\t%v", show)

			switch reporter {
			case "istanbul":
				istanbul, err := istanbulReporter.New()
				if err != nil {
					return err
				}

				if show {
					return istanbul.ShowReport(input, output)
				}
				return istanbul.WriteReport(input, output)

			default:
				return fmt.Errorf("unknown reporter: %s", reporter)
			}
		},
	}

	reportCmd.Flags().BoolVar(&show, "show", false, "Shows written reports")
	reportCmd.Flags().
		StringVarP(&reporter, "reporter", "r", "", "Reporter name e.g. istanbul, lcov, ...")
	reportCmd.Flags().StringVarP(&input, "input", "i", "", "Path to input file")
	reportCmd.Flags().StringVarP(&output, "output", "o", "", "Path to output file")
	reportCmd.MarkFlagRequired("reporter")
	reportCmd.MarkFlagRequired("input")

	return reportCmd
}
