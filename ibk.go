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

func Backup(targetDirPath string, dstDirPath string, now time.Time) error {
	// Get the base name of the target directory
	targetBaseName := filepath.Base(targetDirPath)
	// Difine snar name
	targetSnarName := fmt.Sprintf("%s.snar", targetBaseName)

	// Create destination path if it doesn't exist
	os.MkdirAll(dstDirPath, os.ModePerm)

	// Define tar file name
	tarFileName := fmt.Sprintf("%s_%s.tar", targetBaseName, formatTime(now))

	targetSnarPath := filepath.Join(dstDirPath, targetSnarName)
	tarFilePath := filepath.Join(dstDirPath, tarFileName)

	// Incremental backup
	_, err := util.EchoRunCommand(
		"gtar",
		"-g",
		targetSnarPath,
		"-cf",
		tarFilePath,
		targetDirPath,
	)
	return err
}

func Restore(dstDirPath string) error {
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
			_, err := util.EchoRunCommand(
				"gtar",
				"-g",
				snarFileName,
				"-xf",
				tarFileName,
			)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

