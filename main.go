package main

import (
	"fmt"
	"strconv"
)

type cell struct  {
	value int
	given bool
	row int
	column int
}

type row []cell

func (r row) String() string {
	result := ""

	for _, cell := range r {
		result += fmt.Sprintf("%d ", cell.value)
	}
	result += "\n"

	return result
}

func (r row) contains(n int) (bool, cell) {
	for _, cell := range r {
		if cell.value == n {
			return true, cell
		}
	}
	return false, cell{}
}

type column []cell

func (c column) String() string {
	result := ""

	for _, cell := range c {
		result += fmt.Sprintf("%d ", cell.value)
	}
	result += "\n"

	return result
}

func (c column) contains(n int) (bool, cell) {
	for _, cell := range c {
		if cell.value == n {
			return true, cell
		}
	}
	return false, cell{}
}

type square []cell

func (s square) String() string {
	result := ""

	for i, cell := range s {
		result += fmt.Sprintf("%d ", cell.value)
		if i == 2 || i == 5 {
			result += "\n"
		}
	}

	return result
}

func (s square) contains(n int) (bool, cell) {
	for _, cell := range s {
		if cell.value == n {
			return true, cell
		}
	}
	return false, cell{}
}

type grid []cell

func (g grid) String() string {
	result := "\n"

	for row := 0; row < 9; row++ {
		for col, c:= range g[:9] {
			result += fmt.Sprintf("%d ", c.value)
			if col == 2 || col==5 {
				result += "|"
			}
		}
		result += "\n"
		if row==2 || row==5 {
			result += "------+------+-----"
			result += "\n"
		}
		g = g[9:]
	}

	return result;
}

func (g grid) rows() []row {
	result := make([]row, 9)

	for r := 0; r < 9; r++ {
		result[r] = row(g[:9])
		g = g[9:]
	}

	return result
}

func (g grid) columns() []column {
	result := make([]column, 9)
	rows := g.rows()
	
	for c := 0; c < 9; c++ {
		result[c] = make([]cell, 9)
		for r := 0; r < 9; r++ {
			result[c][r] = rows[r][c]
		}
	}

	return result
}

func (g grid) squares() []square {
	result := make([]square, 9)
	for s := 0; s < 9; s++ {
		var row int = s / 3
		var col int = s % 3
		square := make([]cell, 9)
		si := 0
		for r := row*3; r < row*3 + 3; r++ {
			for c := col*3; c < col*3 + 3; c++ {
				square[si] = g[r*9 + c]
				si++
			}
		}
		result[s] = square
	}
	return result
}

func (g grid) isComplete() bool {
	result := true

	for _, cell := range g {
		if cell.value == 0 {
			result = false
			break
		}
	}

	return result
}

func simpleSolve(g grid) (bool, grid) {
	changed := false
	for guess := 1 ; guess < 10; guess++ {
		for _, square := range g.squares() {
			answer := cell{}
			answerCount := 0
			if ok, _ := square.contains(guess); !ok {
				for _, trial := range square {
					if(trial.value == 0) {
						if ok, _ := g.rows()[trial.row].contains(guess); !ok {
							if ok, _ := g.columns()[trial.column].contains(guess); !ok {
								answer = cell {
									row: trial.row,
									column: trial.column,
									given: trial.given,
									value: guess,
								}
								answerCount++
							} else {
								continue
							}
						} else {
							continue
						}
					}
				}
				if(answerCount == 1) {
					g[answer.row*9+answer.column].value = answer.value
					changed = true
				}
			}
		}
	}
	return changed, g
}

func buildFromString(s string) grid {
	if len(s) != 81 {
		panic("Grid string must be 81 numeric character.")
	}

	result := make(grid, 81)

	for i, c := range s {
		d, err := strconv.Atoi(string(c))
		if(err != nil) {
			panic(err)
		}
		result[i] = cell{value: d, given: c!='0', row: i/9, column: i%9}
	}

	return result
}

func buildTestGrid1() grid {
	return buildFromString("000090000065000380030000050000209000709603208000401000607000102040506090092030560")
}

func buildTestGrid2() grid {
	return buildFromString("000005019063000000050070638008027050009000800020310900634090080000000260870500000")
}

func buildTestGrid3() grid {
	return buildFromString("400086000020007100007250090500010043000008602600300900150000200003000080208503004")
}

func main() {
	fmt.Println()

	grid := buildTestGrid3()

	fmt.Println("Solving:")
	fmt.Println(grid)

	iterations := 0

	for changed, grid := false, grid;
		!grid.isComplete() || !changed;
		changed, grid = simpleSolve(grid) {
		iterations++
	}

	fmt.Println("Solution is:")
	fmt.Println(grid)
	fmt.Printf("(%v iterations)\n", iterations)
}