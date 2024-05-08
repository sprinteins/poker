package store

type Cards = []string
type Set struct {
	Label string `json:"label"`
	Cards Cards  `json:"cards"`
}

var cardSets = []Set{
	{
		Label: "Scrum Fibonacci",
		Cards: Cards{"0", "1/2", "1", "2", "3", "5", "8", "13", "20", "40", "100", "∞", "?"},
	},
	{
		Label: "T-Shirt",
		Cards: Cards{"xs", "s", "m", "l", "xl", "xxl"},
	},
	{
		Label: "5 Finger Consensus",
		Cards: Cards{"1", "2", "3", "4", "5"},
	},
	{
		Label: "yes/no",
		Cards: Cards{"yes", "no"},
	},
	{
		Label: "Hours",
		Cards: Cards{"1h", "2h", "4h", "8h", "12h", "16h", "20h", "24h", "28h", "32h", "36h", "40h", "44h", "48h", "52h", "56h", "∞"},
	},
}
