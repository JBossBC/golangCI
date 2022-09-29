package prepare

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
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
	timeout, cancelFunc := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancelFunc()
	output, err := exec.CommandContext(timeout, GlobalCommand, GlobalParams, "go version").Output()
	if err != nil {
		return string(output), fmt.Errorf("unable to detect your golang environment")
	}
	log.Printf("The current GO version is:%s", output)
	return string(output), nil

}
func InspectGit() (string, error) {
	timeout, cancelFunc := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancelFunc()
	output, err := exec.CommandContext(timeout, GlobalCommand, GlobalParams, "git version").Output()
	if err != nil {
		return string(output), fmt.Errorf("unable to detect your git environment")
	}
	log.Printf("The current GO version is:%s", output)
	return string(output), nil
}
func InspectDocker() (string, error) {
	timeout, cancelFunc := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancelFunc()
	output, err := exec.CommandContext(timeout, GlobalCommand, GlobalParams, "docker version").Output()
	if err != nil {
		return string(output), fmt.Errorf("unable to detect your docker environment")
	}
	log.Printf("The current docker version is:\n%s", output)
	//初略判断docker server是否启动
	contains := strings.Contains(string(output), "Server:")
	if !contains {
		log.Printf("[warning]:docker server maybe dont start")
	}
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

/**
  返回该文件系统支持的默认文件位置
*/
func GetSystemLocation() (string, error) {
	if CurrentSystem == Default {
		return "", fmt.Errorf("the operation system cant support:%s", os.Getenv("os"))
	}
	if CurrentSystem == LINUX {
		return "/usr/", nil
	}
	if CurrentSystem == WINDOWS {
		return "/var/", nil
	}
	return "", fmt.Errorf("programming occur panic error")
}
