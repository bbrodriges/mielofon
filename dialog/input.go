package dialog

import (
	"encoding/json"
	"time"
)

type Input struct {
	Meta    Meta    `json:"meta,omitempty"`
	Request Request `json:"request,omitempty"`
	Session Session `json:"session,omitempty"`
	Version string  `json:"version,omitempty"`
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

// DialogInitiated indicates initial dialog request
func (i Input) DialogInitiated() bool {
	return i.Session.New
}

// HasEntities method checks whether entities of certain
// type exists or not in dialog request
func (i Input) HasEntities(t EntityType) bool {
	for _, e := range i.Request.Nlu.Entities {
		if e.Type == t {
			return true
		}
	}
	return false
}

// Entities is a convenient method to get all YANDEX.<TYPE>
// entities from input request. It is safe to make type assertion
// without validation
func (i Input) Entities(t EntityType) []Entity {
	res := make([]Entity, 0, len(i.Request.Nlu.Entities))
	for _, e := range i.Request.Nlu.Entities {
		if e.Type == t {
			res = append(res, e)
		}
	}
	return res
}

// HasKeyword checks if given string presents in
// input command tokens
func (i Input) HasKeyword(k string) bool {
	for _, t := range i.Request.Nlu.Tokens {
		if k == t {
			return true
		}
	}
	return false
}

// HasKeywords returns an intersection between input command
// tokens and given keywords
func (i Input) HasKeywords(ks ...string) []string {
	if len(ks) == 0 || len(i.Request.Nlu.Tokens) == 0 {
		return nil
	}

	p, s := i.Request.Nlu.Tokens, ks
	if len(ks) > len(i.Request.Nlu.Tokens) {
		p, s = ks, i.Request.Nlu.Tokens
	}

	m := make(map[string]struct{})
	for _, i := range p {
		m[i] = struct{}{}
	}

	res := make([]string, 0, len(m))
	for _, v := range s {
		if _, exists := m[v]; exists {
			res = append(res, v)
		}
	}

	return res
}

type Meta struct {
	Locale     string     `json:"locale,omitempty"`
	Timezone   string     `json:"timezone,omitempty"`
	ClientID   string     `json:"client_id,omitempty"`
	Interfaces Interfaces `json:"interfaces,omitempty"`
}

type Interfaces struct {
	Screen json.RawMessage `json:"screen,omitempty"`
}

type Session struct {
	New       bool   `json:"new,omitempty"`
	MessageID int    `json:"message_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	SkillID   string `json:"skill_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}
