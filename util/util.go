package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func EchoRunCommand(name string, args ...string) ([]byte, error) {
	var cmdList []string
	for _, x := range append([]string{name}, args...) {
		// (%#v from: https://stackoverflow.com/a/50054686/2885946)
		cmdList = append(cmdList, fmt.Sprintf("%#v", x))
	}
	cmdListStr := fmt.Sprintf(
		"CMD [ %s ]",
		strings.Join(cmdList, ", "),
	)
	fmt.Println(cmdListStr)
	return exec.Command(name, args...).Output()
}

func Chdir(dirpath string, process func() error) error {
	err := os.Chdir(dirpath)
	if err != nil {
		return err
	}
	defer os.Chdir(dirpath)
	return process()
}
