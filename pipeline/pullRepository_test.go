package pipeline

import (
	"golangCI/prepare"
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
	err = PullRepository("/var/CI")
	if err != nil {
		return
	}
}
