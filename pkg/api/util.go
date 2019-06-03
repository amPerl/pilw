package api

import (
	"strings"
	"time"
)

// PilwTime is a wrapper over Go's native time only for unmarshaling from JSON
type PilwTime struct {
	time.Time
}

// UnmarshalJSON specifies formatting to unmarshal Pilw timestamps
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
