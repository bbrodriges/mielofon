package mielofon

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
						Screen: make(map[string]interface{}),
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

func TestInputHelpers(t *testing.T) {
	locMoscow, _ := time.LoadLocation("Europe/Moscow")
	locLondon, _ := time.LoadLocation("Europe/London")

	testCases := []struct {
		name            string
		rawRequest      []byte
		expectLocation  *time.Location
		expectTime      time.Time
		expectHasScreen bool
		expectInitiated bool
	}{
		{
			"moscow",
			[]byte(`{"meta":{"timezone":"Europe/Moscow","interfaces":{"screen":{}}},"session":{"new":true}}`),
			locMoscow,
			time.Now().In(locMoscow),
			true,
			true,
		},
		{
			"london",
			[]byte(`{"meta":{"timezone":"Europe/London","interfaces":{}},"session":{"new":false}}`),
			locLondon,
			time.Now().In(locLondon),
			false,
			false,
		},
		{
			"unknown_location",
			[]byte(`{"meta":{"timezone":"Unknown/Location","interfaces":{"screen":{}}},"session":{"new":false}}`),
			time.Local,
			time.Now(),
			true,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var actualInput Input
			err := json.Unmarshal(tc.rawRequest, &actualInput)

			assert.NoError(t, err)

			assert.Equal(t, tc.expectLocation, actualInput.DeviceLocation())
			assert.Equal(t, tc.expectTime.Format(time.RFC3339), actualInput.DeviceTime().Format(time.RFC3339))
			assert.Equal(t, tc.expectHasScreen, actualInput.HasScreen())
			assert.Equal(t, tc.expectInitiated, actualInput.Initiated())
		})
	}
}
