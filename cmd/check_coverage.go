package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCheckCoverageCmd() *cobra.Command {
	var coverProfile string
	var threshold float64
	var doNotFail bool

	checkCoverageCmd := &cobra.Command{
		Use:   "check-coverage",
		Short: "Checks total coverage against thresholds",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := exec.Command("go", "tool", "cover", "-func", coverProfile).Output()
			if err != nil {
				return fmt.Errorf("failed to run 'go tool cover' on %s: %w", coverProfile, err)
			}

			content := string(result)
			index := strings.Index(content, "total")
			if index == -1 {
				return fmt.Errorf("could not find total coverage in %s", coverProfile)
			}
			line := content[index:]
			re := regexp.MustCompile(`([0-9]*\.?[0-9]*)\s*%`)
			match := re.FindStringSubmatch(line)
			if match == nil {
				return fmt.Errorf("could not parse coverage percentage from %s", coverProfile)
			}
			totalCoverage, err := strconv.ParseFloat(match[1], 32)
			if err != nil {
				return err
			}

			if totalCoverage < threshold {
				msg := fmt.Sprintf(
					"total coverage %.2f is lower than threshold %.2f",
					totalCoverage,
					threshold,
				)
				if doNotFail {
					log.Warn(msg)
				} else {
					return fmt.Errorf("%s", msg)
				}
			}

			return nil
		},
	}

	f := checkCoverageCmd.Flags()
	f.StringVarP(&coverProfile, "coverprofile", "c", "coverage.out", "coverprofile")
	f.Float64VarP(&threshold, "threshold", "t", 100, "threshold")
	f.BoolVarP(&doNotFail, "do-not-fail", "d", false, "do-not-fail")

	return checkCoverageCmd
}
