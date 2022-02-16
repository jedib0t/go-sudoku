package sudoku

import (
	"math/rand"
	"testing"

	"github.com/jedib0t/go-sudoku/sudoku/difficulty"
	"github.com/jedib0t/go-sudoku/sudoku/pattern"
	"github.com/stretchr/testify/assert"
)

var (
	testGridDone    = NewGrid(WithRNG(testRNG))
	testGridEmpty   = NewGrid(WithRNG(testRNG))
	testGridPartial = NewGrid(WithRNG(testRNG))
	testRNG         = rand.New(rand.NewSource(1))
)

func init() {
	_ = testGridDone.UnmarshalCSV(
		"4,7,9,5,8,1,2,3,6\n" +
			"8,2,3,6,9,7,1,5,4\n" +
			"1,6,5,3,2,4,8,7,9\n" +
			"7,3,6,8,1,2,4,9,5\n" +
			"2,9,8,4,3,5,7,6,1\n" +
			"5,4,1,7,6,9,3,2,8\n" +
			"9,1,4,2,7,6,5,8,3\n" +
			"6,8,7,1,5,3,9,4,2\n" +
			"3,5,2,9,4,8,6,1,7",
	)
	_ = testGridPartial.UnmarshalCSV(
		"5,3,7,0,0,0,0,8,6\n" +
			"4,0,0,0,9,6,5,0,3\n" +
			"9,0,0,0,0,5,1,7,0\n" +
			"2,4,0,0,3,7,0,9,0\n" +
			"7,0,0,8,0,0,6,3,1\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,6,0\n" +
			"8,6,0,9,7,3,0,1,0\n" +
			"0,9,0,2,6,0,0,5,8",
	)
}

func TestGrid_ApplyDifficulty(t *testing.T) {
	testGrid := testGridDone.Clone()
	assert.Equal(t, 81, testGrid.Count())
	testGrid.ApplyDifficulty(difficulty.Difficulty(36))
	assert.Equal(t, 36, testGrid.Count())

	testGrid2 := testGridDone.Clone()
	assert.Equal(t, 81, testGrid2.Count())
	testGrid2.ApplyDifficulty(difficulty.Difficulty(13))
	assert.Equal(t, 13, testGrid2.Count())
}

func TestGrid_ApplyPattern(t *testing.T) {
	testGrid := testGridDone.Clone()
	assert.Equal(t, 81, testGrid.Count())
	p := pattern.Get("triangle")
	assert.NotNil(t, p)
	testGrid.ApplyPattern(p)
	assert.Equal(t, 45, testGrid.Count())

	expectedCSV := `0,0,0,0,0,0,0,0,6
0,0,0,0,0,0,0,5,4
0,0,0,0,0,0,8,7,9
0,0,0,0,0,2,4,9,5
0,0,0,0,3,5,7,6,1
0,0,0,7,6,9,3,2,8
0,0,4,2,7,6,5,8,3
0,8,7,1,5,3,9,4,2
3,5,2,9,4,8,6,1,7`
	assert.Equal(t, expectedCSV, testGrid.MarshalCSV())
}

func TestGrid_Checksum(t *testing.T) {
	assert.Equal(t, "1e2a971", testGridDone.Checksum())
	assert.Equal(t, "3f73567", testGridPartial.Checksum())
}

func TestGrid_Clone(t *testing.T) {
	testGridClone := testGridDone.Clone()
	assert.Equal(t, 4, testGridDone.Get(0, 0))
	assert.Equal(t, 4, testGridClone.Get(0, 0))
	testGridClone.Reset(0, 0)
	assert.Equal(t, 4, testGridDone.Get(0, 0))
	assert.Equal(t, 0, testGridClone.Get(0, 0))
}

func TestGrid_CopySubGrid(t *testing.T) {
	testGrid := NewGrid(WithRNG(testRNG))
	assert.Equal(t, "[0 0 0 0 0 0 0 0 0]", testGrid.SubGrid(0, 0).String())
	err := testGrid.CopySubGrid(testGridDone, Location{X: 0, Y: 0}, Location{X: 0, Y: 0})
	assert.Nil(t, err)
	assert.Equal(t, "[4 7 9 8 2 3 1 6 5]", testGrid.SubGrid(0, 0).String())

	// invalid from
	err = testGrid.CopySubGrid(testGridDone, Location{X: 9, Y: 0}, Location{X: 0, Y: 0})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to find source sub-grid at [9, 0]")

	// invalid to
	err = testGrid.CopySubGrid(testGridDone, Location{X: 0, Y: 0}, Location{X: 9, Y: 0})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to find destination sub-grid at [9, 0]")
}

func TestGrid_Count(t *testing.T) {
	assert.Equal(t, 81, testGridDone.Count())
	assert.Equal(t, 0, testGridEmpty.Count())
	assert.Equal(t, 36, testGridPartial.Count())
}

func TestGrid_CountToDo(t *testing.T) {
	assert.Equal(t, 0, testGridDone.CountToDo())
	assert.Equal(t, 81, testGridEmpty.CountToDo())
	assert.Equal(t, 45, testGridPartial.CountToDo())
}

func TestGrid_Done(t *testing.T) {
	assert.Equal(t, true, testGridDone.Done())
	assert.Equal(t, false, testGridEmpty.Done())
	assert.Equal(t, false, testGridPartial.Done())
}

func TestGrid_Empty(t *testing.T) {
	assert.Equal(t, false, testGridDone.Empty())
	assert.Equal(t, true, testGridEmpty.Empty())
	assert.Equal(t, false, testGridPartial.Empty())
}

func TestGrid_Get(t *testing.T) {
	assert.Equal(t, 5, testGridPartial.Get(0, 0))
	assert.Equal(t, 0, testGridPartial.Get(1, 1))
	assert.Equal(t, 0, testGridPartial.Get(2, 2))
	assert.Equal(t, 0, testGridPartial.Get(3, 3))
	assert.Equal(t, 0, testGridPartial.Get(4, 4))
	assert.Equal(t, 0, testGridPartial.Get(5, 5))
	assert.Equal(t, 0, testGridPartial.Get(6, 6))
	assert.Equal(t, 1, testGridPartial.Get(7, 7))
	assert.Equal(t, 8, testGridPartial.Get(8, 8))
}

func TestGrid_GetAndSetMetadata(t *testing.T) {
	testGrid := testGridDone.Clone()
	assert.Equal(t, "", testGrid.GetMetadata("foo"))
	testGrid.SetMetadata("foo", "bar")
	assert.Equal(t, "bar", testGrid.GetMetadata("foo"))
}

func TestGrid_HasInColumn(t *testing.T) {
	assert.Equal(t, false, testGridPartial.HasInColumn(0, 1))
	assert.Equal(t, true, testGridPartial.HasInColumn(0, 2))
	assert.Equal(t, false, testGridPartial.HasInColumn(0, 3))
}

func TestGrid_HasInRow(t *testing.T) {
	assert.Equal(t, false, testGridPartial.HasInRow(0, 1))
	assert.Equal(t, false, testGridPartial.HasInRow(0, 2))
	assert.Equal(t, true, testGridPartial.HasInRow(0, 3))
}

func TestGrid_IsSet(t *testing.T) {
	assert.Equal(t, true, testGridPartial.IsSet(0, 1))
	assert.Equal(t, true, testGridPartial.IsSet(0, 2))
	assert.Equal(t, false, testGridPartial.IsSet(0, 3))
	assert.Equal(t, false, testGridPartial.IsSet(0, 4))
	assert.Equal(t, false, testGridPartial.IsSet(0, 5))
}

func TestGrid_MarshalArray(t *testing.T) {
	arr := testGridPartial.MarshalArray()
	assert.Len(t, arr, 9)
	assert.Equal(t, []int{5, 3, 7, 0, 0, 0, 0, 8, 6}, arr[0])
	assert.Equal(t, []int{4, 0, 0, 0, 9, 6, 5, 0, 3}, arr[1])
	assert.Equal(t, []int{9, 0, 0, 0, 0, 5, 1, 7, 0}, arr[2])
	assert.Equal(t, []int{2, 4, 0, 0, 3, 7, 0, 9, 0}, arr[3])
	assert.Equal(t, []int{7, 0, 0, 8, 0, 0, 6, 3, 1}, arr[4])
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, arr[5])
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 6, 0}, arr[6])
	assert.Equal(t, []int{8, 6, 0, 9, 7, 3, 0, 1, 0}, arr[7])
	assert.Equal(t, []int{0, 9, 0, 2, 6, 0, 0, 5, 8}, arr[8])
}

func TestGrid_MarshalCSV(t *testing.T) {
	expectedCSV := `4,7,9,5,8,1,2,3,6
8,2,3,6,9,7,1,5,4
1,6,5,3,2,4,8,7,9
7,3,6,8,1,2,4,9,5
2,9,8,4,3,5,7,6,1
5,4,1,7,6,9,3,2,8
9,1,4,2,7,6,5,8,3
6,8,7,1,5,3,9,4,2
3,5,2,9,4,8,6,1,7`
	csvString := testGridDone.MarshalCSV()
	assert.Equal(t, expectedCSV, csvString)

	expectedCSV2 := `5,3,7,0,0,0,0,8,6
4,0,0,0,9,6,5,0,3
9,0,0,0,0,5,1,7,0
2,4,0,0,3,7,0,9,0
7,0,0,8,0,0,6,3,1
0,0,0,0,0,0,0,0,0
0,0,0,0,0,0,0,6,0
8,6,0,9,7,3,0,1,0
0,9,0,2,6,0,0,5,8`
	csvString2 := testGridPartial.MarshalCSV()
	assert.Equal(t, expectedCSV2, csvString2)
}

func TestGrid_Possibilities(t *testing.T) {
	p := testGridPartial.Possibilities(0, 0)
	assert.Nil(t, p)
	p = testGridPartial.Possibilities(1, 1)
	assert.Equal(t, []int{1, 2, 8}, p.Available())
	p = testGridPartial.Possibilities(7, 6)
	assert.Equal(t, []int{2, 4}, p.Available())
}

func TestGrid_Reset(t *testing.T) {
	testGrid := testGridPartial.Clone()
	assert.Equal(t, true, testGrid.IsSet(0, 1))
	testGrid.Reset(0, 1)
	assert.Equal(t, false, testGrid.IsSet(0, 1))
}

func TestGrid_Set(t *testing.T) {
	testGrid := testGridPartial.Clone()
	assert.Equal(t, true, testGrid.IsSet(0, 1))
	assert.Equal(t, 3, testGrid.Get(0, 1))
	assert.False(t, testGrid.Set(0, 1, 5))
	assert.Equal(t, true, testGrid.IsSet(0, 1))
	assert.Equal(t, 3, testGrid.Get(0, 1))
	assert.True(t, testGrid.Set(0, 1, 1))
	assert.Equal(t, true, testGrid.IsSet(0, 1))
	assert.Equal(t, 1, testGrid.Get(0, 1))
}

func TestGrid_SetRNG(t *testing.T) {
	g := Grid{}
	assert.Nil(t, g.rng)
	assert.NotNil(t, g.getRNG())
	assert.NotEqual(t, testRNG, g.getRNG())
	g.SetRNG(testRNG)
	assert.Equal(t, testRNG, g.getRNG())
}

func TestGrid_String(t *testing.T) {
	expectedTestGridPartialCSV := `5,3,7,0,0,0,0,8,6
4,0,0,0,9,6,5,0,3
9,0,0,0,0,5,1,7,0
2,4,0,0,3,7,0,9,0
7,0,0,8,0,0,6,3,1
0,0,0,0,0,0,0,0,0
0,0,0,0,0,0,0,6,0
8,6,0,9,7,3,0,1,0
0,9,0,2,6,0,0,5,8`
	assert.Equal(t, expectedTestGridPartialCSV, testGridPartial.String())
}

func TestGrid_SubGrid(t *testing.T) {
	sg1 := testGridPartial.SubGrid(0, 0)
	assert.Equal(t, "[5 3 7 4 0 0 9 0 0]", sg1.String())
	sg2 := testGridPartial.SubGrid(3, 3)
	assert.Equal(t, "[0 3 7 8 0 0 0 0 0]", sg2.String())
	sg3 := testGridPartial.SubGrid(6, 6)
	assert.Equal(t, "[0 6 0 0 1 0 0 5 8]", sg3.String())
}

func TestGrid_UnmarshalArray(t *testing.T) {
	testGrid := testGridDone.Clone()
	testGrid2 := NewGrid(WithRNG(testRNG))
	err := testGrid2.UnmarshalArray(testGrid.MarshalArray())
	assert.Nil(t, err)
	assert.Equal(t, testGrid.MarshalCSV(), testGrid2.MarshalCSV())

	// invalid input case 1
	err = testGrid2.UnmarshalArray([][]int{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "input does not have 9 rows")

	// invalid input case 2
	err = testGrid2.UnmarshalArray([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "row[1] does not have 9 columns (0)")

	// invalid input case 3
	err = testGrid2.UnmarshalArray([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, -1},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "row[0] column[8] contains invalid number '-1'")
}

func TestGrid_UnmarshalCSV(t *testing.T) {
	testGrid := testGridDone.Clone()
	testGrid2 := NewGrid(WithRNG(testRNG))
	err := testGrid2.UnmarshalCSV(testGrid.MarshalCSV())
	assert.Nil(t, err)
	assert.Equal(t, testGrid.MarshalCSV(), testGrid2.MarshalCSV())

	// invalid input case 1
	err = testGrid2.UnmarshalCSV("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to read row #0: EOF")

	// invalid input case 2
	err = testGrid2.UnmarshalCSV("1,2,3,4,5,6,7,8")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "row[0] does not have 9 columns (8)")

	// invalid input case 3
	err = testGrid2.UnmarshalCSV("1,2,3,4,5,6,7,8,d")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "row[0] column[8] contains invalid value 'd'")

	// invalid input case 3
	err = testGrid2.UnmarshalCSV("1,2,3,4,5,6,7,8,11")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "row[0] column[8] contains invalid number '11'")
}

func TestGrid_Validate(t *testing.T) {
	isValid, err := testGridPartial.Validate()
	assert.True(t, isValid)
	assert.Nil(t, err)

	testGridInvalid := testGridPartial
	// 5,3,7,0,0,0,0,8,6
	// 4,0,0,0,9,6,5,0,3
	// 9,0,0,0,0,5,1,7,0
	// 2,4,0,0,3,7,0,9,0
	// 7,0,0,8,0,0,6,3,1
	// 0,0,0,0,0,0,0,0,0
	// 0,0,0,0,0,0,0,6,0
	// 8,6,0,9,7,3,0,1,0
	// 0,9,0,2,6,0,0,5,8

	// invalid case 1
	testGridInvalid.grid[0][0].value = 6
	isValid, err = testGridInvalid.Validate()
	assert.False(t, isValid)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "row[1] has more than one instance of '6'")

	// invalid case 2
	testGridInvalid.grid[0][0].value = 9
	isValid, err = testGridInvalid.Validate()
	assert.False(t, isValid)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "column[1] has more than one instance of '9'")

	// invalid case 3
	testGridInvalid.grid[0][0].value = 5
	testGridInvalid.grid[1][1].value = 7
	isValid, err = testGridInvalid.Validate()
	assert.False(t, isValid)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "sub-grid at [1, 1] is invalid: more than one instance of '7'")
}
