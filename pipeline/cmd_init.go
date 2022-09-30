package pipeline

import "os/exec"

func CreateWorkCMD(workDir string, operationInstruction string, arg string) *exec.Cmd {
	preDir = workDir
	preOperationInstruction = operationInstruction
	command := exec.Command(operationInstruction, arg)
	command.Dir = workDir
	return command
}

var (
	preDir                  string
	preOperationInstruction string
)

func CreatePreCMD(arg string) *exec.Cmd {
	return CreateWorkCMD(preDir, preOperationInstruction, arg)
}
