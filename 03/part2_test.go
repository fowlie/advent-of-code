package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSymbols(t *testing.T) {
	string := "....*..*......*"
	assert.EqualValues(t, []int{4, 7, 14}, FindSymbols(string))
}

func TestPartNumbers(t *testing.T) {
	matches := FindPartNumbers("12..34..5")
	expectedMatches := []Part{
		{
			number:   12,
			position: 0,
		},
		{
			number:   34,
			position: 4,
		},
		{
			number:   5,
			position: 8,
		},
	}

	assert.EqualValues(t, expectedMatches, matches)
}

func TestFindNoGearOnlyOneSymbol(t *testing.T) {
	gears := FindGears([]ParsedLine{
		{
			// *
			parts:   []Part{},
			symbols: []int{0},
		},
	})
	assert.Zero(t, len(gears))
}

type AdjacentPartTest struct {
	str   string
	count int
}

func TestAdjacentParts(t *testing.T) {
	tests := []AdjacentPartTest{
		{"4*5", 2},
		{"...4*5", 2},
		{"4*5...", 2},
		{"0..4*5", 2},
		{".0.4*5", 2},
		{"4*5.0.", 2},
		{"1.2.3*4", 2},
	}
	for _, test := range tests {
		parts := FindPartNumbers(test.str)
		pos := strings.Index(test.str, "*")
		assert.Equal(t, test.count, len(FindAdjacentParts(pos, parts)))
	}
}

func TestFindOneGearSingleDigitsOnSameLine(t *testing.T) {
	gears := FindGears(testInput("3*5"))
	assert.Equal(t, 1, len(gears))
	assert.Equal(t, 3, gears[0].part1.number)
	assert.Equal(t, 5, gears[0].part2.number)
}

func TestFindOneGearOneLine(t *testing.T) {
	gears := FindGears(testInput("33*55"))
	assert.Equal(t, 1, len(gears))
	assert.Equal(t, 33, gears[0].part1.number)
	assert.Equal(t, 55, gears[0].part2.number)
}

func TestFindNoGearsOneLine(t *testing.T) {
	gears := FindGears(testInput("3.*.5"))
	assert.Equal(t, 0, len(gears))
}

func TestFindGearTwoLines(t *testing.T) {
	gears := FindGears(testInput(`
    2..
    .*4`))
	assert.Equal(t, 1, len(gears))
	assert.Equal(t, 4, gears[0].part1.number)
	assert.Equal(t, 2, gears[0].part2.number)
}

func TestFindGearThreeLines(t *testing.T) {
	gears := FindGears(testInput(`
    467..114..
    ...*......
    ..35..633.`))
	assert.Equal(t, 1, len(gears))
	assert.Equal(t, 467, gears[0].part1.number)
	assert.Equal(t, 35, gears[0].part2.number)
}

func TestFindNoGearTwoLines(t *testing.T) {
	gears := FindGears(ParseInput(
		splitAndTrim(`
      2.1
      .*4`),
	))
	assert.Equal(t, 0, len(gears))
}

func TestFindTwoGearsOneLine(t *testing.T) {
	gears := FindGears(testInput("2*1...3*4"))
	assert.Equal(t, 2, len(gears))
}

type TestCase struct {
	schematic string
	ratios    []int
}

func TestSuite(t *testing.T) {
	tests := []TestCase{
		{
			schematic: `
      2.1
      .*4`,
			ratios: []int{},
		},
		{
			schematic: `
      2.
      .*4`,
			ratios: []int{0, 8},
		},
	}

	for _, tc := range tests {
		ratios := CalculateRatios(FindGears(ParseInput(splitAndTrim(tc.schematic))))
		assert.EqualValues(t, tc.ratios, ratios)
	}
}

func TestMain(t *testing.T) {
	main()
}

func testInput(s string) []ParsedLine {
	return ParseInput(splitAndTrim(s))
}

func splitAndTrim(s string) []string {
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines
}
