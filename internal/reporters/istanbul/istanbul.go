package istanbul

import (
	"io/fs"
	"io/ioutil"
)

type IstanbulStatementStartEnd struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type IstanbulStatementMap struct {
	Start IstanbulStatementStartEnd `json:"start"`
	End   IstanbulStatementStartEnd `json:"end"`
}

type IstanbulObject struct {
	Path         string                          `json:"path"`
	StatementMap map[string]IstanbulStatementMap `json:"statementMap"`
	FnMap        interface{}                     `json:"fnMap"`
	BranchMap    interface{}                     `json:"branchMap"`
	S            map[string]int                  `json:"s"`
	F            interface{}                     `json:"f"`
	B            interface{}                     `json:"b"`
}

type Istanbul struct {
	writeFile func(filename string, data []byte, perm fs.FileMode) error
}

func New(opts ...func(*Istanbul)) (*Istanbul, error) {
	istanbul := &Istanbul{}
	for _, opt := range opts {
		opt(istanbul)
	}

	if istanbul.writeFile == nil {
		istanbul.writeFile = ioutil.WriteFile
	}

	return istanbul, nil
}
