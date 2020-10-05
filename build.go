package ngbuildcnb

import (
	"fmt"
	"strings"
	"bytes"
	"os"
	"time"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/chronos"
	"github.com/paketo-buildpacks/packit/pexec"
	"github.com/paketo-buildpacks/packit/scribe"
)

func Build(logger scribe.Logger, clock chronos.Clock) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)
		workingDir := context.WorkingDir
		
		executable := pexec.NewExecutable("ng")
		logger.Process("Executing build process")
		buffer := bytes.NewBuffer(nil)
		args := []string{"build", "--prod", "--verbose"}
		command := fmt.Sprintf("ng %s", strings.Join(args, " "))
		logger.Subprocess("Running '%s'", command)
		
		duration, err := clock.Measure(func() error {
			return executable.Execute(pexec.Execution{
					Args:   args,
					Dir:    workingDir,
					Stdout: buffer,
					Stderr: buffer,
					Env:    os.Environ(),
				})	
		})

		if err != nil {
			return packit.BuildResult{}, fmt.Errorf("Failed to run '%s'\nerror: %s", command, buffer.String())
		}

		logger.Action("Completed in %s", duration.Round(time.Millisecond))
		logger.Break()	

		return packit.BuildResult{}, nil
	}
}