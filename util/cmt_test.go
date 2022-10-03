package util

import (
	"fmt"
	"testing"
)

func TestCmd(t *testing.T) {
	println(fmt.Sprintf("docker container ls -a --format \"{{json .}}\""))
}
