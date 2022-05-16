package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ccInput string
	check   float32

	checkCoverageCmd = &cobra.Command{
		Use:   "check-coverage",
		Short: "Checks total coverage against thresholds",
		Run: func(cmd *cobra.Command, args []string) {
			in := viper.GetString("in")
			check := cmd.Flags().Lookup("check").Value

			log.Infof("Input:\t%s", in)
			log.Infof("Check:\t%s%%", check)
		},
	}
)

func init() {
	checkCoverageCmd.Flags().StringVar(&ccInput, "rInput", "", "rInput")
	checkCoverageCmd.Flags().Float32VarP(&check, "check", "c", 100, "check")
	checkCoverageCmd.MarkFlagRequired("rInput")
}
