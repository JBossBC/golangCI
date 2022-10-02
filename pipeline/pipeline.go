package pipeline

import (
	"fmt"
	"golangCI/prepare"
	"golangCI/util"
	"log"
	"os"
	"strings"
)

type Parameters struct {
	GitRepository  string `json:GitRepository`
	BranchName     string `json:branchName`
	DockerfileName string `json:dockerfileName`
	//SshLocation     string `json:ssh`
	//SshLocation     string
	GitLocationRepo string `json:gitLocationRepo`
}

var GlobalParametes Parameters

//func (*Parameters) GetSSHLocation() (*os.File, error) {
//	if GlobalParametes.SshLocation == "" {
//		return nil, fmt.Errorf("ssh location is empty")
//	}
//	open, err := os.Open(GlobalParametes.SshLocation)
//	if err != nil {
//		return nil, fmt.Errorf("open ssh file error:%s", err.Error())
//	}
//	return open, nil
//}
func (global *Parameters) GetDockerfileName() string {
	var name = global.DockerfileName
	if GlobalParametes.DockerfileName == "" {
		name = "Dockerfile"
	}
	return name
}

func (*Parameters) GetBranchName() string {
	if GlobalParametes.BranchName == "" {
		return "main"
	}
	return GlobalParametes.BranchName
}

func (*Parameters) GetGitRepository() string {

	return GlobalParametes.GitRepository
}

func (*Parameters) GetLocationRepository() (string, error) {
	var methodError error
	location, err := prepare.GetSystemLocation()
	if err != nil {
		methodError = err
		return "", methodError
	}
	GlobalParametes.GitLocationRepo = location + "CI"
	err = PullRepository(GlobalParametes.GitLocationRepo)
	if err != nil {
		methodError = err
		return "", methodError
	}
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

	output, err := util.CreateWorkCMD(location, prepare.GlobalCommand, "git init").Output()

	if err != nil {
		return fmt.Errorf("init git repository error:%s", err.Error())
	}

	log.Printf("%s init git repository success:%s", location, output)
	remoteInfo, err := util.CreatePreCMD("git remote -v").Output()
	if !strings.Contains(string(remoteInfo), GlobalParametes.GitRepository) {
		output, err = util.CreatePreCMD(fmt.Sprintf("git remote add origin %s", GlobalParametes.GetGitRepository())).Output()
		if err != nil {
			return fmt.Errorf("git remote add origin error:%s", output)
		}
	}
	log.Printf("%s git remote add origin repository success", location)
	output, err = util.CreatePreCMD(fmt.Sprintf("git pull  origin   %s", GlobalParametes.GetBranchName())).Output()
	if err != nil {
		return fmt.Errorf("git pull remote repository error:%s", output)
	}
	log.Printf("%s pull remote repository success", location)
	return nil
}
