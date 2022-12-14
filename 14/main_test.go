package main

import (
	"testing"
)

func TestOneUnitOfSand(t *testing.T) {
  createLevel(15, 4)
  addLine(10, 3, 13, 3)
  addSand(11)

  expect(t, 1)
}

func TestTwoUnitsOfSand(t *testing.T) {
  createLevel(15, 4)
  addLine(10, 3, 13, 3)
  addSand(11)
  addSand(11)
  printIt()

  expect(t, 1)
}

// ...+...
// ...~...
// ..~o~..
// .~#o#~.
// .~#o#~.
// .~#o#~.
// .~###~.
func TestComplexLevelOneUnitOfSand(t *testing.T) {
  createLevel(7, 7)
  printIt()
  addLine(2, 3, 2, 6) //vertical
  addLine(4, 3, 4, 6) //vertical
  addLine(2, 6, 5, 6) //horizontal
  addNUnitsOfSand(10, 3)
  printIt()

  expect(t, 4)
}

// ..............+...............
// ..............................
// ..............o...............
// .............ooo..............
// ............ooooo.............
// ...........ooooooo............
// ..........#########...........
func TestSimpleWideLevel(t *testing.T) {
  createLevel(30, 10)
  addLine(10, 9, 19, 9)
  addNUnitsOfSand(17, 14)
  printIt()

  expect(t, 16)
}

func TestExampleInput(t *testing.T) {
  createLevel(510, 12)
  printOffset = 490
  enableDebug()

  addLine(498, 4, 498, 6)
  addLine(498, 6, 496, 6)

  addLine(503, 4, 502, 4)
  addLine(502, 4, 502, 9)
  addLine(502, 9, 494, 9)

  addNUnitsOfSand(30, 500)

  printIt()

  expect(t, 24)
}

func expect(t *testing.T, amount int) {
  if countSandUnits() != amount {
    t.Errorf("expected %d unit of sand, but got %d", amount, countSandUnits())
  }
}

func addNUnitsOfSand(units, x int) {
  for i := 0; i < units; i++ {
    addSand(x)
  }
}
