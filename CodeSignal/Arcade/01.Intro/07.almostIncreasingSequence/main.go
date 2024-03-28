package almostincreasingsequence

func almostIncreasingSequence(sequence []int) bool {
	removed := false
	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= sequence[i-1] {
			if removed {
				return false
			}
			removed = true
			if i == 1 || sequence[i] > sequence[i-2] {
				// remove index i-1
				sequence[i-1] = sequence[i]
			} else {
				// remove index i
				sequence[i] = sequence[i-1]
			}
		}

	}
	return true
}
