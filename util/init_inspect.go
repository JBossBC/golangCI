package util

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Operation_System int

const (
	LINUX Operation_System = iota
	WINDOWS
	Default
)

var CurrentSystem Operation_System = Default

var GlobalCommand string
var GlobalParams string

func InspectGo() (string, error) {
	output, err := exec.CommandContext(context.TODO(), GlobalCommand, GlobalParams, "go version").Output()
	if err != nil {
		return string(output), err
	}
	log.Printf("The current GO version is:%s", output)
	return string(output), nil

}
func InspectGit() (string, error) {
	output, err := exec.CommandContext(context.TODO(), GlobalCommand, GlobalParams, "git version").Output()
	if err != nil {
		return string(output), err
	}
	log.Printf("The current GO version is:%s", output)
	return string(output), nil
}
func ConvertSystemToCommand() error {
	if CurrentSystem == Default {
		return fmt.Errorf("the current version does not support this operating system:%s", os.Getenv("os"))
	}
	if CurrentSystem == WINDOWS {
		GlobalCommand = "cmd"
		GlobalParams = "/c"
		return nil
	}
	if CurrentSystem == LINUX {
		GlobalCommand = "/bin/bash"
		GlobalParams = "-c"
		return nil
	}
	return nil
}

func GetSystemInformation() (string, error) {
	operation := os.Getenv("os")
	log.Printf("The Current Operation System is %s\n", operation)
	if strings.Contains(operation, "Windows") {
		CurrentSystem = WINDOWS
		return operation, nil
	}
	if strings.Contains(operation, "Linux") {
		CurrentSystem = LINUX
		return operation, nil
	}
	return operation, fmt.Errorf("the current version does not support this operating system:%s", operation)
}
func InspectDocker() (string, error) {

	return "", nil
}
