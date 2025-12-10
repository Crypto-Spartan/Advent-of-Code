package day8

import (
	"fmt"
	"main/utils"
	// "slices"
	// "strings"
)

const DAY uint = 8

var NUM_ROWS uint
var NUM_COLS uint
var MAX_ROWS_IDX uint
var MAX_COLS_IDX uint

type coordinate struct {
	x uint
	y uint
}
type antenna_coords_T = map[byte][]coordinate
type antinodes_T = map[coordinate]struct{}

// array[width * row + col]

func Puzzle1() {
	fmt.Println("Day 8, Puzzle 1!")

	gridmap := get_gridmap()
	antenna_coords := get_antenna_coords(gridmap)
	// fmt.Println("antenna_coords:", antenna_coords)
	antinodes := get_antinodes_pt1(antenna_coords)
	// fmt.Println("antinodes:", antinodes)

	// for c := range antinodes {
	// 	gridmap[NUM_COLS * c.y + c.x] = '#'
	// }

	// fmt.Println("gridmap w/ antinodes:")
	// for row := uint(0); row < NUM_ROWS; row++ {
	// 	i := row * NUM_COLS
	// 	j := i + NUM_COLS
	// 	chunk := gridmap[i:j]
	//
	// 	fmt.Println(string(chunk))
	// }

	fmt.Println("Number of unique antinodes:", len(antinodes))
}


func Puzzle2() {
	fmt.Println("Day 8, Puzzle 2!")

	gridmap := get_gridmap()
	antenna_coords := get_antenna_coords(gridmap)
	// fmt.Println("antenna_coords:", antenna_coords)
	antinodes := get_antinodes_pt2(antenna_coords)
	// fmt.Println("antinodes:", antinodes)

	// for c := range antinodes {
	// 	gridmap[NUM_COLS * c.y + c.x] = '#'
	// }

	// fmt.Println("gridmap w/ antinodes:")
	// for row := uint(0); row < NUM_ROWS; row++ {
	// 	i := row * NUM_COLS
	// 	j := i + NUM_COLS
	// 	chunk := gridmap[i:j]
	//
	// 	fmt.Println(string(chunk))
	// }

	fmt.Println("Number of unique antinodes:", len(antinodes))
}


func get_antenna_coords(gridmap []byte) antenna_coords_T {
	// size := NUM_COLS
	// gridmap_len := uint(len(gridmap))
	antenna_coords := antenna_coords_T{}

	for row := uint(0); row < NUM_ROWS; row++ {
		i := row * NUM_COLS
		j := i + NUM_COLS
		chunk := gridmap[i:j]

		for col, b := range chunk {
			if b != '.' {
				c := coordinate{x: uint(col), y: row}
				antenna_coords[b] = append(antenna_coords[b], c)
			}
		}
	}

	return antenna_coords
}


func get_antinodes_pt1(antenna_coords antenna_coords_T) antinodes_T {
	antinodes := antinodes_T{}

	for _, coords := range antenna_coords {
		for i1, c1 := range coords {
			for i2, c2 := range coords {
				if i1 == i2 {
					continue
				}

				x_diff := c1.x - c2.x
				y_diff := c1.y - c2.y

				an_x := c2.x - x_diff
				an_y := c2.y - y_diff

				if an_x >= 0 && an_x < NUM_COLS &&
				an_y >= 0 && an_y < NUM_ROWS {
					antinodes[coordinate{x: an_x, y: an_y}] = struct{}{}
				}
			}
		}
	}

	return antinodes
}


func get_antinodes_pt2(antenna_coords antenna_coords_T) antinodes_T {
	antinodes := antinodes_T{}

	for _, coords := range antenna_coords {
		for i1, c1 := range coords {
			for i2, c2 := range coords {
				if i1 == i2 {
					continue
				}

				x_diff := c1.x - c2.x
				y_diff := c1.y - c2.y

				an_x := c2.x + x_diff
				an_y := c2.y + y_diff

				for an_x >= 0 && an_x < NUM_COLS &&
				an_y >= 0 && an_y < NUM_ROWS {
					antinodes[coordinate{x: an_x, y: an_y}] = struct{}{}
					an_x = an_x + x_diff
					an_y = an_y + y_diff
				}
			}
		}
	}

	return antinodes
}


func get_gridmap() []byte {

	input_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
	// for _, s := range input_lines {
	// 	fmt.Println(s)
	// }

	NUM_ROWS = uint(len(input_lines))
	NUM_COLS = uint(len(input_lines[0]))
	MAX_ROWS_IDX = NUM_ROWS - 1
	MAX_COLS_IDX = NUM_COLS - 1

	gridmap := make([]byte, NUM_COLS*NUM_ROWS)
	for i, r := range input_lines {
		// fmt.Println(fmt.Sprintf("%#v", r))
		for j, c := range r {
			gridmap[(NUM_COLS*uint(i))+uint(j)] = byte(c)
		}
	}

	return gridmap
}
