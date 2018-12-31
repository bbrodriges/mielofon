package mielofon

import (
	"encoding/json"
)

type Output struct {
	Response Response `json:"response"`
	Session  Session  `json:"session"`
	Version  string   `json:"version"`
}

type Response struct {
	Text       string     `json:"text"`
	Tts        string     `json:"tts"`
	Card       OutputCard `json:"card"`
	Buttons    []Button   `json:"buttons"`
	EndSession bool       `json:"end_session"`
}

type Button struct {
	Title   string          `json:"title"`
	Payload json.RawMessage `json:"payload"`
	URL     string          `json:"url"`
	Hide    bool            `json:"hide"`
}
