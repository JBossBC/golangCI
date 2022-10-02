package pipeline

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonFile struct {
	Parameters Parameters `json:parameters`
}

func Analy() error {
	file, err := os.OpenFile("./pipeline.json", os.O_APPEND, 0755)
	if err != nil {
		return fmt.Errorf("open pipeline.json error:%s", err.Error())
	}
	var analyObject jsonFile
	err = json.NewDecoder(file).Decode(&analyObject)
	if err != nil {
		return fmt.Errorf("analy json file errror:%s", err.Error())
	}
	GlobalParametes = analyObject.Parameters
	return nil
}
