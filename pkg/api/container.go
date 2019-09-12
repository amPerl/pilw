package api

import (
	"encoding/json"
)

// ContainerServiceVolume represents an individual Pilw container service volume definition
type ContainerServiceVolume struct {
	ContainerPath string   `json:"container_path"`
	CreatedAt     PilwTime `json:"created_at"`
}

// ContainerService represents an individual Pilw container service
type ContainerService struct {
	CPU             float32                  `json:"cpu"`
	CreatedAt       PilwTime                 `json:"created_at"`
	Environment     map[string]string        `json:"env"`
	GroupSUUID      string                   `json:"group_suuid"`
	HTTPExposed     bool                     `json:"http_exposed"`
	Image           string                   `json:"image"`
	Instances       int                      `json:"instances"`
	InternalHost    string                   `json:"internal_host"`
	IsPublic        bool                     `json:"is_public"`
	Name            string                   `json:"name"`
	RAM             int                      `json:"ram"`
	Status          string                   `json:"status"`
	StatusChangedAt PilwTime                 `json:"status_changed_at"`
	SUUID           string                   `json:"suuid"`
	UserID          int                      `json:"user_id"`
	Volumes         []ContainerServiceVolume `json:"volumes"`
}

// ContainerGroup represents a Pilw container group (of services)
type ContainerGroup struct {
	CreatedAt PilwTime           `json:"created_at"`
	Name      string             `json:"name"`
	Services  []ContainerService `json:"services"`
	SUUID     string             `json:"suuid"`
	UserID    int                `json:"user_id"`
}

func parseContainerGroupList(str []byte) ([]ContainerGroup, error) {
	var containerGroupList []ContainerGroup

	err := json.Unmarshal(str, &containerGroupList)
	if err != nil {
		return containerGroupList, err
	}

	return containerGroupList, nil
}

// GetContainerGroupList fetches a list of container groups
func GetContainerGroupList(key string) ([]ContainerGroup, error) {
	resp, err := get(key, "container/groups")
	if err != nil {
		return nil, err
	}

	containerGroupList, err := parseContainerGroupList([]byte(resp))
	if err != nil {
		return nil, err
	}

	return containerGroupList, err
}

// GetContainerServiceList fetches a list of container groups
func GetContainerServiceList(key string) ([]ContainerService, error) {
	containerGroupList, err := GetContainerGroupList(key)
	if err != nil {
		return nil, err
	}

	containerServiceList := make([]ContainerService, 0)
	for _, containerGroup := range containerGroupList {
		for _, containerService := range containerGroup.Services {
			containerServiceList = append(containerServiceList, containerService)
		}
	}

	return containerServiceList, err
}

// StopContainerService stops a container service by suuid
func StopContainerService(key string, suuid string) error {
	_, err := put(key, "container/services/"+suuid+"/stop")
	if err != nil {
		return err
	}

	return err
}

// StartContainerService starts a container service by suuid
func StartContainerService(key string, suuid string) error {
	_, err := put(key, "container/services/"+suuid+"/start")
	if err != nil {
		return err
	}

	return err
}
