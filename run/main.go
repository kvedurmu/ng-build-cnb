package main

import (
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/chronos"
	"github.com/paketo-buildpacks/packit/scribe"
	ngbuildcnb "github.com/kvedurmu/ng-build-cnb"
	"os"
)

func main() {
	logger := scribe.NewLogger(os.Stdout)
	packit.Run(ngbuildcnb.Detect(), ngbuildcnb.Build(logger, chronos.DefaultClock))
}

