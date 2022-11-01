package container

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/docker/docker/api/types"
)

// Function used to gett container ID based on name
func GetContainerId(serverUrl string, containerName string) (string, error) {

	// create query string with filter for name
	// it must looks like
	// filters = {"name": {"<container_name>": true}}
	queryFilter := url.Values{}
	nameFilter := make(map[string]map[string]bool)
	nameFilter["name"] = map[string]bool{containerName: true}

	encodedFilter, _ := json.Marshal(nameFilter)
	queryFilter.Set("filters", string(encodedFilter))

	// do GET request
	resp, err := http.Get(serverUrl + "/containers/json?" + queryFilter.Encode())

	// check error
	if err != nil {
		return "", err
	}

	// decode response body an return
	var containers []types.Container

	err = json.NewDecoder(resp.Body).Decode(&containers)
	if err != nil {
		return "", err
	}

	if len(containers) == 0 {
		log.Printf("Not found container with name = %s\n", containerName)
		return "", nil
	}

	fmt.Printf("Found containerId = %s\n", containers[0].ID)
	return containers[0].ID, nil
}
