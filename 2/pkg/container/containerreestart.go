package container

import (
	"net/http"
)

func ContainerRestart(serverUrl string, containerID string) error {
	_, err := http.Post(serverUrl+"/containers/"+containerID+"/restart", "application/json", nil)
	if err != nil {
		return err
	}
	return nil
}
