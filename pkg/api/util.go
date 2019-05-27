package api

import (
	"strings"
	"time"
)

type PilwTime struct {
	time.Time
}

func (pt *PilwTime) UnmarshalJSON(p []byte) error {
	str := strings.Trim(string(p), "\"")
	if str == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02 15:04:05", str)

	if err != nil {
		return err
	}

	pt.Time = t

	return nil
}
