package store

type ActionType string
type Action struct {
	Type ActionType `json:"type"`
}

const (
	TypeStartVoting        = "StartVoting"
	TypeReveal             = "Reveal"
	TypeReset              = "Reset"
	TypeChoose             = "Choose"
	TypeSetPlayerType      = "SetPlayerType"
	TypeSwitchCardBack     = "SwitchCardBack"
	TypeSetCards           = "SetCards"
	TypeSetAutoReveal      = "SetAutoReveal"
	TypeSetAutoRevealTimer = "SetAutoRevealTimer"
	TypeSetAdmin           = "SetAdmin"
	TypeReVote             = "ReVote"
)
