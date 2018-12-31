package mielofon

import (
	"encoding/json"
	"fmt"
	"time"
)

type Entity struct {
	Tokens Tokens       `json:"tokens"`
	Type   EntityType   `json:"type"`
	Value  YandexEntity `json:"value"`
}

type Tokens struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func (e *Entity) UnmarshalJSON(b []byte) error {
	var rawEntity = struct {
		Tokens Tokens          `json:"tokens"`
		Type   EntityType      `json:"type"`
		Value  json.RawMessage `json:"value"`
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
	Country     string `json:"country"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
}

func (e YandexGeo) Type() EntityType {
	return TypeYandexGeo
}

// YandexFIO is a value of YANDEX.FIO entity type
type YandexFIO struct {
	FirstName      string `json:"first_name"`
	PatronymicName string `json:"patronymic_name"`
	LastName       string `json:"last_name"`
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
	Year             int  `json:"year"`
	YearIsRelative   bool `json:"year_is_relative"`
	Month            int  `json:"month"`
	MonthIsRelative  bool `json:"month_is_relative"`
	Day              int  `json:"day"`
	DayIsRelative    bool `json:"day_is_relative"`
	Hour             int  `json:"hour"`
	HourIsRelative   bool `json:"hour_is_relative"`
	Minute           int  `json:"minute"`
	MinuteIsRelative bool `json:"minute_is_relative"`
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
