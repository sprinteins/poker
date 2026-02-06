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
		Label: "Done",
		Cards: Cards{"Done"},
	},
}
