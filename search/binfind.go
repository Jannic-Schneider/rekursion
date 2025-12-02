package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {

	if len(list) == 0 {
		return -1
	}

	middlepos := len(list) / 2
	middle := list[middlepos]

	if x == middle {
		return middlepos
	}
	if x < middle {
		return FindSorted(list[:middlepos], x)

	}
	if x > middle {
		pos := FindSorted(list[middlepos+1:], x)
		if pos == -1 {
			return -1
		}
		return pos + middlepos + 1

	}
	return -1
}
