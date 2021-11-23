package dialog

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInput_SimpleUtterance(t *testing.T) {
	expected := Input{
		Meta: Meta{
			Locale:   "ru-RU",
			Timezone: "Europe/Moscow",
			ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
			Interfaces: Capabilities{
				Screen:         json.RawMessage(`{}`),
				Payments:       nil,
				AccountLinking: json.RawMessage(`{}`),
				AudioPlayer:    nil,
			},
		},
		Request: SimpleUtteranceRequest{
			Command:           "закажи пиццу на улицу льва толстого 16 на завтра",
			OriginalUtterance: "закажи пиццу на улицу льва толстого, 16 на завтра",
			Markup: &RequestMarkup{
				DangerousContext: true,
			},
			Nlu: SimpleUtteranceNlu{
				Tokens: []string{
					"закажи",
					"пиццу",
					"на",
					"льва",
					"толстого",
					"16",
					"на",
					"завтра",
				},
				Entities: []NamedEntity{
					{
						Tokens: TokensRange{Start: 2, End: 6},
						Type:   TypeYandexGeo,
						Value: YandexGeo{
							Country:     "",
							City:        "",
							Street:      "льва толстого",
							HouseNumber: "16",
							Airport:     "",
						},
					},
					{
						Tokens: TokensRange{Start: 3, End: 5},
						Type:   TypeYandexFIO,
						Value: YandexFIO{
							FirstName: "лев",
							LastName:  "толстой",
						},
					},
					{
						Tokens: TokensRange{Start: 5, End: 6},
						Type:   TypeYandexNumber,
						Value:  YandexNumber(16),
					},
					{
						Tokens: TokensRange{Start: 6, End: 8},
						Type:   TypeYandexDatetime,
						Value:  YandexDatetime{time.Now().AddDate(0, 0, 1).Truncate(time.Minute)},
					},
				},
			},
			Payload: json.RawMessage(`{}`),
		},
		Session: Session{
			MessageID: 0,
			SessionID: "2eac4854-fce721f3-b845abba-20d60",
			SkillID:   "3ad36498-f5rd-4079-a14b-788652932056",
			UserID:    "47C73714B580ED2469056E71081159529FFC676A4E5B059D629A819E857DC2F8",
			User: SessionUser{
				UserID:      "6C91DA5198D1758C6A9F63A7C5CDDF09359F683B13A18A151FBF4C8B092BB0C2",
				AccessToken: "AgAAAAAB4vpbAAApoR1oaCd5yR6eiXSHqOGT8dT",
			},
			Application: SessionApplication{
				ApplicationID: "47C73714B580ED2469056E71081159529FFC676A4E5B059D629A819E857DC2F8",
			},
			New: true,
		},
		State: State{
			Session:     StateValue{Value: 10},
			User:        StateValue{Value: 42},
			Application: StateValue{Value: 37},
		},
		Version: "1.0",
	}

	fd, err := os.Open("testdata/input/simple_utterance.json")
	require.NoError(t, err, "cannot open testdata file")

	defer fd.Close()

	var input Input
	err = json.NewDecoder(fd).Decode(&input)
	require.NoError(t, err, "cannot decode JSON file to struct")

	opts := cmp.Options{
		cmp.Comparer(func(a, b YandexDatetime) bool {
			return a.Unix() == b.Unix()
		}),
	}

	assert.True(t, cmp.Equal(expected, input, opts...), cmp.Diff(expected, input, opts...))
}

func TestInput_ButtonPressed(t *testing.T) {
	expected := Input{
		Meta: Meta{
			Locale:   "ru-RU",
			Timezone: "Europe/Moscow",
			ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
			Interfaces: Capabilities{
				Screen:         json.RawMessage(`{}`),
				Payments:       nil,
				AccountLinking: json.RawMessage(`{}`),
				AudioPlayer:    nil,
			},
		},
		Request: ButtonPressedRequest{
			Nlu: ButtonPressedRequestNlu{
				Tokens: []string{
					"надпись",
					"на",
					"кнопке",
				},
				Entities: []NamedEntity{},
				Intents:  json.RawMessage(`{}`),
			},
			Payload: json.RawMessage(`{}`),
		},
		Session: Session{
			MessageID: 0,
			SessionID: "2eac4854-fce721f3-b845abba-20d60",
			SkillID:   "3ad36498-f5rd-4079-a14b-788652932056",
			UserID:    "47C73714B580ED2469056E71081159529FFC676A4E5B059D629A819E857DC2F8",
			User: SessionUser{
				UserID:      "6C91DA5198D1758C6A9F63A7C5CDDF09359F683B13A18A151FBF4C8B092BB0C2",
				AccessToken: "AgAAAAAB4vpbAAApoR1oaCd5yR6eiXSHqOGT8dT",
			},
			Application: SessionApplication{
				ApplicationID: "47C73714B580ED2469056E71081159529FFC676A4E5B059D629A819E857DC2F8",
			},
			New: false,
		},
		State:   State{},
		Version: "1.0",
	}

	fd, err := os.Open("testdata/input/button_pressed.json")
	require.NoError(t, err, "cannot open testdata file")

	defer fd.Close()

	var input Input
	err = json.NewDecoder(fd).Decode(&input)
	require.NoError(t, err, "cannot decode JSON file to struct")

	assert.True(t, cmp.Equal(expected, input), cmp.Diff(expected, input))
}

func TestInput_ShowPull(t *testing.T) {
	expected := Input{
		Meta: Meta{
			Locale:   "ru-RU",
			Timezone: "Europe/Moscow",
			ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
			Interfaces: Capabilities{
				Screen: json.RawMessage(`{}`),
			},
		},
		Request: ShowPullRequest{
			ShowType: ShowTypeMorning,
		},
		Session: Session{},
		State:   State{},
		Version: "1.0",
	}

	fd, err := os.Open("testdata/input/show_pull.json")
	require.NoError(t, err, "cannot open testdata file")

	defer fd.Close()

	var input Input
	err = json.NewDecoder(fd).Decode(&input)
	require.NoError(t, err, "cannot decode JSON file to struct")

	assert.True(t, cmp.Equal(expected, input), cmp.Diff(expected, input))
}

func TestInput_PurchaseConfirmation(t *testing.T) {
	expected := Input{
		Meta: Meta{
			Locale:   "ru-RU",
			Timezone: "Europe/Moscow",
			ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
			Interfaces: Capabilities{
				Payments: json.RawMessage(`{}`),
			},
		},
		Request: PurchaseConfirmationRequest{
			PurchaseRequestId: "d432de19be8347d09f656d9fe966e2f9",
			PurchaseToken:     "token_value",
			OrderId:           "eeb59d64-9e6a-11ea-bb37-0242ac130002",
			PurchaseTimestamp: 1590399311,
			PurchasePayload:   json.RawMessage(`{"value": "payload"}`),
			SignedData:        "purchase_request_id=id_value&purchase_token=token_value&order_id=id_value",
			Signature:         "cHVyY2hhc2VfcmVxdWVzdF9pZD1pZF92YWx1ZSZwdXJjaGFzZV90b2tlbj10b2tlbl92YWx1ZSZvcmRlcl9pZD1pZF92YWx1ZQ==",
		},
		Session: Session{
			New: true,
		},
		State:   State{},
		Version: "1.0",
	}

	fd, err := os.Open("testdata/input/payment_confirmation.json")
	require.NoError(t, err, "cannot open testdata file")

	defer fd.Close()

	var input Input
	err = json.NewDecoder(fd).Decode(&input)
	require.NoError(t, err, "cannot decode JSON file to struct")

	assert.True(t, cmp.Equal(expected, input), cmp.Diff(expected, input))
}

func TestInput_AudioPlayer(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		expected Input
	}{
		{
			name: "playback_started",
			file: "testdata/input/audio_player_playback_started.json",
			expected: Input{
				Meta: Meta{
					Locale:   "ru-RU",
					Timezone: "Europe/Moscow",
					ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
					Interfaces: Capabilities{
						Payments:    json.RawMessage(`{}`),
						AudioPlayer: json.RawMessage(`{}`),
					},
				},
				Request: AudioPlayerRequest{
					requestType: TypeAudioPlayerPlaybackStarted,
				},
				Session: Session{New: true},
				State:   State{},
				Version: "1.0",
			},
		},
		{
			name: "playback_stopped",
			file: "testdata/input/audio_player_playback_stopped.json",
			expected: Input{
				Meta: Meta{
					Locale:   "ru-RU",
					Timezone: "Europe/Moscow",
					ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
					Interfaces: Capabilities{
						Payments:    json.RawMessage(`{}`),
						AudioPlayer: json.RawMessage(`{}`),
					},
				},
				Request: AudioPlayerRequest{
					requestType: TypeAudioPlayerPlaybackStopped,
				},
				Session: Session{New: true},
				State:   State{},
				Version: "1.0",
			},
		},
		{
			name: "playback_nearly_finished",
			file: "testdata/input/audio_player_playback_nearly_finished.json",
			expected: Input{
				Meta: Meta{
					Locale:   "ru-RU",
					Timezone: "Europe/Moscow",
					ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
					Interfaces: Capabilities{
						Payments:    json.RawMessage(`{}`),
						AudioPlayer: json.RawMessage(`{}`),
					},
				},
				Request: AudioPlayerRequest{
					requestType: TypeAudioPlayerPlaybackNearlyFinished,
				},
				Session: Session{New: true},
				State:   State{},
				Version: "1.0",
			},
		},
		{
			name: "playback_finished",
			file: "testdata/input/audio_player_playback_finished.json",
			expected: Input{
				Meta: Meta{
					Locale:   "ru-RU",
					Timezone: "Europe/Moscow",
					ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
					Interfaces: Capabilities{
						Payments:    json.RawMessage(`{}`),
						AudioPlayer: json.RawMessage(`{}`),
					},
				},
				Request: AudioPlayerRequest{
					requestType: TypeAudioPlayerPlaybackFinished,
				},
				Session: Session{New: true},
				State:   State{},
				Version: "1.0",
			},
		},
		{
			name: "playback_failed",
			file: "testdata/input/audio_player_playback_failed.json",
			expected: Input{
				Meta: Meta{
					Locale:   "ru-RU",
					Timezone: "Europe/Moscow",
					ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
					Interfaces: Capabilities{
						Payments:    json.RawMessage(`{}`),
						AudioPlayer: json.RawMessage(`{}`),
					},
				},
				Request: AudioPlayerErrorRequest{
					Error: AudioPlayerErrorData{
						Message: "fail details",
						Type:    MediaErrorServiceUnavailable,
					},
				},
				Session: Session{New: true},
				State:   State{},
				Version: "1.0",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fd, err := os.Open(tc.file)
			require.NoError(t, err, "cannot open testdata file")

			defer fd.Close()

			var input Input
			err = json.NewDecoder(fd).Decode(&input)
			require.NoError(t, err, "cannot decode JSON file to struct")

			opts := cmp.Options{
				cmp.AllowUnexported(AudioPlayerRequest{}),
			}

			assert.True(t, cmp.Equal(tc.expected, input, opts...), cmp.Diff(tc.expected, input, opts...))
		})
	}
}
