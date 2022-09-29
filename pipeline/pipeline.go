package pipeline

import (
	"fmt"
	"golangCI/prepare"
	"log"
	"os"
	"os/exec"
	"sync"
)

type Parameters struct {
	GitRepositoy    string `json:repository`
	SshLocation     string `json:repository`
	BranchName      string `json:branch`
	GitLocationRepo string
}

var GlobalParametes Parameters

func GetSSHLocation() (*os.File, error) {
	if GlobalParametes.SshLocation == "" {
		return nil, fmt.Errorf("ssh location is empty")
	}
	open, err := os.Open(GlobalParametes.SshLocation)
	if err != nil {
		return nil, fmt.Errorf("open ssh file error:%s", err.Error())
	}
	return open, nil
}

func GetBranchName() string {
	return GlobalParametes.BranchName
}

var once sync.Once

func GetLocationRepository() (string, error) {
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
	file,err:=os.Open(location)
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	stat.IsDir()
	if ;
	err := exec.Command(prepare.GlobalCommand, prepare.GlobalParams, "cd ", location).Run()
	if err != nil {
		return fmt.Errorf("cd the locationRepository error:%s", err.Error())
	}
	output, err := exec.Command(prepare.GlobalCommand, prepare.GlobalCommand, "git init").Output()
	if err != nil {
		return fmt.Errorf("init git repository error:%s", err.Error())
	}
	log.Printf("%s init git repository success:%s", location, output)
	_, err = exec.Command(prepare.GlobalCommand, prepare.GlobalCommand, "git remote add origin ", GlobalParametes.GitRepositoy).Output()
	if err != nil {
		return fmt.Errorf("git remote add origin error:%s", err.Error())
	}
	log.Printf("%s git remote add origin repository success", location)
	_, err = exec.Command(prepare.GlobalCommand, prepare.GlobalCommand, "git pull origin", GlobalParametes.BranchName).Output()
	if err != nil {
		return fmt.Errorf("git pull remote repository error:%s", err.Error())
	}
	log.Printf("%s pull remote repository success", location)
	return nil
}
