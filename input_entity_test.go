package mielofon

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUnmarshalEntity(t *testing.T) {
	testCases := []struct {
		name   string
		raw    []byte
		expect Entity
	}{
		{
			name: string(TypeYandexNumber),
			raw:  []byte(`{"tokens":{"start":5,"end":6},"type":"YANDEX.NUMBER","value":42}`),
			expect: Entity{
				Tokens: Tokens{
					Start: 5,
					End:   6,
				},
				Type:  TypeYandexNumber,
				Value: YandexNumber(42),
			},
		},
		{
			name: string(TypeYandexGeo),
			raw:  []byte(`{"tokens":{"start":2,"end":6},"type":"YANDEX.GEO","value":{"street":"льва толстого","house_number":"16"}}`),
			expect: Entity{
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
		},
		{
			name: string(TypeYandexFIO),
			raw:  []byte(`{"tokens":{"start":3,"end":5},"type":"YANDEX.FIO","value":{"first_name":"лев","last_name":"толстой"}}`),
			expect: Entity{
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
		},
		{
			name: string(TypeYandexDatetime),
			raw:  []byte(`{"tokens":{"start":6,"end":8},"type":"YANDEX.DATETIME","value":{"day":1,"day_is_relative":true}}`),
			expect: Entity{
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var actual Entity
			err := json.Unmarshal(tc.raw, &actual)
			assert.NoError(t, err)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func TestEntityType(t *testing.T) {
	assert.Equal(t, YandexGeo{}.Type(), TypeYandexGeo)
	assert.Equal(t, YandexFIO{}.Type(), TypeYandexFIO)
	assert.Equal(t, YandexNumber(42).Type(), TypeYandexNumber)
	assert.Equal(t, YandexDatetime{}.Type(), TypeYandexDatetime)
}

func TestYandexDatetimeIsRelative(t *testing.T) {
	testCases := []struct {
		name   string
		entity YandexDatetime
		expect bool
	}{
		{
			"non_relative",
			YandexDatetime{
				Year: 1984,
			},
			false,
		},
		{
			"relative_year",
			YandexDatetime{
				Year:           -2,
				YearIsRelative: true,
			},
			true,
		},
		{
			"relative_month",
			YandexDatetime{
				Month:           2,
				MonthIsRelative: true,
			},
			true,
		},
		{
			"relative_day",
			YandexDatetime{
				Day:           42,
				DayIsRelative: true,
			},
			true,
		},
		{
			"relative_hour",
			YandexDatetime{
				Hour:           -42,
				HourIsRelative: true,
			},
			true,
		},
		{
			"relative_minute",
			YandexDatetime{
				Minute:           0,
				MinuteIsRelative: true,
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, tc.entity.IsRelative())
		})
	}
}

func TestYandexDatetimeGetTime(t *testing.T) {
	testCases := []struct {
		name       string
		entity     YandexDatetime
		assertFunc func(t *testing.T, e YandexDatetime)
	}{
		{
			"nil_location",
			YandexDatetime{},
			func(t *testing.T, e YandexDatetime) {
				assert.Equal(t, time.Local, e.GetTime(nil).Location())
			},
		},
		{
			"what_happened_in_1984",
			YandexDatetime{
				Year:           1984,
				YearIsRelative: false,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now()
				assert.Equal(t, 1984, e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
			},
		},
		{
			"what_happened_2_years_ago",
			YandexDatetime{
				Year:           -2,
				YearIsRelative: true,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now().AddDate(-2, 0, 0)
				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
			},
		},
		{
			"what_happened_in_march",
			YandexDatetime{
				Month:           3,
				MonthIsRelative: false,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now()
				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, time.Month(3), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
			},
		},
		{
			"what_will_happen_in_a_month",
			YandexDatetime{
				Month:           1,
				MonthIsRelative: true,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now().AddDate(0, 1, 0)
				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
			},
		},
		{
			"what_happened_on_4th_day",
			YandexDatetime{
				Day:           4,
				DayIsRelative: false,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now()
				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, 4, e.GetTime(time.Local).Day())
			},
		},
		{
			"what_happened_5_days_ago",
			YandexDatetime{
				Day:           -5,
				DayIsRelative: true,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now().AddDate(0, 0, -5)
				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
			},
		},
		{
			"what_happened_at_5pm",
			YandexDatetime{
				Hour:           17,
				HourIsRelative: false,
			},
			func(t *testing.T, e YandexDatetime) {
				now := time.Now()
				expect := time.Date(now.Year(), now.Month(), now.Day(), 17, now.Minute(), now.Second(), now.Nanosecond(), now.Location())
				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
				assert.Equal(t, expect.Hour(), e.GetTime(time.Local).Hour())
			},
		},
		{
			"what_happened_3_hours_ago",
			YandexDatetime{
				Hour:           -3,
				HourIsRelative: true,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now().Add(-3 * time.Hour)
				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
				assert.Equal(t, expect.Hour(), e.GetTime(time.Local).Hour())
			},
		},
		{
			"what_happened_on_4th_of_july_2_years_ago_at_6am",
			YandexDatetime{
				Year:            -2,
				YearIsRelative:  true,
				Month:           7,
				MonthIsRelative: false,
				Day:             4,
				DayIsRelative:   false,
				Hour:            6,
				HourIsRelative:  false,
			},
			func(t *testing.T, e YandexDatetime) {
				expect := time.Now().AddDate(-2, 0, 0)
				expect = time.Date(expect.Year(), time.July, 4, 6, 0, 0, 0, expect.Location())

				assert.Equal(t, expect.Year(), e.GetTime(time.Local).Year())
				assert.Equal(t, expect.Month(), e.GetTime(time.Local).Month())
				assert.Equal(t, expect.Day(), e.GetTime(time.Local).Day())
				assert.Equal(t, expect.Hour(), e.GetTime(time.Local).Hour())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.assertFunc(t, tc.entity)
		})
	}
}
