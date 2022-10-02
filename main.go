package main

import (
	"golangCI/deploy"
	"golangCI/pipeline"
	"golangCI/prepare"
	"log"
)

type Operation_System int

const (
	LINUX Operation_System = iota
	WINDOWS
	Default
)

var CurrentSystem Operation_System = Default

var GlobalCommand string

func main() {
	err := pipeline.Analy()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	prepareOperation()
	deploy.ContinuousDeploy()

}
func prepareOperation() {
	_, err := prepare.GetSystemInformation()
	if err != nil {
		panic(err.Error())
	}
	err = prepare.ConvertSystemToCommand()
	if err != nil {
		panic(err.Error())
	}
	_, err = prepare.InspectGo()
	if err != nil {
		panic(err.Error())
	}
	_, err = prepare.InspectGit()
	if err != nil {
		panic(err.Error())
	}
	_, err = prepare.InspectDocker()
	if err != nil {
		panic(err.Error())
	}
}
