package istanbul

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/farbodsalimi/dokimi/internal/reporters/std"
)

func (istanbul *Istanbul) WriteReport(input string, output string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	absoluteDir, projectFolder := path.Split(wd)

	log.Printf("Directory: %s", absoluteDir)
	log.Printf("Project: %s", projectFolder)

	istanbulObject := map[string]IstanbulObject{}

	profiles, err := std.ParseProfiles(input)
	if err != nil {
		return err
	}

	for _, p := range profiles {
		relativePath := strings.Split(p.FileName, projectFolder)[1]
		absolutePath := path.Join(absoluteDir, projectFolder, relativePath)

		istanbulObj := IstanbulObject{
			Path:         absolutePath,
			StatementMap: map[string]IstanbulStatementMap{},
			FnMap:        map[string]string{},
			BranchMap:    map[string]string{},
			S:            map[string]int{},
			F:            map[string]string{},
			B:            map[string]string{},
		}

		for i, b := range p.Blocks {
			istanbulObj.S[strconv.Itoa(i)] = b.Count
			istanbulObj.StatementMap[strconv.Itoa(i)] = IstanbulStatementMap{
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

		istanbulObject[absolutePath] = istanbulObj
	}

	file, err := json.MarshalIndent(istanbulObject, "", " ")
	if err != nil {
		return err
	}

	err = istanbul.writeFile(output, file, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
