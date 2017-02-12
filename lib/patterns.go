package lib

// Patterns is a struct containing all accessible patterns
type Patterns struct {
	Glider  [][]int
	Blinker [][]int
	Pulsar  [][]int
}

// CreatePatterns forms a struct containing the patterns with coordinates
func CreatePatterns() Patterns {
	glider := [][]int{{3, 2}, {3, 3}, {3, 4}, {2, 4}, {1, 3}}

	blinker := [][]int{{3, 1}, {3, 2}, {3, 3}}

	pulsar := [][]int{
		{12, 16}, {13, 16}, {14, 16}, {10, 14}, {15, 14}, {10, 13}, {15, 13}, {10, 12}, {15, 12}, {12, 11}, {13, 11}, {14, 11}, {18, 16}, {19, 16}, {20, 16},
		{17, 14}, {22, 14}, {17, 13}, {22, 13}, {17, 12}, {22, 12}, {18, 11}, {19, 11}, {20, 11}, {12, 9}, {13, 9}, {14, 9}, {10, 8}, {15, 8}, {10, 7}, {15, 7}, {10, 6}, {15, 6}, {12, 4}, {13, 4}, {14, 4}, {18, 9}, {19, 9}, {20, 9},
		{17, 8}, {22, 8}, {17, 7}, {22, 7}, {17, 6}, {22, 6}, {18, 4}, {19, 4}, {20, 4}}
	patterns := Patterns{glider, blinker, pulsar}
	return patterns
}
