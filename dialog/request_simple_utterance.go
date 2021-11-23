package dialog

import (
	"encoding/json"
)

var _ Request = (*SimpleUtteranceRequest)(nil)

// SimpleUtteranceRequest represents request of type `SimpleUtterance`
type SimpleUtteranceRequest struct {
	Command           string             `json:"command,omitempty"`
	OriginalUtterance string             `json:"original_utterance,omitempty"`
	Markup            *RequestMarkup     `json:"markup,omitempty"`
	Nlu               SimpleUtteranceNlu `json:"nlu,omitempty"`
	Payload           json.RawMessage    `json:"payload,omitempty"`
}

func (s SimpleUtteranceRequest) Type() RequestType {
	return TypeSimpleUtterance
}

type SimpleUtteranceNlu struct {
	Tokens   []string        `json:"tokens,omitempty"`
	Entities []NamedEntity   `json:"entities,omitempty"`
	Intents  json.RawMessage `json:"intents,omitempty"`
}
