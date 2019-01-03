package dialog

import (
	"encoding/json"
)

type RequestType string

const (
	SimpleUtterance RequestType = "SimpleUtterance"
	ButtonPressed   RequestType = "ButtonPressed"
)

type Request struct {
	Command           string          `json:"command,omitempty"`
	OriginalUtterance string          `json:"original_utterance,omitempty"`
	Type              RequestType     `json:"type,omitempty"`
	Markup            Markup          `json:"markup,omitempty"`
	Payload           json.RawMessage `json:"payload,omitempty"`
	Nlu               Nlu             `json:"nlu,omitempty"`
}

type Markup struct {
	DangerousContext bool `json:"dangerous_context,omitempty"`
}

type Nlu struct {
	Tokens   []string `json:"tokens,omitempty"`
	Entities []Entity `json:"entities,omitempty"`
}
