package dialog

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DirectiveType string

const (
	DirectiveAudioPlayer         DirectiveType = "audio_player"
	DirectiveStartAccountLinking DirectiveType = "start_account_linking"
	DirectiveStartPurchase       DirectiveType = "start_purchase"
	DirectiveConfirmPurchase     DirectiveType = "confirm_purchase"
)

// Directive represents abstract directive field of input
type Directive interface {
	Type() DirectiveType
}

// Directives is a set of output directives
type Directives []Directive

func (d Directives) MarshalJSON() ([]byte, error) {
	if len(d) == 0 {
		return []byte{}, nil
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	buf.WriteByte('{')
	for i, directive := range d {
		if i > 0 {
			buf.WriteByte(',')
		}
		dt := directive.Type()
		buf.WriteString(`"` + string(dt) + `":`)
		if err := enc.Encode(directive); err != nil {
			return nil, fmt.Errorf("cannot encode directive %s: %w", dt, err)
		}
	}
	buf.WriteByte('}')

	return buf.Bytes(), nil
}
