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
	GlobalParametes.GitRepositoy = ""
	err = PullRepository("C://var/CI")
	if err != nil {
		return
	}
}

func TestCMD(t *testing.T) {
	output, err := exec.Command("cmd ", "cd C://var/CI", " git init").Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(output))
}
