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

// GetContainerGroupList fetches a list of tokens
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

// // CreateToken registers a new token. Returns a list of tokens on success
// func CreateToken(key string, description string, restricted bool, billingAccountID int) ([]Token, error) {
// 	form := url.Values{}
// 	form.Add("billing_account_id", fmt.Sprintf("%d", billingAccountID))
// 	form.Add("description", description)
// 	form.Add("restricted", fmt.Sprintf("%v", restricted))

// 	resp, err := postForm(key, "user-resource/token", form)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tokenList, err := parseContainerGroupList([]byte(resp))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tokenList, nil
// }

// // DeleteToken deletes a token by its ID
// func DeleteToken(key string, tokenID int) error {
// 	form := url.Values{}
// 	form.Add("token_id", fmt.Sprintf("%d", tokenID))

// 	_, err := deleteForm(key, "user-resource/token", form)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateToken updates a token with the values supplied
// func UpdateToken(key string, tokenID int, changedFields url.Values) error {
// 	changedFields.Set("token_id", fmt.Sprintf("%d", tokenID))

// 	_, err := patchForm(key, "user-resource/token", changedFields)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
