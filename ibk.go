package ibk

import (
	"fmt"
	"github.com/nwtgck/ibk/util"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// Format time for file name
func formatTime(t time.Time) string {
	zone, _ := t.Zone()
	// NOTE: Tried to use time#Format(), but it is very weird implementation.
	// Third-party libraries seem not to be maintained.
	return fmt.Sprintf(
		"%d%02d%02d_%02d%02d_%02d_%03d_%s",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(),
		t.Second(),
		t.Nanosecond() / 1000000,
		zone,
	)
}

func Backup() error {
	// TODO: Hard code
	targetDirPath := "./mydir"
	// TODO: Hard code
	dstDirPath := "./mydst"
	// Get the base name of the target directory
	targetBaseName := filepath.Base(targetDirPath)
	// Difine snar name
	targetSnarName := fmt.Sprintf("%s.snar", targetBaseName)

	// Create destination path if it doesn't exist
	os.MkdirAll(dstDirPath, os.ModePerm)

	// Define tar file name
	tarFileName := fmt.Sprintf("%s_%s.tar", targetBaseName, formatTime(time.Now()))

	targetSnarpath := filepath.Join(dstDirPath, targetSnarName)
	tarFilePath := filepath.Join(dstDirPath, tarFileName)

	// Incremental backup
	_, err := util.EchoRunCmdStr(fmt.Sprintf(
		"gtar -g %s -cf %s %s",
		targetSnarpath,
		tarFilePath,
		targetDirPath,
	))
	return err
}

func Restore() error {
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

