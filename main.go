package main

import (
	"github.com/farbodsalimi/dokimi/cmd"
	log "github.com/sirupsen/logrus"
)

// Version is the version of this binary. It'll be overridden as part of the build process through linker flags.
var Version = "development"

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})
}

func main() {
	err := cmd.Execute(Version)
	if err != nil {
		log.Fatalln(err)
	}
}
