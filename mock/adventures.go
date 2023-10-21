package mock

import "github.com/Tesohh/xlearn/data"

// "forklift-certification-123456"
// "forklift-certification-abcdef"
// "forklift-certification-fedcba"

var adventures = []data.Adventure{
	{
		Name:        "Forklift Certification (for silandrish people)",
		Tag:         "forklift-certification-123456",
		Description: "Learn how to drive a forklift",
		Steps:       []string{"forkliftstep1-123456"},
	},
	{
		Name:        "Forklift Certification (for tubrish people)",
		Tag:         "forklift-certification-abcdef",
		Description: "Learn how to drive a forklift",
		Steps:       []string{"forkliftstep1-abcdef", "forkliftstep-parent-abcdef", "forkliftstep-join-abcdef"},
	},
	{
		Name:        "Forklift Certification (for meranish people)",
		Tag:         "forklift-certification-fedcba",
		Description: "Learn how to drive a forklift",
		Steps:       []string{"forkliftstep1-fedcba"},
	},
}
