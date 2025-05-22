package imagebuildah

import (
	"fmt"
	"os"
	"io"
	"github.com/containers/image/v5/docker/reference"

)

const (
)

func bakeFiles(dfile string, print bool)(id string, ref reference.Canonical, err error) {
	
	// var data io.Reader
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
	// data = contents

	// If print flag simply print output of bakefile
	if(print){
		b, _ := io.ReadAll(contents)
		fmt.Print(b)
	}




	return "finished", nil, fmt.Errorf("Correct file in %q", dfile)

}
