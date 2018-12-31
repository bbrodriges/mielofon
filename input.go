package mielofon

import (
	"time"
)

type Input struct {
	Meta    Meta    `json:"meta"`
	Request Request `json:"request"`
	Session Session `json:"session"`
	Version string  `json:"version"`
}

// HasScreen indicates device screen capability
func (i Input) HasScreen() bool {
	return i.Meta.Interfaces.Screen != nil
}

// DeviceLocation returns location object based on request location info.
// It will return current time.Local on location parse error
func (i Input) DeviceLocation() *time.Location {
	loc, err := time.LoadLocation(i.Meta.Timezone)
	if err != nil {
		return time.Local
	}
	return loc
}

// DeviceTime returns time instance based on request location info.
// It will return current local time on location parse error
func (i Input) DeviceTime() time.Time {
	return time.Now().In(i.DeviceLocation())
}

// Initiated indicates initial dialog request
func (i Input) Initiated() bool {
	return i.Session.New
}

type Meta struct {
	Locale     string     `json:"locale"`
	Timezone   string     `json:"timezone"`
	ClientID   string     `json:"client_id"`
	Interfaces Interfaces `json:"interfaces"`
}

type Interfaces struct {
	Screen interface{} `json:"screen,omitempty"`
}

type Session struct {
	New       bool   `json:"new"`
	MessageID int    `json:"message_id"`
	SessionID string `json:"session_id"`
	SkillID   string `json:"skill_id"`
	UserID    string `json:"user_id"`
}
