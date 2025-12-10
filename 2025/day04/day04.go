package day04

import (
    "fmt"
    "main/utils"
    // "strconv"
    "strings"
)

const DAY uint = 4

func get_adj_rolls_count(row_idx, col_idx int, gridmap [][]bool) uint {
    adj_rolls_count := uint(0)
    if row_idx == 0 {
        adj_rolls_count += get_rolls_count_nextto(col_idx, gridmap[row_idx])
        adj_rolls_count += get_rolls_count_abv_bel(get_cells_to_check(col_idx, gridmap[row_idx + 1]))
    } else if row_idx == len(gridmap) - 1 {
        adj_rolls_count += get_rolls_count_abv_bel(get_cells_to_check(col_idx, gridmap[row_idx - 1]))
        adj_rolls_count += get_rolls_count_nextto(col_idx, gridmap[row_idx])
    } else {
        adj_rolls_count += get_rolls_count_abv_bel(get_cells_to_check(col_idx, gridmap[row_idx - 1]))
        adj_rolls_count += get_rolls_count_nextto(col_idx, gridmap[row_idx])
        adj_rolls_count += get_rolls_count_abv_bel(get_cells_to_check(col_idx, gridmap[row_idx + 1]))
    }
    return adj_rolls_count
}

func get_rolls_count_abv_bel(cells_to_check []bool) uint {
    rolls_count := uint(0)
    for _, b := range cells_to_check {
        if b == true {
            rolls_count++
        }
    }
    return rolls_count
}

func get_rolls_count_nextto(col_idx int, row []bool) uint {
    if col_idx == 0 {
        if row[col_idx + 1] {
            return uint(1)
        } else {
            return uint(0)
        }
    } else if col_idx + 1 == len(row) {
        if row[col_idx - 1] {
            return uint(1)
        } else {
            return uint(0)
        }
    } else {
        rolls_count := uint(0)
        if row[col_idx - 1] {
            rolls_count++
        }
        if row[col_idx + 1] {
            rolls_count++
        }
        return rolls_count
    }
}

func get_cells_to_check(col_idx int, row []bool) []bool {
    col_idx_start := col_idx - 1
    if col_idx_start < 0 {
        col_idx_start = 0
    }

    var cells_to_check []bool
    col_idx_end := col_idx + 2
    if col_idx_end >= len(row) {
        cells_to_check = row[col_idx_start:]
    } else {
        cells_to_check = row[col_idx_start:col_idx_end]
    }

    if len(cells_to_check) == 0 {
        panic("cells_to_check is empty")
    } else if len(cells_to_check) > 3 {
        panic("cells_to_check is too long")
    }

    return cells_to_check
}

func Puzzle1() {
    fmt.Println(fmt.Sprintf("Day %d, Puzzle 1!", DAY))

    gridmap_strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    row_count := len(gridmap_strs)
    col_count := len(gridmap_strs[0])
    // gridmap_accessible := make([][]string, row_count)
    gridmap := make([][]bool, row_count)
    for row_idx, row_str := range gridmap_strs {
        row_strs := strings.Split(strings.TrimSpace(row_str), "")
        // gridmap_accessible[row_idx] = row_strs
        row_bools := make([]bool, col_count)
        for col_idx, s := range row_strs {
            if s == "." {
                row_bools[col_idx] = false
            } else if s == "@" {
                row_bools[col_idx] = true
            } else {
                panic(fmt.Sprintf("invalid grid cell value: %s", s))
            }
        }
        gridmap[row_idx] = row_bools
    }
    // fmt.Println("gridmap:", gridmap)

    accessible_rolls_count := uint(0)
    for row_idx := 0; row_idx < row_count; row_idx++ {
        for col_idx := 0; col_idx < col_count; col_idx++ {
            if gridmap[row_idx][col_idx] == true {
                adj_rolls_count := get_adj_rolls_count(row_idx, col_idx, gridmap)
                if adj_rolls_count < 4 {
                    accessible_rolls_count++
                    // gridmap_accessible[row_idx][col_idx] = "x"
                }
            }
        }
        // fmt.Println(strings.Join(gridmap_accessible[row_idx], ""))
    }
    // fmt.Println()

    fmt.Println("# of accessible rolls:", accessible_rolls_count)
}

func Puzzle2() {
    fmt.Println(fmt.Sprintf("Day %d, Puzzle 2!", DAY))

    gridmap_strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    row_count := len(gridmap_strs)
    col_count := len(gridmap_strs[0])
    // gridmap_accessible := make([][]string, row_count)
    gridmap := make([][]bool, row_count)
    for row_idx, row_str := range gridmap_strs {
        row_strs := strings.Split(strings.TrimSpace(row_str), "")
        // gridmap_accessible[row_idx] = row_strs
        row_bools := make([]bool, col_count)
        for col_idx, s := range row_strs {
            if s == "." {
                row_bools[col_idx] = false
            } else if s == "@" {
                row_bools[col_idx] = true
            } else {
                panic(fmt.Sprintf("invalid grid cell value: %s", s))
            }
        }
        gridmap[row_idx] = row_bools
    }
    // fmt.Println("gridmap:", gridmap)

    all_accessible_rolls_count := uint(0)
    for {
        accessible_rolls_count := uint(0)
        for row_idx := 0; row_idx < row_count; row_idx++ {
            for col_idx := 0; col_idx < col_count; col_idx++ {
                if gridmap[row_idx][col_idx] {
                    adj_rolls_count := get_adj_rolls_count(row_idx, col_idx, gridmap)
                    if adj_rolls_count < 4 {
                        accessible_rolls_count++
                        gridmap[row_idx][col_idx] = false
                    }
                }
            }
        }

        if accessible_rolls_count > 0 {
            all_accessible_rolls_count += accessible_rolls_count
        } else {
            break
        }
    }

    fmt.Println("# of all accessible rolls:", all_accessible_rolls_count)
}

