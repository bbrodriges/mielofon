package dialogutil

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bbrodriges/mielofon/dialog"
)

func TestGetDialogPair(t *testing.T) {
	raw := bytes.NewBuffer(
		[]byte(`{"meta":{"client_id":"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)","interfaces":{"screen":{}},"locale":"ru-RU","timezone":"UTC"},"request":{"command":"hello, world","nlu":{"tokens":["hello","world"]},"original_utterance":"hello, world","type":"SimpleUtterance"},"session":{"message_id":1,"new":true,"session_id":"18021ad-ae9b2d24-c86e8926-6b64f1a","skill_id":"01b43d41-806b-4b95-8d19-87ba2a8edefa","user_id":"0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2"},"version":"1.0"}`),
	)

	expectInput := &dialog.Input{
		Meta: dialog.Meta{
			Locale:   "ru-RU",
			Timezone: "UTC",
			ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
			Interfaces: dialog.Interfaces{
				Screen: []byte(`{}`),
			},
		},
		Request: dialog.Request{
			Command:           "hello, world",
			OriginalUtterance: "hello, world",
			Type:              dialog.SimpleUtterance,
			Nlu: dialog.Nlu{
				Tokens: []string{"hello", "world"},
			},
		},
		Session: dialog.Session{
			New:       true,
			MessageID: 1,
			SessionID: "18021ad-ae9b2d24-c86e8926-6b64f1a",
			SkillID:   "01b43d41-806b-4b95-8d19-87ba2a8edefa",
			UserID:    "0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2",
		},
		Version: "1.0",
	}

	expectOutput := &dialog.Output{
		Version: "1.0",
		Session: dialog.Session{
			MessageID: 1,
			SessionID: "18021ad-ae9b2d24-c86e8926-6b64f1a",
			SkillID:   "01b43d41-806b-4b95-8d19-87ba2a8edefa",
			UserID:    "0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2",
		},
	}

	input, output, err := GetDialogPair(raw)
	assert.NoError(t, err)
	assert.Equal(t, expectInput, input)
	assert.Equal(t, expectOutput, output)

	// error
	input, output, err = GetDialogPair(bytes.NewBuffer(nil))
	assert.Error(t, err)
	assert.Nil(t, input)
	assert.Nil(t, output)
}

func TestReadInput(t *testing.T) {
	testCases := []struct {
		name      string
		input     io.Reader
		expect    *dialog.Input
		expectErr bool
	}{
		{
			"basic",
			bytes.NewBuffer(
				[]byte(`{"meta":{"client_id":"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)","interfaces":{"screen":{}},"locale":"ru-RU","timezone":"UTC"},"request":{"command":"hello, world","nlu":{"tokens":["hello","world"]},"original_utterance":"hello, world","type":"SimpleUtterance"},"session":{"message_id":1,"new":true,"session_id":"18021ad-ae9b2d24-c86e8926-6b64f1a","skill_id":"01b43d41-806b-4b95-8d19-87ba2a8edefa","user_id":"0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2"},"version":"1.0"}`),
			),
			&dialog.Input{
				Meta: dialog.Meta{
					Locale:   "ru-RU",
					Timezone: "UTC",
					ClientID: "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
					Interfaces: dialog.Interfaces{
						Screen: []byte(`{}`),
					},
				},
				Request: dialog.Request{
					Command:           "hello, world",
					OriginalUtterance: "hello, world",
					Type:              dialog.SimpleUtterance,
					Nlu: dialog.Nlu{
						Tokens: []string{"hello", "world"},
					},
				},
				Session: dialog.Session{
					New:       true,
					MessageID: 1,
					SessionID: "18021ad-ae9b2d24-c86e8926-6b64f1a",
					SkillID:   "01b43d41-806b-4b95-8d19-87ba2a8edefa",
					UserID:    "0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2",
				},
				Version: "1.0",
			},
			false,
		},
		{
			"json_unmarshal_error",
			bytes.NewBuffer(
				[]byte(`something bad here`),
			),
			nil,
			true,
		},
		{
			"empty_reader",
			bytes.NewBuffer(nil),
			nil,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ReadInput(tc.input)

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expect, actual)
		})
	}
}

func TestOutputFromInput(t *testing.T) {
	testCases := []struct {
		name   string
		input  *dialog.Input
		expect *dialog.Output
	}{
		{
			"basic",
			&dialog.Input{
				Version: "1.0",
				Session: dialog.Session{
					MessageID: 1,
					SessionID: "18021ad-ae9b2d24-c86e8926-6b64f1a",
					SkillID:   "01b43d41-806b-4b95-8d19-87ba2a8edefa",
					UserID:    "0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2",
				},
			},
			&dialog.Output{
				Version: "1.0",
				Session: dialog.Session{
					MessageID: 1,
					SessionID: "18021ad-ae9b2d24-c86e8926-6b64f1a",
					SkillID:   "01b43d41-806b-4b95-8d19-87ba2a8edefa",
					UserID:    "0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := OutputFromInput(tc.input)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
