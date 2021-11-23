package dialog

var _ Request = (*ShowPullRequest)(nil)

type ShowType string

const (
	ShowTypeMorning ShowType = "MORNING"
)

type ShowPullRequest struct {
	ShowType ShowType `json:"show_type"`
}

func (ShowPullRequest) Type() RequestType {
	return TypeShowPull
}
