package dialog

const (
	AudioPlayerActionPlay = "play"
	AudioPlayerActionStop = "stop"
)

// AudioPlayerDirective represents directive of type `audio_player`
type AudioPlayerDirective struct {
	Action string                   `json:"action"`
	Item   AudioPlayerDirectiveItem `json:"item,omitempty"`
}

func (AudioPlayerDirective) Type() DirectiveType {
	return DirectiveAudioPlayer
}

type AudioPlayerDirectiveItem struct {
	Stream   AudioPlayerDirectiveStream   `json:"stream"`
	Metadata AudioPlayerDirectiveMetadata `json:"metadata"`
}

type AudioPlayerDirectiveStream struct {
	URL      string `json:"url"`
	OffsetMs int64  `json:"offset_ms"`
	Token    string `json:"token"`
}

type AudioPlayerDirectiveMetadata struct {
	Title           string                    `json:"title"`
	SubTitle        string                    `json:"sub_title"`
	Art             AudioPlayerDirectiveImage `json:"art"`
	BackgroundImage AudioPlayerDirectiveImage `json:"background_image"`
}

type AudioPlayerDirectiveImage struct {
	URL string `json:"url"`
}
