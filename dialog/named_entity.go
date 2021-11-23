package dialog

import (
	"encoding/json"
	"fmt"
	"time"
)

type NamedEntityType string

const (
	TypeYandexGeo      NamedEntityType = "YANDEX.GEO"
	TypeYandexFIO      NamedEntityType = "YANDEX.FIO"
	TypeYandexNumber   NamedEntityType = "YANDEX.NUMBER"
	TypeYandexDatetime NamedEntityType = "YANDEX.DATETIME"
)

// NamedEntity is a universal named entity
type NamedEntity struct {
	Tokens TokensRange     `json:"tokens"`
	Type   NamedEntityType `json:"type"`
	Value  interface{}     `json:"value"`
}

type TokensRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// YandexGeo is a value of YANDEX.GEO entity type
type YandexGeo struct {
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
	Airport     string `json:"airport,omitempty"`
}

// YandexFIO is a value of YANDEX.FIO entity type
type YandexFIO struct {
	FirstName      string `json:"first_name,omitempty"`
	PatronymicName string `json:"patronymic_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
}

// YandexNumber is a value of YANDEX.NUMBER entity type
type YandexNumber float64

// YandexDatetime is a value of YANDEX.DATETIME entity type
type YandexDatetime struct {
	time.Time
}

func (e *NamedEntity) UnmarshalJSON(b []byte) error {
	var rawEntity = struct {
		Tokens TokensRange     `json:"tokens"`
		Type   NamedEntityType `json:"type"`
		Value  json.RawMessage `json:"value,omitempty"`
	}{}

	if err := json.Unmarshal(b, &rawEntity); err != nil {
		return fmt.Errorf("cannot unmarshal named entity: %w", err)
	}

	switch rawEntity.Type {
	case TypeYandexGeo:
		var value YandexGeo
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return fmt.Errorf("cannot unmarshal named entity of type %s: %w", rawEntity.Type, err)
		}
		e.Value = value
	case TypeYandexFIO:
		var value YandexFIO
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return fmt.Errorf("cannot unmarshal named entity of type %s: %w", rawEntity.Type, err)
		}
		e.Value = value
	case TypeYandexNumber:
		var value YandexNumber
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return fmt.Errorf("cannot unmarshal named entity of type %s: %w", rawEntity.Type, err)
		}
		e.Value = value
	case TypeYandexDatetime:
		var value YandexDatetime
		if err := json.Unmarshal(rawEntity.Value, &value); err != nil {
			return fmt.Errorf("cannot unmarshal named entity of type %s: %w", rawEntity.Type, err)
		}
		e.Value = value
	default:
		// put raw bytes of unknown type
		e.Value = rawEntity.Value
	}

	e.Type = rawEntity.Type
	e.Tokens = rawEntity.Tokens
	return nil
}

func (d *YandexDatetime) UnmarshalJSON(b []byte) error {
	var raw = struct {
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
	}{}

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	t := time.Now().Truncate(time.Minute)

	if raw.Year != 0 {
		if raw.YearIsRelative {
			t = t.AddDate(raw.Year, 0, 0)
		} else {
			t = time.Date(raw.Year, 0, 0, 0, 0, 0, 0, time.Local)
		}
	}

	if raw.Month != 0 {
		if raw.MonthIsRelative {
			t = t.AddDate(0, raw.Month, 0)
		} else {
			t = time.Date(t.Year(), time.Month(raw.Month), 0, 0, 0, 0, 0, time.Local)
		}
	}

	if raw.Day != 0 {
		if raw.DayIsRelative {
			t = t.AddDate(0, 0, raw.Day)
		} else {
			t = time.Date(t.Year(), t.Month(), raw.Day, 0, 0, 0, 0, time.Local)
		}
	}

	if raw.Hour != 0 {
		if raw.HourIsRelative {
			t = t.Add(time.Duration(raw.Hour) * time.Hour)
		} else {
			t = time.Date(t.Year(), t.Month(), t.Day(), raw.Hour, 0, 0, 0, time.Local)
		}
	}

	if raw.Minute != 0 {
		if raw.MinuteIsRelative {
			t = t.Add(time.Duration(raw.Minute) * time.Minute)
		} else {
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), raw.Minute, 0, 0, time.Local)
		}
	}

	d.Time = t
	return nil
}
