package dialog

import (
	"encoding/json"
	"fmt"
	"time"
)

type Entity struct {
	Tokens Tokens       `json:"tokens,omitempty"`
	Type   EntityType   `json:"type,omitempty"`
	Value  YandexEntity `json:"value,omitempty"`
}

type Tokens struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

func (e *Entity) UnmarshalJSON(b []byte) error {
	var rawEntity = struct {
		Tokens Tokens          `json:"tokens,omitempty"`
		Type   EntityType      `json:"type,omitempty"`
		Value  json.RawMessage `json:"value,omitempty"`
	}{}

	if err := json.Unmarshal(b, &rawEntity); err != nil {
		return err
	}

	switch rawEntity.Type {
	case TypeYandexGeo:
		var value YandexGeo
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return err
		}
		e.Value = value
	case TypeYandexFIO:
		var value YandexFIO
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return err
		}
		e.Value = value
	case TypeYandexNumber:
		var value YandexNumber
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return err
		}
		e.Value = value
	case TypeYandexDatetime:
		var value YandexDatetime
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return err
		}
		e.Value = value
	default:
		return fmt.Errorf("cannot unmarshal value of Entity type %s", rawEntity.Type)
	}

	e.Tokens = rawEntity.Tokens
	e.Type = rawEntity.Type
	return nil
}

type EntityType string

const (
	TypeYandexGeo      EntityType = "YANDEX.GEO"
	TypeYandexFIO      EntityType = "YANDEX.FIO"
	TypeYandexNumber   EntityType = "YANDEX.NUMBER"
	TypeYandexDatetime EntityType = "YANDEX.DATETIME"
)

type YandexEntity interface {
	Type() EntityType
}

var (
	_ YandexEntity = YandexGeo{}
	_ YandexEntity = YandexFIO{}
	_ YandexEntity = YandexNumber(42)
	_ YandexEntity = YandexDatetime{}
)

// YandexGeo is a value of YANDEX.GEO entity type
type YandexGeo struct {
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
}

func (e YandexGeo) Type() EntityType {
	return TypeYandexGeo
}

// YandexFIO is a value of YANDEX.FIO entity type
type YandexFIO struct {
	FirstName      string `json:"first_name,omitempty"`
	PatronymicName string `json:"patronymic_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
}

func (e YandexFIO) Type() EntityType {
	return TypeYandexFIO
}

// YandexNumber is a value of YANDEX.NUMBER entity type
type YandexNumber float64

func (e YandexNumber) Type() EntityType {
	return TypeYandexNumber
}

// YandexDatetime is a value of YANDEX.DATETIME entity type
type YandexDatetime struct {
	Year             int  `json:"year,omitempty"`
	YearIsRelative   bool `json:"year_is_relative,omitempty"`
	Month            int  `json:"month,omitempty"`
	MonthIsRelative  bool `json:"month_is_relative,omitempty"`
	Day              int  `json:"day,omitempty"`
	DayIsRelative    bool `json:"day_is_relative,omitempty"`
	Hour             int  `json:"hour,omitempty"`
	HourIsRelative   bool `json:"hour_is_relative,omitempty"`
	Minute           int  `json:"minute,omitempty"`
	MinuteIsRelative bool `json:"minute_is_relative,omitempty"`
}

func (e YandexDatetime) Type() EntityType {
	return TypeYandexDatetime
}

// IsRelative checks if any part of YANDEX.DATETIME entity
// contains relative time
func (e YandexDatetime) IsRelative() bool {
	return e.YearIsRelative || e.MonthIsRelative ||
		e.DayIsRelative || e.HourIsRelative ||
		e.MinuteIsRelative
}

// GetTime returns proper time based on every YANDEX.DATETIME entity
// part in given location
func (e YandexDatetime) GetTime(l *time.Location) time.Time {
	if l == nil {
		l = time.Local
	}

	t := time.Now().In(l)

	if e.Year != 0 {
		if e.YearIsRelative {
			t = t.AddDate(e.Year, 0, 0)
		} else {
			t = time.Date(e.Year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		}
	}

	if e.Month != 0 {
		if e.MonthIsRelative {
			t = t.AddDate(0, e.Month, 0)
		} else {
			t = time.Date(t.Year(), time.Month(e.Month), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		}
	}

	if e.Day != 0 {
		if e.DayIsRelative {
			t = t.AddDate(0, 0, e.Day)
		} else {
			t = time.Date(t.Year(), t.Month(), e.Day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		}
	}

	if e.Hour != 0 {
		if e.HourIsRelative {
			t = t.Add(time.Duration(e.Hour) * time.Hour)
		} else {
			t = time.Date(t.Year(), t.Month(), t.Day(), e.Hour, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		}
	}

	if e.Minute != 0 {
		if e.MinuteIsRelative {
			t = t.Add(time.Duration(e.Minute) * time.Minute)
		} else {
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), e.Minute, t.Second(), t.Nanosecond(), t.Location())
		}
	}

	return t
}
