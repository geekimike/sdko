package main

import (
	"fmt"
	"strconv"
)

type cell struct  {
	value int
	given bool
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

type column []cell

func (c column) String() string {
	result := ""

	for _, cell := range c {
		result += fmt.Sprintf("%d ", cell.value)
	}
	result += "\n"

	return result
}

type square []cell

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
	
}

func solve(g grid) grid {
	result := g

	return result
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
		result[i] = cell{value: d, given: c!='0'}
	}

	return result
}

func buildTestGrid1() grid {
	return buildFromString("000090000065000380030000050000209000709603208000401000607908102040506090092030560")
}

func main() {
	fmt.Println()

	grid := buildTestGrid1()

	fmt.Println("Solving:")
	fmt.Println(grid)

	solution := solve(grid)
	
	fmt.Println("Solution is:")
	fmt.Println(solution)
}