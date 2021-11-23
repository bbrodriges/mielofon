package dialog

import (
	"encoding/json"
	"fmt"
)

var _ Request = (*AudioPlayerRequest)(nil)
var _ Request = (*AudioPlayerErrorRequest)(nil)

const (
	MediaErrorUnknown            = "MEDIA_ERROR_UNKNOWN"
	MediaErrorServiceUnavailable = "MEDIA_ERROR_SERVICE_UNAVAILABLE"
)

// AudioPlayerRequest represents request of type `AudioPlayer`.
// It is used to represent any non error request
type AudioPlayerRequest struct {
	requestType RequestType
}

func (a AudioPlayerRequest) Type() RequestType {
	return a.requestType
}

func (a *AudioPlayerRequest) UnmarshalJSON(b []byte) error {
	var raw = struct {
		Type RequestType `json:"type"`
	}{}

	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("cannot unmarshal request field: %w", err)
	}

	a.requestType = raw.Type
	return nil
}

type AudioPlayerErrorRequest struct {
	Error AudioPlayerErrorData `json:"error"`
}

type AudioPlayerErrorData struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (a AudioPlayerErrorRequest) Type() RequestType {
	return TypeAudioPlayerPlaybackFailed
}
