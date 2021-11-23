package dialog

type Session struct {
	MessageID   int                `json:"message_id,omitempty"`
	SessionID   string             `json:"session_id,omitempty"`
	SkillID     string             `json:"skill_id,omitempty"`
	UserID      string             `json:"user_id,omitempty"`
	User        SessionUser        `json:"user"`
	Application SessionApplication `json:"application"`
	New         bool               `json:"new,omitempty"`
}

type SessionUser struct {
	UserID      string `json:"user_id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type SessionApplication struct {
	ApplicationID string `json:"application_id,omitempty"`
}
