package istanbul

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/farbodsalimi/dokimi/internal/configs"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

func (istanbul *Istanbul) ShowReport(input string, output string) error {
	err := istanbul.WriteReport(
		input,
		configs.IstanbulTmpJsonPath,
	)
	if err != nil {
		return err
	}

	exeCmd := exec.Command("istanbul", "report",
		"--include", configs.IstanbulTmpJsonPath,
		"--dir", configs.IstanbulTmpDir, "html",
	)

	cmdOutput, err := exeCmd.CombinedOutput()
	if err != nil {
		log.WithError(err).WithField("output", string(cmdOutput)).
			Error("failed to run 'istanbul report'")
		log.Info("if istanbul is not installed, run: npm i -g istanbul")

		table := tablewriter.NewWriter(os.Stdout)
		table.Append([]string{"npm i -g istanbul"})
		table.Render()

		return err
	}

	exeCmd = exec.Command("open", configs.IstanbulTmpIndexPath)
	err = exeCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open report at %s: %w", configs.IstanbulTmpIndexPath, err)
	}

	return nil
}
