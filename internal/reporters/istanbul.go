package reporters

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/farbodsalimi/dokimi/internal/configs"
)

type IstanbulStatementStartEnd struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type IstanbulStatementMap struct {
	Start IstanbulStatementStartEnd `json:"start"`
	End   IstanbulStatementStartEnd `json:"end"`
}

type Istanbul struct {
	Path         string                          `json:"path"`
	StatementMap map[string]IstanbulStatementMap `json:"statementMap"`
	FnMap        interface{}                     `json:"fnMap"`
	BranchMap    interface{}                     `json:"branchMap"`
	S            map[string]int                  `json:"s"`
	F            interface{}                     `json:"f"`
	B            interface{}                     `json:"b"`
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func WriteIstanbulReports(input string, output string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	absoluteDir, projectFolder := path.Split(wd)

	log.Printf("Directory: %s", absoluteDir)
	log.Printf("Project: %s", projectFolder)

	istanbul := map[string]Istanbul{}

	profiles, err := ParseProfiles(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range profiles {
		relativePath := strings.Split(p.FileName, projectFolder)[1]
		absolutePath := path.Join(absoluteDir, projectFolder, relativePath)

		istb := Istanbul{
			Path:         absolutePath,
			StatementMap: map[string]IstanbulStatementMap{},
			FnMap:        map[string]string{},
			BranchMap:    map[string]string{},
			S:            map[string]int{},
			F:            map[string]string{},
			B:            map[string]string{},
		}

		for i, b := range p.Blocks {

			istb.S[strconv.Itoa(i)] = b.Count
			istb.StatementMap[strconv.Itoa(i)] = IstanbulStatementMap{
				Start: IstanbulStatementStartEnd{
					Line:   b.StartLine,
					Column: b.EndCol,
				},
				End: IstanbulStatementStartEnd{
					Line:   b.EndLine,
					Column: b.EndCol,
				},
			}
		}

		istanbul[absolutePath] = istb
	}

	file, err := json.MarshalIndent(istanbul, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(output, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ShowIstanbulReports(input string, output string) {
	//
	WriteIstanbulReports(
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
