package day4

import (
    "fmt"
    //"strconv"
    "main/utils"
    //"slices"
    // "strconv"
    // "strings"
)

const DAY uint = 4
type check_op func(int, int, byte, []string) bool
var NUM_ROWS int
var NUM_COLS int
var MAX_ROWS_IDX int
var MAX_COLS_IDX int


func Puzzle1() {
    fmt.Println("Day 4, Puzzle 1!")

    wordsearch := get_wordsearch()
    // fmt.Println(wordsearch)
    NUM_ROWS = len(wordsearch)
    NUM_COLS = len(wordsearch[0])
    MAX_ROWS_IDX = NUM_ROWS - 1
    MAX_COLS_IDX = NUM_COLS - 1

    var total_num_xmas uint = 0
    for ri, rv := range wordsearch {
        // fmt.Println("ri:", ri, "rv:", rv)

        if len(rv) != NUM_COLS {
            panic(fmt.Sprintf("row %d doesn't have %d cols", ri, NUM_COLS))
        }

        for ci, cv := range rv {
            if cv == 'X' {
                total_num_xmas += find_xmas_puz1(ri, ci, wordsearch)
            }
        }
    }

    fmt.Println("Number of XMASs in wordsearch:", total_num_xmas)
}


func get_wordsearch() []string {
    input_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    return input_lines
}


func find_xmas_puz1(row, col int, ws []string) uint {
    check_ops_arr := [8]check_op{check_upleft, check_up, check_upright, check_left, check_right, check_downleft, check_down, check_downright}
    var num_xmas uint = 0

    for _, fn := range check_ops_arr {
        if fn(row, col, 'X', ws) {
            num_xmas++
        }
    }

    return num_xmas
}


func get_next_char(char byte) byte {
    switch char {
    case 'X':
        return 'M'
    case 'M':
        return 'A'
    case 'A':
        return 'S'
    default:
        panic("should be unreachable")
    }
}


func check_upleft(row, col int, char byte, ws []string) bool {
    if char == 'S' {
        return true
    }
    if row == 0 || col == 0 {
        return false
    }
    next_r := row - 1
    next_c := col - 1

    next_char := get_next_char(char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_upleft(next_r, next_c, next_char, ws)
}


func check_up(row, col int, char byte, ws []string) bool {
    if char == 'S' {
        return true
    }
    if row == 0 {
        return false
    }
    next_r := row - 1
    next_c := col

    next_char := get_next_char(char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_up(next_r, next_c, next_char, ws)
}


func check_upright(row, col int, char byte, ws []string) bool {
    if char == 'S' {
        return true
    }
    if row == 0 || col == MAX_COLS_IDX {
        return false
    }
    next_r := row - 1
    next_c := col + 1

    next_char := get_next_char(char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_upright(next_r, next_c, next_char, ws)
}


func check_left(row, col int, char byte, ws []string) bool {
    if char == 'S' {
        return true
    }
    if col == 0 {
        return false
    }
    next_r := row
    next_c := col - 1

    next_char := get_next_char(char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_left(next_r, next_c, next_char, ws)
}


func check_right(row, col int, curr_char byte, ws []string) bool {
    if curr_char == 'S' {
        return true
    }
    if col == MAX_COLS_IDX {
        return false
    }
    next_r := row
    next_c := col + 1

    next_char := get_next_char(curr_char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_right(next_r, next_c, next_char, ws)
}


func check_downleft(row, col int, char byte, ws []string) bool {
    if char == 'S' {
        return true
    }
    if row == MAX_ROWS_IDX || col == 0 {
        return false
    }
    next_r := row + 1
    next_c := col - 1

    next_char := get_next_char(char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_downleft(next_r, next_c, next_char, ws)
}


func check_down(row, col int, char byte, ws []string) bool {
    if char == 'S' {
        return true
    }
    if row == MAX_ROWS_IDX {
        return false
    }
    next_r := row + 1
    next_c := col

    next_char := get_next_char(char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_down(next_r, next_c, next_char, ws)
}


func check_downright(row, col int, char byte, ws []string) bool {
    if char == 'S' {
        return true
    }
    if row == MAX_ROWS_IDX || col == MAX_COLS_IDX {
        return false
    }
    next_r := row + 1
    next_c := col + 1

    next_char := get_next_char(char)
    if ws[next_r][next_c] != next_char {
        return false
    }

    return check_downright(next_r, next_c, next_char, ws)
}



func Puzzle2() {
    fmt.Println("Day 4, Puzzle 2!")

    wordsearch := get_wordsearch()
    NUM_ROWS = len(wordsearch)
    NUM_COLS = len(wordsearch[0])
    MAX_ROWS_IDX = NUM_ROWS - 1
    MAX_COLS_IDX = NUM_COLS - 1

    var total_num_xmas uint = 0
    for ri := 1; ri < MAX_ROWS_IDX; ri++ {
        // fmt.Println("ri:", ri, "rv:", rv)

        rv := wordsearch[ri]
        
        if len(rv) != NUM_COLS {
            panic(fmt.Sprintf("row %d doesn't have %d cols", ri, NUM_COLS))
        }

        for ci := 1; ci < MAX_COLS_IDX; ci++ {
            cv := rv[ci]
            if cv == 'A' && check_xmas_puz2(ri, ci, wordsearch) {
                total_num_xmas++
            }
        }
    }

    fmt.Println("Number of X-MASs in wordsearch:", total_num_xmas)
}


func check_xmas_puz2(r, c int, ws []string) bool {

    first_diag := false
    if (ws[r-1][c-1] == 'M' && ws[r+1][c+1] == 'S') ||
    (ws[r-1][c-1] == 'S' && ws[r+1][c+1] == 'M') {
        first_diag = true
    }

    if first_diag &&
    ((ws[r-1][c+1] == 'S' && ws[r+1][c-1] == 'M') ||
    (ws[r-1][c+1] == 'M' && ws[r+1][c-1] == 'S')) {
        return true
    }

    return false
}
