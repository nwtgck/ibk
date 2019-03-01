package util

import (
	"fmt"
	"github.com/mattn/go-shellwords"
	"os"
	"os/exec"
)

// Run command and return the output
// (from: https://qiita.com/tanksuzuki/items/9205ff70c57c4c03b5e5)
func RunCmdStr(cmdstr string) ([]byte, error) {
	var out []byte
	// Parse the command string
	c, err := shellwords.Parse(cmdstr)
	if err != nil {
		return out, err
	}
	switch len(c) {
	// Empty string
	case 0:
		return out, nil
	// Only command without options
	case 1:
		out, err = exec.Command(c[0]).CombinedOutput()
	// Command with options
	default:
		out, err = exec.Command(c[0], c[1:]...).CombinedOutput()
	}
	if err != nil {
		return out, err
	}
	return out, nil
}


// Print command and run the command
func EchoRunCmdStr(cmdstr string) ([]byte, error) {
	// Print command
	fmt.Printf("+ %s\n", cmdstr)
	res, err := RunCmdStr(cmdstr)
	return res, err
}

func Chdir(dirpath string, process func() error) error {
	err := os.Chdir(dirpath)
	if err != nil {
		return err
	}
	defer os.Chdir(dirpath)
	return process()
}
