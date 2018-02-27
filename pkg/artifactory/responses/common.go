package responses

import (
	"errors"
	"strings"
	"time"
)

// JSONTime is for custom marshal/unmarshal of artifactory datetime values
type JSONTime struct {
	time.Time
}

// ISO8601Time is a custom/marshal/unmarshal of iso8601 time values
type ISO8601Time struct {
	time.Time
}

const artifTime = "2006-01-02T15:04:05Z"
const iso8601Time = "2006-01-02T15:04:05.000-0700"

// UnmarshalJSON parses the rundeck datetime format
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("JSONTime: UnmarshalText on nil pointer")
	}
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	tempTime, tErr := time.Parse(artifTime, s)
	if tErr != nil {
		return errors.New("JSONTime: " + tErr.Error())
	}
	t.Time = tempTime
	return nil
}

// UnmarshalJSON parses the iso8601 datetime format
func (t *ISO8601Time) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("ISO8601Time: UnmarshalText on nil pointer")
	}
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	tempTime, tErr := time.Parse(iso8601Time, s)
	if tErr != nil {
		return errors.New("ISO8601Time: " + tErr.Error())
	}
	t.Time = tempTime
	return nil
}
