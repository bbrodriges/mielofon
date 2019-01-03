package mielofon

import (
	"encoding/json"
	"io"

	"github.com/bbrodriges/mielofon/dialog"
)

// GetDialogPair is a convenient wrapper around ReadInput
// and OutputFromInput
func GetDialogPair(r io.Reader) (*dialog.Input, *dialog.Output, error) {
	input, err := ReadInput(r)
	if err != nil {
		return nil, nil, err
	}
	output := OutputFromInput(input)
	return input, output, nil
}

// ReadInput reads http POST body into dialog input struct
func ReadInput(r io.Reader) (*dialog.Input, error) {
	input := new(dialog.Input)
	dec := json.NewDecoder(r)
	if err := dec.Decode(input); err != nil {
		return nil, err
	}
	return input, nil
}

// OutputFromInput returns new dialog output with basic fields
// filled from dialog input
func OutputFromInput(input *dialog.Input) *dialog.Output {
	return &dialog.Output{
		Version: input.Version,
		Session: dialog.Session{
			MessageID: input.Session.MessageID,
			SessionID: input.Session.SessionID,
			UserID:    input.Session.UserID,
			SkillID:   input.Session.SkillID,
		},
	}
}
