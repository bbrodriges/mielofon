package dialog

import (
	"encoding/json"
)

type Capability int

const (
	ScreenCapability Capability = iota + 1
	PaymentsCapability
	AccountLinkingCapability
	AudioPlayerCapability
)

type Capabilities struct {
	Screen         json.RawMessage `json:"screen,omitempty"`
	Payments       json.RawMessage `json:"payments,omitempty"`
	AccountLinking json.RawMessage `json:"account_linking,omitempty"`
	AudioPlayer    json.RawMessage `json:"audio_player,omitempty"`
}

// HasCapability checks device capability
func (i Input) HasCapability(c Capability) bool {
	switch c {
	case ScreenCapability:
		return i.Meta.Interfaces.Screen != nil
	case PaymentsCapability:
		return i.Meta.Interfaces.Payments != nil
	case AccountLinkingCapability:
		return i.Meta.Interfaces.AccountLinking != nil
	case AudioPlayerCapability:
		return i.Meta.Interfaces.AudioPlayer != nil
	default:
		return false
	}
}
