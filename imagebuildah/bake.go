package imagebuildah

import (
	"bytes"
	"fmt"
	"os"
	"io"
	"slices"
	"context"
 	"github.com/containers/image/v5/docker/reference"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/containers/buildah/define"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/containers/storage"
	// gotypes "go/types"

	// "github.com/containers/image/v5/types"
)

const (
)

type Bakefile = define.BakeFile
type GroupOptions = define.GroupOptions
type TargetOptions = define.TargetOptions

func BakeFiles(ctx context.Context, store storage.Store, dfile string, print bool, inputArgs ...string)(id string, ref reference.Canonical, err error) {
	
	var data io.Reader
	dinfo, err := os.Stat(dfile)

	if err != nil {
		return "", nil, err
	}

	var contents *os.File
	// If given a directory error out since `-f` does not supports path to directory
	if dinfo.Mode().IsDir() {
		return "", nil, fmt.Errorf("containerfile: %q cannot be path to a directory", dfile)
	}
	contents, err = os.Open(dfile)
	if err != nil {
		return "", nil, fmt.Errorf("reading bake instructions: %w", err)
	}
	defer contents.Close()
	dinfo, err = contents.Stat()
	if err != nil {
		return "", nil, fmt.Errorf("reading info about %q: %w", dfile, err)
	}
	if dinfo.Mode().IsRegular() && dinfo.Size() == 0 {
		return "", nil, fmt.Errorf("no contents in %q", dfile)
	}
	data = contents

	// If print flag simply print output of bakefile
	var b bytes.Buffer
	if _, err := b.ReadFrom(data); err != nil {
		return "", nil, err
	}

	if print {
		fmt.Printf("%s", b.String())
	}

	// REMOVE THIS
	if(context.Cause(ctx) == nil){
		fmt.Printf("No errors yippee \n")
	}


	bakefile, diag := parseBakefileHCL(dfile)

	if(diag != nil){
		return "", nil, fmt.Errorf("error parsing HCL file %s", diag[0].Summary)
	}	

	// Check for default 
	if (len(inputArgs) == 0){
		inputArgs = append(inputArgs, "default")
	}

	groups := bakefile.Groups
	targets := bakefile.Targets
	// Groups and targets specified in the input parameters
	var targetGroups []GroupOptions
	var targetTargets []string
	

	for _, input := range inputArgs {
		flag := false
		for _, group := range groups {
			if(group.GroupName == input){
				flag = true
				targetGroups = append(targetGroups, group)
				break
			} 
		}
		if(!flag){
			for _, target := range targets {
				if(target.TargetName == input){
					targetTargets = append(targetTargets, target.TargetName)
					break
				} 
			}
		}
	} 

	// If no default group and default asked for build everything
	if(len(targetGroups) == 0 && len(targetTargets) == 0){
		for _, group := range groups {
			targetGroups = append(targetGroups, group)
		}
		for _, target := range targets {
			targetTargets = append(targetTargets, target.TargetName)
		}
	}

	// DELETE THIS
	for _, i := range targets {
		fmt.Printf("%s \n", i.TargetName)
	}
	for _, tG := range targetGroups {
		fmt.Printf("%s \n", tG)
	}

	dedupTargets(targets, targetGroups, targetTargets)


	return "finished", nil, fmt.Errorf("Correct file in %q", dfile)
}

// parseBakefileHCL uses the hashicorp HCL library to convert the HCL file
// into go structs
func parseBakefileHCL(path string) (Bakefile, hcl.Diagnostics) {
	var bakefile Bakefile

	hclparser := hclparse.NewParser()
	result, diag := hclparser.ParseHCLFile(path)

	if diag.HasErrors() {
		return bakefile, diag
	}
	
	ctx := &hcl.EvalContext{}
	gohcl.DecodeBody(result.Body, ctx, &bakefile)

	return bakefile, nil
}

// dedup processes the list of given targets/groups and returns a deduplicated version with just targets
func dedupTargets(targets []TargetOptions, targetGroups []GroupOptions, targetTargets []string) (FinT []TargetOptions){
	
	var finT []TargetOptions
	finTString := targetTargets

	for _, tG := range targetGroups {
		for _, target := range tG.GroupMembers{
			if(!(slices.Contains(finTString, target))){
				finTString = append(finTString, target)
			}
		}
	}

	for _, t := range targets {
		if(slices.Contains(finTString, t.TargetName)){
			finT = append(finT, t)
		}
	}

	return finT
}
// Step 2: Setup the context and environments for each of those