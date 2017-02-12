package lib

import (
	"fmt"
	"strings"
)

var (
	glider = [][]int{{3, 2}, {3, 3}, {3, 4}, {2, 4}, {1, 3}}

	pulsar = [][]int{
		{12, 16}, {13, 16}, {14, 16}, {10, 14}, {15, 14}, {10, 13}, {15, 13}, {10, 12}, {15, 12}, {12, 11}, {13, 11}, {14, 11}, {18, 16}, {19, 16}, {20, 16},
		{17, 14}, {22, 14}, {17, 13}, {22, 13}, {17, 12}, {22, 12}, {18, 11}, {19, 11}, {20, 11}, {12, 9}, {13, 9}, {14, 9}, {10, 8}, {15, 8}, {10, 7}, {15, 7}, {10, 6}, {15, 6}, {12, 4}, {13, 4}, {14, 4}, {18, 9}, {19, 9}, {20, 9},
		{17, 8}, {22, 8}, {17, 7}, {22, 7}, {17, 6}, {22, 6}, {18, 4}, {19, 4}, {20, 4}}

	gliderGun = [][]int{{3, 13}, {4, 13}, {3, 12}, {4, 12}, {13, 13}, {13, 12}, {13, 11}, {14, 14}, {14, 10}, {15, 15}, {16, 15}, {15, 9}, {16, 9}, {17, 12}, {19, 13}, {19, 12}, {19, 11}, {18, 10}, {20, 12}, {18, 14}, {23, 15}, {24, 15}, {23, 14}, {24, 14}, {23, 13}, {24, 13}, {25, 16}, {25, 12}, {27, 17}, {27, 16}, {27, 12}, {27, 11}, {37, 15}, {38, 15}, {37, 14}, {38, 14}}
)

// PromptPattern prompts player for the pattern player wants to see
func PromptPattern(s *string) ([][]int, int, int) {
	ClearScreen()
	fmt.Print("Choose a pattern-- Glider, Pulsar, or Gun?: ")
	fmt.Scanln(s)
	*s = strings.ToLower(*s)
	pattern, width, height := findPattern(*s)
	return pattern, width, height
}

func findPattern(s string) ([][]int, int, int) {
	var (
		pattern [][]int
		width   int
		height  int
	)
	switch s {
	case "glider":
		pattern, width, height = glider, 30, 15
	case "pulsar":
		pattern, width, height = pulsar, 40, 20
	case "gun":
		pattern, width, height = gliderGun, 50, 30
	default:
		var start string
		pattern, width, height = PromptPattern(&start)
	}
	return pattern, width, height
}
