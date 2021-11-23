package dialogutil

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/bbrodriges/mielofon/v2/dialog"
)

type ResponseBuilder struct {
	Output dialog.Output
}

// NewResponse returns empty response builder
func NewResponse() *ResponseBuilder {
	return &ResponseBuilder{}
}

func (b *ResponseBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Output)
}

func (b *ResponseBuilder) Write(w io.Writer) (int, error) {
	j, err := json.Marshal(b)
	if err != nil {
		return 0, fmt.Errorf("cannot marshal response: %w", err)
	}
	return w.Write(j)
}

func (b *ResponseBuilder) SetText(t string) *ResponseBuilder {
	b.Output.Response.Text = t
	return b
}

func (b *ResponseBuilder) SetTTS(t string) *ResponseBuilder {
	b.Output.Response.TTS = t
	return b
}

func (b *ResponseBuilder) SetCard(c dialog.OutputCard) *ResponseBuilder {
	b.Output.Response.Card = c
	return b
}

func (b *ResponseBuilder) SetButtons(bs ...dialog.Button) *ResponseBuilder {
	b.Output.Response.Buttons = bs
	return b
}

func (b *ResponseBuilder) AddButtons(bs ...dialog.Button) *ResponseBuilder {
	b.Output.Response.Buttons = append(b.Output.Response.Buttons, bs...)
	return b
}

func (b *ResponseBuilder) SetDirectives(ds ...dialog.Directive) *ResponseBuilder {
	b.Output.Response.Directives = ds
	return b
}

func (b *ResponseBuilder) AddDirectives(ds ...dialog.Directive) *ResponseBuilder {
	b.Output.Response.Directives = append(b.Output.Response.Directives, ds...)
	return b
}

func (b *ResponseBuilder) EndSession(end bool) *ResponseBuilder {
	b.Output.Response.EndSession = end
	return b
}

func (b *ResponseBuilder) SetShowItemMeta(m dialog.ShowItemMeta) *ResponseBuilder {
	b.Output.Response.ShowItemMeta = m
	return b
}

func (b *ResponseBuilder) SetVersion(v string) *ResponseBuilder {
	b.Output.Version = v
	return b
}
