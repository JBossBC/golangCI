package main

import (
	"golangCI/util"
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
	prepare()

}
func prepare() {
	_, err := util.GetSystemInformation()
	if err != nil {
		panic(err.Error())
	}
	err = util.ConvertSystemToCommand()
	if err != nil {
		panic(err.Error())
	}
	_, err = util.InspectGo()
	if err != nil {
		panic(err.Error())
	}
	_, err = util.InspectGit()
	if err != nil {
		panic(err.Error())
	}
}
