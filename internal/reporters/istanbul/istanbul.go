package istanbul

import (
	"io/fs"
	"os"

)

type Istanbul struct {
	writeFileFn func(filename string, data []byte, perm fs.FileMode) error
}

func (i Istanbul) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	writeFn := i.writeFileFn
	if writeFn == nil {
		writeFn = os.WriteFile
	}
	return writeFn(filename, data, perm)
}

func New(opts ...func(*Istanbul)) (*Istanbul, error) {
	istanbul := &Istanbul{}
	for _, opt := range opts {
		opt(istanbul)
	}

	return istanbul, nil
}

type IstanbulStatementStartEnd struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type IstanbulStatementMap struct {
	Start IstanbulStatementStartEnd `json:"start"`
	End   IstanbulStatementStartEnd `json:"end"`
}

type IstanbulFnMapEntry struct {
	Name string              `json:"name"`
	Line int                 `json:"line"`
	Loc  IstanbulStatementMap `json:"loc"`
}

type IstanbulBranchMapEntry struct {
	Type      string                `json:"type"`
	Line      int                   `json:"line"`
	Locations []IstanbulStatementMap `json:"locations"`
}

type IstanbulObject struct {
	Path         string                             `json:"path"`
	StatementMap map[string]IstanbulStatementMap     `json:"statementMap"`
	FnMap        map[string]IstanbulFnMapEntry       `json:"fnMap"`
	BranchMap    map[string]IstanbulBranchMapEntry   `json:"branchMap"`
	S            map[string]int                      `json:"s"`
	F            map[string]int                      `json:"f"`
	B            map[string][]int                    `json:"b"`
}
