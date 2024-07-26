package pkg_test

import (
	"fmt"
	"github.com/vetali/sudoku/pkg"
	"testing"
)

func TestSudoku4x4(t *testing.T) {

	t.Run("Check update options", func(t *testing.T) {

		sudoku := pkg.NewSudoku([81]int{
			0, 3, 1, 0, 0, 9, 2, 5, 0,
			0, 7, 2, 5, 3, 1, 0, 8, 4,
			0, 4, 0, 0, 0, 7, 6, 0, 3,

			1, 0, 7, 0, 9, 0, 0, 0, 0,
			0, 9, 0, 0, 0, 8, 0, 0, 5,
			0, 0, 0, 7, 0, 3, 0, 0, 6,

			9, 0, 8, 0, 0, 4, 0, 6, 1,
			7, 0, 0, 1, 0, 0, 0, 0, 0,
			4, 1, 0, 9, 6, 2, 3, 0, 0,
		})

		sudoku.Print()
		err := sudoku.Solve(true)
		if err != nil {
			t.Fatalf("Sudoku is invalid, %v", err)
		}
		AssertTrue(t, sudoku.IsValid())
		AssertValues(t, sudoku.GetValues(), []int{
			8, 3, 1, 6, 4, 9, 2, 5, 7,
			6, 7, 2, 5, 3, 1, 9, 8, 4,
			5, 4, 9, 8, 2, 7, 6, 1, 3,

			1, 5, 7, 4, 9, 6, 8, 3, 2,
			3, 9, 6, 2, 1, 8, 7, 4, 5,
			2, 8, 4, 7, 5, 3, 1, 9, 6,

			9, 2, 8, 3, 7, 4, 5, 6, 1,
			7, 6, 3, 1, 8, 5, 4, 2, 9,
			4, 1, 5, 9, 6, 2, 3, 7, 8,
		})
	})
}

func BenchmarkSudoku4x4(b *testing.B) {

	for i := 0; i < b.N; i++ {
		pkg.NewSudoku([81]int{
			0, 3, 1, 0, 0, 9, 2, 5, 0,
			0, 7, 2, 5, 3, 1, 0, 8, 4,
			0, 4, 0, 0, 0, 7, 6, 0, 3,
			1, 0, 7, 0, 9, 0, 0, 0, 0,
			0, 9, 0, 0, 0, 8, 0, 0, 5,
			0, 0, 0, 7, 0, 3, 0, 0, 6,
			9, 0, 8, 0, 0, 4, 0, 6, 1,
			7, 0, 0, 1, 0, 0, 0, 0, 0,
			4, 1, 0, 9, 6, 2, 3, 0, 0,
		}).Solve(false)
	}
}

func AssertTrue(t *testing.T, b bool) {
	t.Helper()
	if !b {
		fmt.Errorf("sudoku is not valid, expected valid")
	}
}

func AssertValues(t *testing.T, expected, actual []int) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Errorf("sudoku solution is not valid, expected %v, got %v", expected, actual)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("sudoku solution is not valid, \n expected:\n %+v , got:\n %+v", expected, actual)
			break
		}
	}
}
