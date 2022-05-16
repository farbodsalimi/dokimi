package cmd

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/farbodsalimi/dokimi/internal/reporters"
)

func init() {
	reportCmd.Flags().BoolVar(&show, "show", false, "Shows written reports")
	reportCmd.Flags().StringVarP(&reporter, "reporter", "r", "", "Reporter name e.g. istanbul, lcov, ...")
	reportCmd.Flags().StringVarP(&rInput, "input", "i", "", "Path to input file")
	reportCmd.Flags().StringVarP(&rOutput, "output", "o", "", "Path to output file")
	reportCmd.MarkFlagRequired("reporter")
	reportCmd.MarkFlagRequired("rInput")
}

var (
	reporter string
	rInput   string
	rOutput  string
	show     bool

	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Writes reports for Go coverage profiles",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("Reporter:\t%s", reporter)
			log.Infof("Input:\t%s", rInput)
			log.Infof("Output:\t%s", rOutput)
			log.Infof("Show:\t%s", strconv.FormatBool(show))

			switch reporter {
			case "istanbul":
				if show {
					reporters.ShowIstanbulReports(rInput, rOutput)
				} else {
					reporters.WriteIstanbulReports(rInput, rOutput)
				}
			default:
				log.Fatalf("Unknown reporter: %s", reporter)
			}
		},
	}
)
