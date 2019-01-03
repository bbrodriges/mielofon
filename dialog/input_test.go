package dialog

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInputUnmarshal(t *testing.T) {
	testCases := []struct {
		name        string
		rawRequest  []byte
		expectInput Input
		expectError bool
	}{
		{
			"docs example",
			[]byte(`{"meta":{"locale":"ru-RU","timezone":"Europe/Moscow","client_id":"ru.yandex.searchplugin/5.80 (Samsung Galaxy; Android 4.4)","interfaces":{"screen":{}}},"request":{"command":"закажи пиццу на улицу льва толстого, 16 на завтра","original_utterance":"закажи пиццу на улицу льва толстого, 16 на завтра","type":"SimpleUtterance","markup":{"dangerous_context":true},"payload":{},"nlu":{"tokens":["закажи","пиццу","на","льва","толстого","16","на","завтра"],"entities":[{"tokens":{"start":2,"end":6},"type":"YANDEX.GEO","value":{"house_number":"16","street":"льва толстого"}},{"tokens":{"start":3,"end":5},"type":"YANDEX.FIO","value":{"first_name":"лев","last_name":"толстой"}},{"tokens":{"start":5,"end":6},"type":"YANDEX.NUMBER","value":16},{"tokens":{"start":6,"end":8},"type":"YANDEX.DATETIME","value":{"day":1,"day_is_relative":true}}]}},"session":{"new":true,"message_id":4,"session_id":"2eac4854-fce721f3-b845abba-20d60","skill_id":"3ad36498-f5rd-4079-a14b-788652932056","user_id":"AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC"},"version":"1.0"}`),
			Input{
				Meta: Meta{
					Locale:   "ru-RU",
					Timezone: "Europe/Moscow",
					ClientID: "ru.yandex.searchplugin/5.80 (Samsung Galaxy; Android 4.4)",
					Interfaces: Interfaces{
						Screen: []byte(`{}`),
					},
				},
				Request: Request{
					Command:           "закажи пиццу на улицу льва толстого, 16 на завтра",
					OriginalUtterance: "закажи пиццу на улицу льва толстого, 16 на завтра",
					Type:              SimpleUtterance,
					Markup: Markup{
						DangerousContext: true,
					},
					Payload: []byte(`{}`),
					Nlu: Nlu{
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
						Entities: []Entity{
							{
								Tokens: Tokens{
									Start: 2,
									End:   6,
								},
								Type: TypeYandexGeo,
								Value: YandexGeo{
									Street:      "льва толстого",
									HouseNumber: "16",
								},
							},
							{
								Tokens: Tokens{
									Start: 3,
									End:   5,
								},
								Type: TypeYandexFIO,
								Value: YandexFIO{
									FirstName: "лев",
									LastName:  "толстой",
								},
							},
							{
								Tokens: Tokens{
									Start: 5,
									End:   6,
								},
								Type:  TypeYandexNumber,
								Value: YandexNumber(16),
							},
							{
								Tokens: Tokens{
									Start: 6,
									End:   8,
								},
								Type: TypeYandexDatetime,
								Value: YandexDatetime{
									Day:           1,
									DayIsRelative: true,
								},
							},
						},
					},
				},
				Session: Session{
					New:       true,
					MessageID: 4,
					SessionID: "2eac4854-fce721f3-b845abba-20d60",
					SkillID:   "3ad36498-f5rd-4079-a14b-788652932056",
					UserID:    "AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC",
				},
				Version: "1.0",
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var actualInput Input
			err := json.Unmarshal(tc.rawRequest, &actualInput)

			if !tc.expectError {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			assert.Equal(t, tc.expectInput, actualInput)
		})
	}
}

func TestHasScreen(t *testing.T) {
	input := Input{Meta: Meta{Interfaces: Interfaces{Screen: []byte(`{}`)}}}
	assert.True(t, input.HasScreen())

	input = Input{Meta: Meta{Interfaces: Interfaces{Screen: nil}}}
	assert.False(t, input.HasScreen())
}

func TestDeviceLocation(t *testing.T) {
	locMoscow, _ := time.LoadLocation("Europe/Moscow")
	input := Input{Meta: Meta{Timezone: "Europe/Moscow"}}
	assert.Equal(t, locMoscow, input.DeviceLocation())

	locLondon, _ := time.LoadLocation("Europe/London")
	input = Input{Meta: Meta{Timezone: "Europe/London"}}
	assert.Equal(t, locLondon, input.DeviceLocation())

	input = Input{Meta: Meta{Timezone: "Unknown/Unknown"}}
	assert.Equal(t, time.Local, input.DeviceLocation())
}

func TestDeviceTime(t *testing.T) {
	locMoscow, _ := time.LoadLocation("Europe/Moscow")
	expect := time.Now().In(locMoscow)
	input := Input{Meta: Meta{Timezone: "Europe/Moscow"}}
	assert.Equal(t, expect.Format(time.RFC3339), input.DeviceTime().Format(time.RFC3339))

	locLondon, _ := time.LoadLocation("Europe/London")
	expect = time.Now().In(locLondon)
	input = Input{Meta: Meta{Timezone: "Europe/London"}}
	assert.Equal(t, expect.Format(time.RFC3339), input.DeviceTime().Format(time.RFC3339))

	input = Input{Meta: Meta{Timezone: "Unknown/Unknown"}}
	assert.Equal(t, time.Now().Format(time.RFC3339), input.DeviceTime().Format(time.RFC3339))
}

func TestDialogInitiated(t *testing.T) {
	input := Input{Session: Session{New: true}}
	assert.True(t, input.DialogInitiated())

	input = Input{Session: Session{New: false}}
	assert.False(t, input.DialogInitiated())
}

func TestHasEntities(t *testing.T) {
	input := Input{
		Request: Request{
			Nlu: Nlu{
				Entities: []Entity{
					{
						Type:  TypeYandexNumber,
						Value: YandexNumber(42),
					},
					{
						Type: TypeYandexFIO,
						Value: YandexFIO{
							FirstName: "Freddy",
							LastName:  "Mercury",
						},
					},
				},
			},
		},
	}

	assert.True(t, input.HasEntities(TypeYandexNumber))
	assert.False(t, input.HasEntities(TypeYandexGeo))
	assert.True(t, input.HasEntities(TypeYandexFIO))
	assert.False(t, input.HasEntities(TypeYandexDatetime))
}

func TestEntities(t *testing.T) {
	input := Input{
		Request: Request{
			Nlu: Nlu{
				Entities: []Entity{
					{
						Type:  TypeYandexNumber,
						Value: YandexNumber(42),
					},
					{
						Type: TypeYandexFIO,
						Value: YandexFIO{
							FirstName: "Douglas",
							LastName:  "Adams",
						},
					},
					{
						Type: TypeYandexFIO,
						Value: YandexFIO{
							FirstName: "Arthur",
							LastName:  "Dent",
						},
					},
				},
			},
		},
	}

	ents := input.Entities(TypeYandexFIO)
	assert.Equal(t, 2, len(ents))
	assert.Equal(t, []Entity{
		{
			Type: TypeYandexFIO,
			Value: YandexFIO{
				FirstName: "Douglas",
				LastName:  "Adams",
			},
		},
		{
			Type: TypeYandexFIO,
			Value: YandexFIO{
				FirstName: "Arthur",
				LastName:  "Dent",
			},
		},
	}, ents)

	ents = input.Entities(TypeYandexNumber)
	assert.Equal(t, 1, len(ents))
	assert.Equal(t, []Entity{
		{
			Type:  TypeYandexNumber,
			Value: YandexNumber(42),
		},
	}, ents)

	ents = input.Entities(TypeYandexGeo)
	assert.Equal(t, 0, len(ents))
	assert.Equal(t, []Entity{}, ents)

	ents = input.Entities(TypeYandexDatetime)
	assert.Equal(t, 0, len(ents))
	assert.Equal(t, []Entity{}, ents)
}

func TestHasKeyword(t *testing.T) {
	input := Input{
		Request: Request{
			Nlu: Nlu{
				Tokens: []string{"hello", "world"},
			},
		},
	}

	assert.True(t, input.HasKeyword("hello"))
	assert.False(t, input.HasKeyword("goodbye"))
}

func TestHasKeywords(t *testing.T) {
	input := Input{
		Request: Request{
			Nlu: Nlu{
				Tokens: []string{"hello", "world", "and", "thanks", "for", "the", "fish"},
			},
		},
	}

	expect := []string{"hello", "thanks", "fish"}
	ks := input.HasKeywords(expect...)
	assert.Equal(t, len(expect), len(ks))
	assert.Equal(t, expect, ks)

	expect = []string{"hello", "thanks"}
	ks = input.HasKeywords(append(expect, "goodbye")...)
	assert.Equal(t, len(expect), len(ks))
	assert.Equal(t, expect, ks)

	// test on empty keywords
	input = Input{
		Request: Request{
			Nlu: Nlu{
				Tokens: nil,
			},
		},
	}

	expect = nil
	ks = input.HasKeywords("hello", "thanks", "fish")
	assert.Equal(t, 0, len(ks))
	assert.Equal(t, expect, ks)

	// test on lesser tokens
	input = Input{
		Request: Request{
			Nlu: Nlu{
				Tokens: []string{"hello", "world"},
			},
		},
	}

	expect = []string{"hello"}
	ks = input.HasKeywords("hello", "thanks", "fish")
	assert.Equal(t, len(expect), len(ks))
	assert.Equal(t, expect, ks)
}
