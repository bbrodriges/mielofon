package dialog

import (
	"encoding/json"
)

type Output struct {
	Response Response `json:"response,omitempty"`
	Session  Session  `json:"session,omitempty"`
	Version  string   `json:"version,omitempty"`
}

type Response struct {
	Text       string     `json:"text,omitempty"`
	Tts        string     `json:"tts,omitempty"`
	Card       OutputCard `json:"card,omitempty"`
	Buttons    []Button   `json:"buttons,omitempty"`
	EndSession bool       `json:"end_session,omitempty"`
}

type Button struct {
	Title   string          `json:"title,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
	URL     string          `json:"url,omitempty"`
	Hide    bool            `json:"hide,omitempty"`
}
