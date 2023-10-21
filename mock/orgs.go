package mock

import "github.com/Tesohh/xlearn/data"

var orgs = []data.Org{
	{
		Name:       "Silandro Investors",
		Tag:        "silandro-investors",
		Secret:     "tubetti",
		Adventures: []string{"forklift-certification-123456"},
		Codes: map[string]int{
			"123456": 20,
		},
	},
	{
		Name:       "Tubre Investors",
		Tag:        "tubre-investors",
		Secret:     "tubetti",
		Adventures: []string{"forklift-certification-abcdef"},
		Codes: map[string]int{
			"abcdef": 5,
		},
	},
	{
		Name:       "Merano Holdings LLC",
		Tag:        "merano-holdings",
		Secret:     "tubetti",
		Adventures: []string{"forklift-certification-fedcba"},
		Codes: map[string]int{
			"fedcba": 1,
		},
	},
}
