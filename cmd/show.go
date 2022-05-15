package cmd

import (
	"os"
	"os/exec"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/farbodsalimi/dokimi/internal/configs"
	"github.com/farbodsalimi/dokimi/internal/parsers"
)

func init() {
	showCmd.PersistentFlags().StringVarP(&reporter, "reporter", "r", "", "Reporter name e.g. istanbul, lcov, ...")
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows coverage files in reporter formats e.g istanbul",
	Run: func(cmd *cobra.Command, args []string) {
		reporter := viper.GetString("reporter")
		in := viper.GetString("in")

		log.Infof("Reporter:\t%s", reporter)
		log.Infof("Input:\t%s", in)

		switch reporter {
		case "istanbul":
			//
			parsers.IstanbulGenerator(
				in,
				configs.IstanbulTmpJsonPath,
			)

			//
			exeCmd := exec.Command("istanbul", "report",
				"--include", configs.IstanbulTmpJsonPath,
				"--dir", configs.IstanbulTmpDir, "html",
			)

			err := exeCmd.Run()
			if err != nil {
				log.Errorln("Please make sure you have istanbul globally installed on your local machine")

				table := tablewriter.NewWriter(os.Stdout)
				table.Append([]string{"npm i -g istanbul"})
				table.Render()

				log.Fatal(err)
			}

			exeCmd = exec.Command("open", configs.IstanbulTmpIndexPath)
			err = exeCmd.Run()
			if err != nil {
				log.Fatal(err)
			}

		default:
			log.Fatalf("Unknown reporter: %s", reporter)
		}

	},
}
