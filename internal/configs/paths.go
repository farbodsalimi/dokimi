package configs

import (
	"path"

	log "github.com/sirupsen/logrus"

	"github.com/mitchellh/go-homedir"
)

var (
	DokimiHomeDir        string
	IstanbulTmpDir       string
	IstanbulTmpJsonPath  string
	IstanbulTmpIndexPath string
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	DokimiHomeDir = path.Join(home, ".dokimi")
	IstanbulTmpDir = path.Join(DokimiHomeDir, "istanbul_tmp")
	IstanbulTmpJsonPath = path.Join(IstanbulTmpDir, "istanbul_tmp.json")
	IstanbulTmpIndexPath = path.Join(IstanbulTmpDir, "index.html")
}
