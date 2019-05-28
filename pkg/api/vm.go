package api

import (
	"encoding/json"
)

type VMStorage struct {
	CreatedAt  PilwTime `json:"created_at"`
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Pool       string   `json:"pool"`
	Primary    bool     `json:"primary"`
	PublicIPv4 string   `json:"public_ipv4"`
	// replica
	Shared    bool     `json:"shared"`
	Size      int      `json:"size"`
	Type      string   `json:"type"`
	UpdatedAt PilwTime `json:"updated_at"`
	UserID    int      `json:"user_id"`
	UUID      string   `json:"uuid"`
}

type VM struct {
	Backup           bool        `json:"backup"`
	BillingAccountID int         `json:"billing_account"`
	CreatedAt        PilwTime    `json:"created_at"`
	Description      string      `json:"description"`
	Hostname         string      `json:"hostname"`
	ID               int         `json:"id"`
	Mac              string      `json:"mac"`
	Memory           int         `json:"memory"`
	Name             string      `json:"name"`
	OSName           string      `json:"os_name"`
	OSVersion        string      `json:"os_version"`
	PrivateIPv4      string      `json:"private_ipv4"`
	PublicIPv4       string      `json:"public_ipv4"`
	Status           string      `json:"status"`
	Storage          []VMStorage `json:"storage"`
	UpdatedAt        PilwTime    `json:"updated_at"`
	UserID           int         `json:"user_id"`
	Username         string      `json:"username"`
	UUID             string      `json:"uuid"`
	VCPU             int         `json:"vcpu"`
}

func ParseVMList(str []byte) ([]VM, error) {
	var vmList []VM

	err := json.Unmarshal(str, &vmList)
	if err != nil {
		return vmList, err
	}

	return vmList, nil
}

func GetVMList(key string) ([]VM, error) {
	resp, err := get(key, "user-resource/vm/list")
	if err != nil {
		return nil, err
	}

	vmList, err := ParseVMList([]byte(resp))
	if err != nil {
		return nil, err
	}

	return vmList, err
}
