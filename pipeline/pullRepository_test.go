package pipeline

import (
	"fmt"
	"golangCI/prepare"
	"os/exec"
	"testing"
)

func TestPullRepository(t *testing.T) {
	_, err := prepare.GetSystemInformation()
	if err != nil {
		panic(err.Error())
	}
	err = prepare.ConvertSystemToCommand()
	if err != nil {
		panic(err.Error())
	}
	GlobalParametes.GitRepositoy = "git@github.com:JBossBC/golangCI.git"
	err = PullRepository("C://var/CI")
	if err != nil {
		return
	}
}

func TestCMD(t *testing.T) {

	cmd := exec.Command("cmd", "/c", "git init")
	cmd.Dir = "C:/var/CI"
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
