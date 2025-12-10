package day6

import (
	"fmt"
	"slices"
	"strings"
	"main/utils"
)

const DAY uint = 6
type check_op func(int, int, byte, []string) bool
var NUM_ROWS int
var NUM_COLS int
var MAX_ROWS_IDX int
var MAX_COLS_IDX int
var PREV_TURNS = []turn{}

type turn struct {
    direction byte
    idx int
}

// array[width * row + col]


func Puzzle1() {
    fmt.Println("Day 6, Puzzle 1!")

    PREV_TURNS = PREV_TURNS[:0]
    gridmap := get_gridmap()
    // fmt.Println(fmt.Sprintf("%#v", gridmap))

    init_idx, guard_direction := find_guard(gridmap)
    gridmap = move_guard_part1(init_idx, guard_direction, gridmap)

    var num_positions uint = 0
    for _, c := range gridmap {
        if c == 'X' {
            num_positions++
        }
    }

    fmt.Println("Number of distinct guard positions:", num_positions)
}


func find_guard(gridmap []byte) (int, byte) {
    for i, char := range gridmap {
        if char == '^' {
            return i, '^'
        } else if char == '>' {
            return i, '>'
        } else if char == 'v' {
            return i, 'v'
        } else if char == '<' {
            return i, '<'
        }
    }

    panic("should be unreachable")
}


func move_guard_part1(init_idx int, guard_direction byte, gridmap []byte) []byte {

    var complete_grid []byte
    var idx int

    switch guard_direction {
    case '^':
        idx, complete_grid = move_up(init_idx, gridmap)
    case '>':
        idx, complete_grid = move_right(init_idx, gridmap)
    case 'v':
        idx, complete_grid = move_down(init_idx, gridmap)
    case '<':
        idx, complete_grid = move_left(init_idx, gridmap)
    }

    if idx != -1 {
        fmt.Println("idx:", idx)
        panic("grid not actually complete")
    }

    return complete_grid
}


func Puzzle2() {
    fmt.Println("Day 6, Puzzle 2!")

    PREV_TURNS = PREV_TURNS[:0]
    gridmap := get_gridmap()
    init_idx, guard_direction := find_guard(gridmap)
    gridmap = move_guard_part1(init_idx, guard_direction, gridmap)
    gridmap[init_idx] = guard_direction
    fmt.Println("got part1")
    // fmt.Println(fmt.Sprintf("%#v", gridmap))

    var num_positions uint = 0
    for i, c := range gridmap {
        PREV_TURNS = PREV_TURNS[:0]
        if c == 'X' {
            grid_copy := make([]byte, len(gridmap))
            copy(grid_copy, gridmap)
            grid_copy[i] = '#'
            if move_guard_part2(init_idx, guard_direction, grid_copy) {
                num_positions++
                // fmt.Println("found loop")
            }
        }
    }

    fmt.Println("Number of distinct blocker positions:", num_positions)
}


func move_guard_part2(init_idx int, guard_direction byte, gridmap []byte) bool {

    var idx int

    switch guard_direction {
    case '^':
        idx, _ = move_up(init_idx, gridmap)
    case '>':
        idx, _ = move_right(init_idx, gridmap)
    case 'v':
        idx, _ = move_down(init_idx, gridmap)
    case '<':
        idx, _ = move_left(init_idx, gridmap)
    }

    if idx >= 0 {
        fmt.Println("idx:", idx)
        panic("loop search failed")
    } else if idx == -1 {
        return false
    } else if idx == -2 {
        return true
    }

    panic("should be unreachable")
}


func get_row_and_col(idx int) (int, int) {
    row := idx / NUM_COLS
    col := idx - (row * NUM_COLS)
    return row, col
}


func handle_prev_turns(direction byte, idx int) bool {
    t := turn{direction: direction, idx: idx}

    if slices.Contains(PREV_TURNS, t) {
        return true
    }
    PREV_TURNS = append(PREV_TURNS, t)

    return false
}


func move_up(idx int, gridmap []byte) (int, []byte) {
    // fmt.Println("Up")
    // fmt.Println("idx:", idx)

    if idx < NUM_COLS {
        gridmap[idx] = 'X'
        return -1, gridmap
    }

    idx_row_above := idx - NUM_COLS
    if gridmap[idx_row_above] == '#' {
        if handle_prev_turns(0, idx) {
            return -2, gridmap
        }
        gridmap[idx] = '>'
        return move_right(idx, gridmap)
    }

    gridmap[idx] = 'X'
    return move_up(idx_row_above, gridmap)
}


func move_right(idx int, gridmap []byte) (int, []byte) {
    // fmt.Println("Right")
    // fmt.Println("idx:", idx)

    _, col := get_row_and_col(idx)

    if col == MAX_COLS_IDX {
        gridmap[idx] = 'X'
        return -1, gridmap
    }

    idx_row_right := idx + 1
    if gridmap[idx_row_right] == '#' {
        if handle_prev_turns(1, idx) {
            return -2, gridmap
        }
        gridmap[idx] = 'v'
        return move_down(idx, gridmap)
    }

    gridmap[idx] = 'X'
    return move_right(idx_row_right, gridmap)
}


func move_down(idx int, gridmap []byte) (int, []byte) {
    // fmt.Println("Down")
    // fmt.Println("idx:", idx)

    if idx > NUM_COLS * MAX_ROWS_IDX {
        gridmap[idx] = 'X'
        return -1, gridmap
    }

    idx_row_below := idx + NUM_COLS
    if gridmap[idx_row_below] == '#' {
        if handle_prev_turns(2, idx) {
            return -2, gridmap
        }
        gridmap[idx] = '<'
        return move_left(idx, gridmap)
    }

    gridmap[idx] = 'X'
    return move_down(idx_row_below, gridmap)
}


func move_left(idx int, gridmap []byte) (int, []byte) {
    // fmt.Println("Left")
    // fmt.Println("idx:", idx)

    _, col := get_row_and_col(idx)

    if col == 0 {
        gridmap[idx] = 'X'
        return -1, gridmap
    }

    idx_row_left := idx - 1
    if gridmap[idx_row_left] == '#' {
        if handle_prev_turns(3, idx) {
            return -2, gridmap
        }
        gridmap[idx] = '^'
        return move_up(idx, gridmap)
    }

    gridmap[idx] = 'X'
    return move_left(idx_row_left, gridmap)
}


func get_gridmap() []byte {

    input_str := utils.ReadInputToString(utils.Get_Input_Filepath(DAY))
    split := strings.Split(input_str, "\n")

    NUM_ROWS = len(split)
    NUM_COLS = len(split[0])
    MAX_ROWS_IDX = NUM_ROWS - 1
    MAX_COLS_IDX = NUM_COLS - 1
    
    gridmap := make([]byte, NUM_COLS * NUM_ROWS)
    for i, r := range split {
        // fmt.Println(fmt.Sprintf("%#v", r))
        for j, c := range r {
            gridmap[(NUM_COLS * i) + j] = byte(c)
        }
    }

    return gridmap
}


