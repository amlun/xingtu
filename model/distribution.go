package model

type Distribution struct {
	Description string   `json:"description"`
	Type        int      `json:"type"`
	TypeDisplay string   `json:"type_display"`
	Image       []string `json:"image"`
}
