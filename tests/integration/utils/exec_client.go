package utils

import (
	"os"
	"os/exec"
	"strings"

	"github.com/kiali/kiali/log"
)

var ocCommand = NewExecCommand()

func NewExecCommand() string {
	command := os.Getenv("CLIENT_EXE")
	if command != "" {
		return command
	} else {
		return "oc"
	}
}

func ApplyFile(yamlFile, namespace string) bool {
	cmd := exec.Command(ocCommand, "apply", "-n="+namespace, "-f="+yamlFile)
	stdout, err := cmd.Output()

	if err != nil {
		log.Errorf(err.Error())
		return false
	}
	log.Debugf(string(stdout))
	return strings.Contains(string(stdout), "created") || strings.Contains(string(stdout), "configure")
}

func DeleteFile(yamlFile, namespace string) bool {
	cmd := exec.Command(ocCommand, "delete", "-n="+namespace, "-f="+yamlFile)
	stdout, err := cmd.Output()

	if err != nil {
		log.Errorf(err.Error())
		return false
	}
	log.Debugf(string(stdout))
	return strings.Contains(string(stdout), "deleted")
}
