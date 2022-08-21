package istanbul

import (
	"os"
	"os/exec"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/farbodsalimi/dokimi/internal/configs"
)

func (istanbul *Istanbul) ShowReport(input string, output string) {
	//
	istanbul.WriteReport(
		input,
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
}
