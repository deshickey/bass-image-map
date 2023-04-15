package music

type Scale struct {
	Index   int
	Name    string
	Root    string
	Pattern []int
}

func GetScale(scaleIndex int, root string) (scale []string) {
	scales := []Scale{
		{
			Index:   1,
			Name:    "Major",
			Root:    root,
			Pattern: []int{2, 2, 1, 2, 2, 2, 1},
		},
		{
			Index:   2,
			Name:    "Natural Minor",
			Root:    root,
			Pattern: []int{2, 1, 2, 2, 1, 2, 2},
		},
		{
			Index:   3,
			Name:    "Harmonic Major",
			Root:    root,
			Pattern: []int{2, 1, 2, 2, 1, 3, 1},
		},
		{
			Index:   4,
			Name:    "Melodic Major",
			Root:    root,
			Pattern: []int{2, 1, 2, 2, 2, 2, 1},
		},
		{
			Index:   5,
			Name:    "Major Pentatonic",
			Root:    root,
			Pattern: []int{2, 2, 3, 2, 3},
		},
		{
			Index:   6,
			Name:    "Minor Pentatonic",
			Root:    root,
			Pattern: []int{3, 2, 2, 3, 2},
		},
		{
			Index:   7,
			Name:    "Blues",
			Root:    root,
			Pattern: []int{3, 2, 1, 1, 3, 2},
		},
		{
			Index:   8,
			Name:    "Chromatic",
			Root:    root,
			Pattern: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Index:   9,
			Name:    "Half Whole Diminished",
			Root:    root,
			Pattern: []int{1, 2, 1, 2, 1, 2, 1, 2},
		},
		{
			Index:   10,
			Name:    "Whole Half Diminished",
			Root:    root,
			Pattern: []int{2, 1, 2, 1, 2, 1, 2, 1},
		},
		{
			Index:   11,
			Name:    "Bebop Dominant",
			Root:    root,
			Pattern: []int{2, 2, 1, 2, 2, 1, 1, 1},
		},
		{
			Index:   12,
			Name:    "Bebop Major",
			Root:    root,
			Pattern: []int{2, 2, 1, 2, 1, 1, 2, 1},
		},
		{
			Index:   13,
			Name:    "Harmonic Major",
			Root:    root,
			Pattern: []int{2, 2, 1, 2, 1, 3, 1},
		},
		{
			Index:   14,
			Name:    "Hungarian Minor",
			Root:    root,
			Pattern: []int{2, 1, 3, 1, 1, 3, 1},
		},
	}

	pattern := []int{}
	notes := []string{}

	for _, s := range scales {
		if scaleIndex == s.Index {
			pattern = s.Pattern
			notes = OrderNotes(s.Root)
		}
	}

	count := 0
	for _, p := range pattern {
		count = count + p
		scale = append(scale, notes[count])
	}
	return scale
}

func OrderNotes(root string) (orderedNotes []string) {
	allnotes := []string{"A", "As", "B", "C", "Cs", "D", "Ds", "E", "F", "Fs", "G", "Gs"}

	rootIndex := 0
	for i, j := range allnotes {
		if root == j {
			rootIndex = i
		}
	}

	var prenotes []int

	// find indexs before the new root index position
	for k := 0; k <= rootIndex; k++ {
		prenotes = append(prenotes, k)
	}

	// Notes in their new index position starting with root note
	for rootIndex < len(allnotes) {
		orderedNotes = append(orderedNotes, allnotes[rootIndex])
		rootIndex++
	}

	//append the first notes to the end
	for l := 0; l < len(prenotes); l++ {
		orderedNotes = append(orderedNotes, allnotes[l])
	}
	return orderedNotes
}
