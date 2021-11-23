package dialog

import (
	"encoding/json"
	"fmt"
)

// Input represents request payload
type Input struct {
	Meta    Meta    `json:"meta,omitempty"`
	Request Request `json:"request,omitempty"`
	Session Session `json:"session,omitempty"`
	State   State   `json:"state,omitempty"`
	Version string  `json:"version,omitempty"`
}

type Meta struct {
	Locale     string       `json:"locale,omitempty"`
	Timezone   string       `json:"timezone,omitempty"`
	ClientID   string       `json:"client_id,omitempty"`
	Interfaces Capabilities `json:"interfaces,omitempty"`
}

type State struct {
	Session     StateValue `json:"session"`
	User        StateValue `json:"user"`
	Application StateValue `json:"application"`
}

type StateValue struct {
	Value int64 `json:"value"`
}

func (i *Input) UnmarshalJSON(b []byte) error {
	var raw = struct {
		Meta    Meta            `json:"meta,omitempty"`
		Request json.RawMessage `json:"request,omitempty"`
		Session Session         `json:"session,omitempty"`
		State   State           `json:"state,omitempty"`
		Version string          `json:"version,omitempty"`
	}{}

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	req, err := unmarshalRequest(raw.Request)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request field: %w", err)
	}

	i.Meta = raw.Meta
	i.Request = req
	i.Session = raw.Session
	i.State = raw.State
	i.Version = raw.Version
	return nil
}
