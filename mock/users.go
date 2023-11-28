package mock

import "github.com/Tesohh/xlearn/data"

var Users = []data.User{
	{
		Display:      "Michele",
		Username:     "michele",
		PasswordHash: "$2a$10$2ohNV/1gszuQXPqHXdnpB.WRlwCCe5e.G9MmXH60QTxLPq2wRpciu",
		XP:           19092,
		Level:        67,
		Coins:        23,
		Role:         2,
		JoinedOrgs:   []string{"silandro-investors", "tubre-investors", "merano-holdings"},
		Settings: data.UserSettings{
			Language: "en_us",
			Theme:    "workshop",
		},
	},
	{
		Display:      "The Jolly Joker",
		Username:     "jollyjoker",
		PasswordHash: "$2a$10$2ohNV/1gszuQXPqHXdnpB.WRlwCCe5e.G9MmXH60QTxLPq2wRpciu",
		XP:           800,
		Level:        10,
		Coins:        60000000,
		Role:         1,
		JoinedOrgs:   []string{"silandro-investors", "tubre-investors"},
		Settings: data.UserSettings{
			Language: "en_us",
			Theme:    "workshop",
		},
	},
	{
		Display:      "Mr. MongoDB",
		Username:     "mr-mongodb",
		PasswordHash: "$2a$10$2ohNV/1gszuQXPqHXdnpB.WRlwCCe5e.G9MmXH60QTxLPq2wRpciu",
		XP:           850,
		Level:        11,
		Coins:        4,
		Role:         1,
		JoinedOrgs:   []string{"silandro-investors", "tubre-investors"},
		Settings: data.UserSettings{
			Language: "en_us",
			Theme:    "workshop",
		},
	},
	{
		Display:      "PolaroidKing123",
		Username:     "polaroidking123",
		PasswordHash: "$2a$10$2ohNV/1gszuQXPqHXdnpB.WRlwCCe5e.G9MmXH60QTxLPq2wRpciu",
		XP:           0,
		Level:        0,
		Coins:        6,
		Role:         0,
		JoinedOrgs:   []string{"silandro-investors", "tubre-investors"},
		Settings: data.UserSettings{
			Language: "it",
			Theme:    "workshop",
		},
	},
	{
		Display:      "Zesty Man",
		Username:     "zestyman",
		PasswordHash: "$2a$10$2ohNV/1gszuQXPqHXdnpB.WRlwCCe5e.G9MmXH60QTxLPq2wRpciu",
		XP:           6,
		Level:        5,
		Coins:        3,
		Role:         0,
		JoinedOrgs:   []string{"silandro-investors"},
		Settings: data.UserSettings{
			Language: "en_us",
			Theme:    "workshop",
		},
	},
}
