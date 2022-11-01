package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/docker/docker/api/types"
	cntr "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/piotmni/rekrutacja5g/pkg/container"
	"github.com/robfig/cron/v3"
)

func main() {

	// if missing argument end this.
	if len(os.Args) <= 1 {
		log.Printf("Error: missing argument. Give ip address of server as a argument.\n ex. %s 192.168.0.100", os.Args[0])
		return
	}

	// building sever url based on argument
	serverUrl := fmt.Sprintf("http://%s:2375/v1.41", os.Args[1])
	log.Printf("Using server %s", serverUrl)
	c := cron.New()

	// solution .1
	c.AddFunc("@every 1h", func() {
		log.Println("Starting cron: restarting container test1")
		err := container.ContainerRestart(serverUrl, "test1")
		if err != nil {
			log.Printf("Error while restarting container: %s", err.Error())
		}
	})

	// solution .2
	c.AddFunc("@every 30m", func() {
		log.Println("Starting cron: create container test2 if it not exist")
		// prepare config for container
		var containerConfig types.ContainerCreateConfig

		containerConfig.Config = &cntr.Config{Image: "demo/dummy:latest"}
		containerConfig.NetworkingConfig = &network.NetworkingConfig{}
		containerConfig.HostConfig = &cntr.HostConfig{
			Binds: []string{"/test2:/test2"},
		}

		// call create container
		err := container.CreateContainer(serverUrl, containerConfig, "test2")

		if err != nil {
			log.Printf("Error while creating container: %s", err.Error())
		}
	})

	// Start cron and end when signal
	go c.Start()
	singalChannel := make(chan os.Signal)
	signal.Notify(singalChannel, os.Interrupt, os.Kill)
	<-singalChannel

}
