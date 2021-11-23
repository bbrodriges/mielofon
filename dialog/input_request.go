package dialog

import (
	"encoding/json"
	"fmt"
)

type RequestType string

const (
	TypeSimpleUtterance RequestType = "SimpleUtterance"
	TypeButtonPressed   RequestType = "ButtonPressed"

	TypeAudioPlayerPlaybackStarted        RequestType = "AudioPlayer.PlaybackStarted"
	TypeAudioPlayerPlaybackFinished       RequestType = "AudioPlayer.PlaybackFinished"
	TypeAudioPlayerPlaybackNearlyFinished RequestType = "AudioPlayer.PlaybackNearlyFinished"
	TypeAudioPlayerPlaybackStopped        RequestType = "AudioPlayer.PlaybackStopped"
	TypeAudioPlayerPlaybackFailed         RequestType = "AudioPlayer.PlaybackFailed"

	TypePurchaseConfirmation RequestType = "Purchase.Confirmation"

	TypeShowPull RequestType = "Show.Pull"
)

// Request represents abstract request field of input
type Request interface {
	// Type returns type of input request
	Type() RequestType
}

type RequestMarkup struct {
	DangerousContext bool `json:"dangerous_context,omitempty"`
}

// unmarshalRequest unmarshals request object from JSON
func unmarshalRequest(b json.RawMessage) (Request, error) {
	var raw = struct {
		Type RequestType `json:"type"`
	}{}

	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, fmt.Errorf("cannot unmarshal to intermediate struct: %w", err)
	}

	switch raw.Type {
	case TypeSimpleUtterance:
		var req SimpleUtteranceRequest
		if err := json.Unmarshal(b, &req); err != nil {
			return nil, err
		}
		return req, nil
	case TypeButtonPressed:
		var req ButtonPressedRequest
		if err := json.Unmarshal(b, &req); err != nil {
			return nil, err
		}
		return req, nil
	case TypeShowPull:
		var req ShowPullRequest
		if err := json.Unmarshal(b, &req); err != nil {
			return nil, err
		}
		return req, nil
	case TypePurchaseConfirmation:
		var req PurchaseConfirmationRequest
		if err := json.Unmarshal(b, &req); err != nil {
			return nil, err
		}
		return req, nil
	case TypeAudioPlayerPlaybackStarted,
		TypeAudioPlayerPlaybackStopped,
		TypeAudioPlayerPlaybackNearlyFinished,
		TypeAudioPlayerPlaybackFinished:
		var req AudioPlayerRequest
		if err := json.Unmarshal(b, &req); err != nil {
			return nil, err
		}
		return req, nil
	case TypeAudioPlayerPlaybackFailed:
		var req AudioPlayerErrorRequest
		if err := json.Unmarshal(b, &req); err != nil {
			return nil, err
		}
		return req, nil
	default:
		return nil, fmt.Errorf("unsupported type: %s", raw.Type)
	}
}
