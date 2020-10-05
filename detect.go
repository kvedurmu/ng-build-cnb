package ngbuildcnb

import (
	"path/filepath"
	"fmt"
	"os"
	"encoding/json"
	"github.com/paketo-buildpacks/packit"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		dependencies, err := parseDependencies(filepath.Join(context.WorkingDir, "package.json"))
		if err != nil {
			return packit.DetectResult{}, err
		}
		
		if _, ok := dependencies["@angular/cli"]; !ok {
			return packit.DetectResult{}, fmt.Errorf("Missing @angular/cli in package.json")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "ng_static"},
				},
				Requires: []packit.BuildPlanRequirement{
					{Name: "ng_static"},
					{
						Name: "node_modules",
						Metadata: map[string]interface{}{
							"build": true,	
						},
					},
				},
			},
		}, nil
	}
}

func parseDependencies(packageJsonPath string) (map[string]string, error) {
	file, err := os.Open(packageJsonPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pkg struct {
		Dependencies map[string]string `json:"dependencies"`
	}

	err = json.NewDecoder(file).Decode(&pkg)
	if err != nil {
		return nil, err
	}

	return pkg.Dependencies, nil
}