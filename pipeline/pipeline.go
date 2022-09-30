package pipeline

import (
	"fmt"
	"golangCI/prepare"
	"log"
	"os"
	"sync"
)

type Parameters struct {
	GitRepositoy    string `json:repository`
	SshLocation     string `json:repository`
	BranchName      string `json:branch`
	GitLocationRepo string
}

var GlobalParametes Parameters

func (*Parameters) GetSSHLocation() (*os.File, error) {
	if GlobalParametes.SshLocation == "" {
		return nil, fmt.Errorf("ssh location is empty")
	}
	open, err := os.Open(GlobalParametes.SshLocation)
	if err != nil {
		return nil, fmt.Errorf("open ssh file error:%s", err.Error())
	}
	return open, nil
}

func (*Parameters) GetBranchName() string {
	if GlobalParametes.BranchName == "" {
		return "main"
	}
	return GlobalParametes.BranchName
}

func (*Parameters) GetGitRepository() string {

	return GlobalParametes.GitRepositoy
}

var once sync.Once

func (*Parameters) GetLocationRepository() (string, error) {
	var methodError error
	once.Do(func() {
		location, err := prepare.GetSystemLocation()
		if err != nil {
			methodError = err
			return
		}
		GlobalParametes.GitLocationRepo = location + "CI"
		err = PullRepository(GlobalParametes.GitLocationRepo)
		if err != nil {
			methodError = err
			return
		}
	})
	return GlobalParametes.GitLocationRepo, methodError
}
func PullRepository(location string) error {
	_, err := os.Stat(location)
	if os.IsNotExist(err) {
		err = os.Mkdir(location, 0755)
		if err != nil {
			return fmt.Errorf("create dir error:%s", err.Error())
		}
	}

	output, err := CreateWorkCMD(location, prepare.GlobalCommand, "git init").Output()

	if err != nil {
		return fmt.Errorf("init git repository error:%s", err.Error())
	}

	log.Printf("%s init git repository success:%s", location, output)
	_, err = CreatePreCMD(fmt.Sprintf("git remote add origin %s", GlobalParametes.GetGitRepository())).Output()
	if err != nil {
		return fmt.Errorf("git remote add origin error:%s", err.Error())
	}
	log.Printf("%s git remote add origin repository success", location)
	_, err = CreatePreCMD(fmt.Sprintf("git pull  origin   %s", GlobalParametes.GetBranchName())).Output()
	if err != nil {
		return fmt.Errorf("git pull remote repository error:%s", err.Error())
	}
	log.Printf("%s pull remote repository success", location)
	return nil
}
