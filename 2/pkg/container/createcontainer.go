package container

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

var (
	ErrContainerExists = errors.New("Container with specified name already exists")
)

// helper struct used for json encoding config
type configBodyWrapper struct {
	*container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
}

// function used to create container if it does not exists
func CreateContainer(serverUrl string, config types.ContainerCreateConfig, containerName string) error {

	query := url.Values{}
	if containerName != "" {
		containerId, err := GetContainerId(serverUrl, containerName)

		if err != nil {
			return err
		}

		if containerId != "" {
			return ErrContainerExists
		}

		query.Set("name", containerName)
	}

	configBody := configBodyWrapper{
		Config:           config.Config,
		HostConfig:       config.HostConfig,
		NetworkingConfig: config.NetworkingConfig,
	}

	encodedConfig, _ := json.Marshal(configBody)

	// after this container is in state Created
	_, err := http.Post(serverUrl+"/containers/create?"+query.Encode(), "application/json", bytes.NewBuffer(encodedConfig))
	if err != nil {
		return err
	}

	// start container
	_, err = http.Post(serverUrl+"/containers/"+containerName+"/start", "application/json", nil)
	if err != nil {
		return err
	}

	// start container
	// bodyBytes, _ := io.ReadAll(resp.Body)
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	return nil
}
