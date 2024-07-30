package pkg_test

import (
	"github.com/vetali/sudoku/pkg"
	"testing"
)

func TestGetRelatedCells(t *testing.T) {
	t.Run("get related cells for (0)[0,0]", func(t *testing.T) {

		related := pkg.GetRelatedCells(0)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 18, 19, 20, 27, 36, 45, 54, 63, 72}

		AssertSameArray(t, related, want)
	})
}

func AssertSameArray[T comparable](t *testing.T, got, want []T) {
	if len(got) != len(want) {
		t.Fatalf("expected arrays have different length. got %d want %d", len(got), len(want))
	}
	exists := make(map[T]bool)

	for _, v := range got {
		exists[v] = true
	}
	for _, v := range want {
		if !exists[v] {
			t.Fatalf("arrays do not have the same items.\n got %+v \n want %+v", got, want)
		}
	}
}
