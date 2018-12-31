package mielofon

import (
	"encoding/json"
)

type RequestType string

const (
	SimpleUtterance RequestType = "SimpleUtterance"
	ButtonPressed   RequestType = "ButtonPressed"
)

type Request struct {
	Command           string          `json:"command"`
	OriginalUtterance string          `json:"original_utterance"`
	Type              RequestType     `json:"type"`
	Markup            Markup          `json:"markup"`
	Payload           json.RawMessage `json:"payload"`
	Nlu               Nlu             `json:"nlu"`
}

type Markup struct {
	DangerousContext bool `json:"dangerous_context"`
}

type Nlu struct {
	Tokens   []string `json:"tokens"`
	Entities []Entity `json:"entities"`
}
