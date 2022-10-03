package deploy

import (
	"encoding/json"
	"fmt"
	"golangCI/pipeline"
	"golangCI/prepare"
	"golangCI/util"
	"log"
	"strings"
	"time"
)

var preImages = make(map[string]int, 100)

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
	preImages[imageHash]++
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
		containerRm()

		time.Sleep(time.Second * 60)
	}
}

var containerSet []*ContainerInfo

type ContainerInfo struct {
	Command   string `json:Command`
	CreatedAt string `json:CreatedAt`
	ID        string `json:ID`
	Image     string `json:Image`
	Names     string `json:Names`
	Ports     string `json:Ports`
	Status    string `json:Status`
	State     string `json:State`
	Networks  string `json:Network`
}

func containerRm() {
	repository := pipeline.GlobalParametes.GitLocationRepo
	var lsContainers = fmt.Sprintf(`docker container ls -a --format "{{json .}}"`)
	containers, err := util.CreateWorkCMD(repository, prepare.GlobalCommand, lsContainers).Output()
	if err != nil {
		log.Fatalf("docker container ls error:%s", lsContainers)
		return
	}
	err = json.Unmarshal(containers, &containerSet)
	if err != nil {
		log.Fatalf("json analy error:%s", err.Error())
		return
	}
	for i := 0; i < len(containerSet); i++ {
		var temp = containerSet[i]
		if data, ok := preImages[temp.Image]; ok && data != 0 {
			if temp.State != "running" {
				var rmContaienr = fmt.Sprintf("docker container rm %s", temp.ID)
				_, err := util.CreatePreCMD(rmContaienr).Output()
				if err != nil {
					log.Printf("remove docker container %s error:%s", temp.ID, err.Error())
					return
				}
				preImages[temp.Image]--
				if preImages[temp.Image] <= 0 {
					var rmImage = fmt.Sprintf("docker image rm %s", temp.Image)
					_, err := util.CreatePreCMD(rmImage).Output()
					if err != nil {
						log.Printf("remove docker image %s error:%s", temp.Image, err.Error())
						return
					}
				}
			}
		}
	}
}
