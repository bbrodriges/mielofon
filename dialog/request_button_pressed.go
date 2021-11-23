package dialog

import (
	"encoding/json"
)

var _ Request = (*ButtonPressedRequest)(nil)

// ButtonPressedRequest represents request of type `ButtonPressed`
type ButtonPressedRequest struct {
	Markup  *RequestMarkup          `json:"markup,omitempty"`
	Nlu     ButtonPressedRequestNlu `json:"nlu"`
	Payload json.RawMessage         `json:"payload,omitempty"`
}

type ButtonPressedRequestNlu struct {
	Tokens   []string        `json:"tokens,omitempty"`
	Entities []NamedEntity   `json:"entities,omitempty"`
	Intents  json.RawMessage `json:"intents,omitempty"`
}

func (s ButtonPressedRequest) Type() RequestType {
	return TypeButtonPressed
}
