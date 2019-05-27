package api

import (
	"encoding/json"
	"strings"
	"time"
)

type PilwTime struct {
	time.Time
}

func (pt *PilwTime) UnmarshalJSON(p []byte) error {
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(p), "\""))

	if err != nil {
		return err
	}

	pt.Time = t

	return nil
}

type UserInfo struct {
	CookieID     string   `json:"cookie_id"`
	ID           int      `json:"id"`
	LastActivity PilwTime `json:"last_activity"`
	Name         string   `json:"name"`
}

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
