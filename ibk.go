package ibk

import (
	"fmt"
	"github.com/nwtgck/ibk/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

// tar command name
func getTarName() string {
	if runtime.GOOS == "darwin" {
		return "gtar"
	} else {
		return "tar"
	}
}

func Backup(srcDirPath string, backupDirPath string, now time.Time) error {
	tarName := getTarName()
	// If tar is not available
	// (from: https://qiita.com/suin/items/2724f15fff9c1948c5b3)
	if _, err := exec.LookPath(tarName); err != nil {
		return fmt.Errorf("%s command is not available\n", tarName)
	}

	// Get the base name of the source directory
	srcBaseName := filepath.Base(srcDirPath)
	// Define snar name
	srcSnarName := fmt.Sprintf("%s.snar", srcBaseName)

	// Create destination path if it doesn't exist
	os.MkdirAll(backupDirPath, os.ModePerm)

	// Define tar file name
	tarFileName := fmt.Sprintf("%s_%s.tar", srcBaseName, formatTime(now))

	backupSnarPath := filepath.Join(backupDirPath, srcSnarName)
	tarFilePath := filepath.Join(backupDirPath, tarFileName)

	// Incremental backup
	_, err := util.EchoRunCommand(
		tarName,
		"-g",
		backupSnarPath,
		"-cf",
		tarFilePath,
		srcDirPath,
	)
	return err
}

func Restore(backupDirPath string, restoredDirPath string) error {
	tarName := getTarName()
	// If tar is not available
	// (from: https://qiita.com/suin/items/2724f15fff9c1948c5b3)
	if _, err := exec.LookPath(tarName); err != nil {
		return fmt.Errorf("%s command is not available\n", tarName)
	}

	// Create restored path if it doesn't exist
	os.MkdirAll(restoredDirPath, os.ModePerm)

	//err := util.Chdir(backupDirPath, func() error {
	// Find .snar files
	matches, err := filepath.Glob(filepath.Join(backupDirPath, "**.snar"))
	if err != nil {
		return err
	}
	// Get snar file name
	// TODO: Should handle empty slice
	snarFilePath := matches[0]
	// Get tar file names
	tarFilePaths, err := filepath.Glob(filepath.Join(backupDirPath, "**.tar"))
	if err != nil {
		return err
	}
	// Sort tar file names
	sort.Strings(tarFilePaths)
	for _, tarFilePath := range(tarFilePaths) {
		_, err := util.EchoRunCommand(
			tarName,
			"-g",
			snarFilePath,
			"-xf",
			tarFilePath,
			"-C",
			restoredDirPath,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

