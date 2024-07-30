package pkg

import "sort"

func GetRelatedCells(cellIdx int) (related []int) {
	var colIdx, rowIdx, boxIdx int
	if cellIdx > 0 {
		rowIdx = cellIdx / 9
		colIdx = cellIdx % 9
		boxIdx = (colIdx / 3) + ((rowIdx / 3) * 3)
	}
	result := make(map[int]bool)

	for _, i := range boxes[boxIdx] {
		result[i] = true
	}

	for _, i := range cols[colIdx] {
		result[i] = true
	}

	for _, i := range rows[rowIdx] {
		result[i] = true
	}

	for i := range result {
		if i != cellIdx {
			related = append(related, i)
		}
	}
	sort.Ints(related)

	return
}
