package api

import (
	"encoding/json"
)

// UserInfo represents a Pilw user's metadata
type UserInfo struct {
	CookieID     string   `json:"cookie_id"`
	ID           int      `json:"id"`
	LastActivity PilwTime `json:"last_activity"`
	Name         string   `json:"name"`
}

// GetUserInfo returns the current authenticated user's metadata
func GetUserInfo(key string) (UserInfo, error) {
	var userInfo UserInfo

	resp, err := get(key, "user-resource/user")
	if err != nil {
		return userInfo, err
	}

	err = json.Unmarshal([]byte(resp), &userInfo)
	if err != nil {
		return userInfo, err
	}

	return userInfo, err
}
