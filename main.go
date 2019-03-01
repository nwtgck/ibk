package main

import (
	"fmt"
	"github.com/nwtgck/ibk/util"
	"path/filepath"
	"sort"
)

// TODO: impl
func backup(){
}

func restore() error {
	// TODO: Hard code
	dstDirPath := "./mydst"
	err := util.Chdir(dstDirPath, func() error {
		// Find .snar files
		matches, err := filepath.Glob("**.snar")
		if err != nil {
			return err
		}
		// Get snar file name
		// TODO: Should handle empty slice
		snarFileName := matches[0]
		// Get tar file names
		tarFileNames, err := filepath.Glob("**.tar")
		if err != nil {
			return err
		}
		// Sort tar file names
		sort.Strings(tarFileNames)
		for _, tarFileName := range(tarFileNames) {
			command := fmt.Sprintf("gtar -g %s -xf %s", snarFileName, tarFileName)
			_, err := util.EchoRunCmdStr(command)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func main() {
	restore()
}
