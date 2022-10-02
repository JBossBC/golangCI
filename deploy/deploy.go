package deploy

import (
	"fmt"
	"golangCI/pipeline"
	"golangCI/prepare"
	"golangCI/util"
	"log"
	"strings"
	"time"
)

func Deploy(contextPath string) error {
	log.Println("begin deploy...")
	var aimName = "project"
	_, err := util.CreateWorkCMD(contextPath, prepare.GlobalCommand, fmt.Sprintf("docker build -f %s -t %s .", pipeline.GlobalParametes.DockerfileName, aimName)).Output()
	if err != nil {
		return fmt.Errorf("docker build failed:%s", err.Error())
	}
	images, err := util.CreatePreCMD("docker image ls -a").Output()
	info := strings.Fields(string(images))
	var locationInfo = -1
	for i := 0; i < len(info); i++ {
		if info[i] == aimName {
			locationInfo = i + 2
			break
		}
	}
	if locationInfo == -1 {
		return fmt.Errorf("cant find the image %s", images)
	}
	var imageHash = info[locationInfo]
	var globalErr error
	go func() {
		_, err = util.CreatePreCMD(fmt.Sprintf("docker run -p 8081:8081  %s", imageHash)).Output()
		if err != nil {
			globalErr = fmt.Errorf("docker run error:%s", err.Error())
		}
	}()
	time.Sleep(15 * time.Second)
	log.Println("deploy success...")
	return nil
}

func ContinuousDeploy() {
	log.Println("begin continuous deploy...")
	for {
		repository, err := pipeline.GlobalParametes.GetLocationRepository()
		if err != nil {
			log.Fatal(err.Error())
			time.Sleep(time.Second * 60)
			continue
		}
		err = Deploy(repository)
		if err != nil {
			panic(err.Error())
		}

		time.Sleep(time.Second * 60)
	}
}
