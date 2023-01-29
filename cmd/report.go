package cmd

import (
	"fmt"
	"strconv"

	istanbulReporter "github.com/farbodsalimi/dokimi/internal/reporters/istanbul"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewReportCmd() *cobra.Command {
	var reporter string
	var rInput string
	var rOutput string
	var show bool

	reportCmd := &cobra.Command{
		Use:   "report",
		Short: "Writes reports for Go coverage profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infof("Reporter:\t%s", reporter)
			log.Infof("Input:\t%s", rInput)
			log.Infof("Output:\t%s", rOutput)
			log.Infof("Show:\t%s", strconv.FormatBool(show))

			switch reporter {
			case "istanbul":
				istanbul, err := istanbulReporter.New()
				if err != nil {
					return err
				}

				if show {
					istanbul.ShowReport(rInput, rOutput)
				} else {
					istanbul.WriteReport(rInput, rOutput)
				}

			default:
				return fmt.Errorf("unknown reporter: %s", reporter)
			}

			return nil
		},
	}

	reportCmd.Flags().BoolVar(&show, "show", false, "Shows written reports")
	reportCmd.Flags().
		StringVarP(&reporter, "reporter", "r", "", "Reporter name e.g. istanbul, lcov, ...")
	reportCmd.Flags().StringVarP(&rInput, "input", "i", "", "Path to input file")
	reportCmd.Flags().StringVarP(&rOutput, "output", "o", "", "Path to output file")
	reportCmd.MarkFlagRequired("reporter")
	reportCmd.MarkFlagRequired("rInput")

	return reportCmd
}
