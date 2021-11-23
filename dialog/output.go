package dialog

import (
	"encoding/json"
	"time"
)

type Output struct {
	Response         Response        `json:"response,omitempty"`
	SessionState     StateValue      `json:"session_state,omitempty"`
	UserStateUpdate  StateValue      `json:"user_state_update,omitempty"`
	ApplicationState StateValue      `json:"application_state,omitempty"`
	Analytics        OutputAnalytics `json:"analytics,omitempty"`
	Version          string          `json:"version,omitempty"`
}

type Response struct {
	Text         string       `json:"text,omitempty"`
	TTS          string       `json:"tts,omitempty"`
	Card         OutputCard   `json:"card,omitempty"`
	Buttons      []Button     `json:"buttons,omitempty"`
	EndSession   bool         `json:"end_session,omitempty"`
	Directives   Directives   `json:"directives,omitempty"`
	ShowItemMeta ShowItemMeta `json:"show_item_meta,omitempty"`
}

type Button struct {
	Title   string          `json:"title,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
	URL     string          `json:"url,omitempty"`
	Hide    bool            `json:"hide,omitempty"`
}

type ShowItemMeta struct {
	ContentID       string    `json:"content_id"`
	Title           string    `json:"title"`
	TitleTTS        string    `json:"title_tts"`
	PublicationDate time.Time `json:"publication_date"`
	ExpirationDate  time.Time `json:"expiration_date"`
}

type OutputAnalytics struct {
	Events []AnalyticsEvent `json:"events"`
}

type AnalyticsEvent struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}
