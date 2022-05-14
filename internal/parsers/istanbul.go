package parsers

import (
	log "github.com/sirupsen/logrus"

	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
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

func IstanbulGenerator(in string, out string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	absoluteDir, projectFolder := path.Split(wd)

	log.Printf("Directory: %s", absoluteDir)
	log.Printf("Project: %s", projectFolder)

	istanbul := map[string]Istanbul{}

	profiles, err := ParseProfiles(in)
	if err != nil {
		log.Errorln(err)
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
		log.Errorln(err)
	}

	err = ioutil.WriteFile(out, file, 0644)
	if err != nil {
		log.Errorln(err)
	}
}
