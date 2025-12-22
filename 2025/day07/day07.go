package day07

import (
    "fmt"
    "main/utils"
    // "strconv"
    "strings"
)

const DAY uint = 7

func Puzzle1() {
    fmt.Printf("Day %d, Puzzle 1!\n", DAY)

    gridlines_strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    // fmt.Println("gridlines_strs:", gridlines_strs)

    grid := make([][]rune, len(gridlines_strs))
    first_line := gridlines_strs[0]
    width := len(first_line)
    grid[0] = make([]rune, width)

    split_count := uint(0)
    for i, line_str := range gridlines_strs {
        line := make([]rune, width)
        if i == 0 {
            for j, cell := range line_str {
                if cell == '.' {
                    line[j] = '.'
                } else if cell == 'S' {
                    line[j] = '|'
                } else {
                    panic(fmt.Sprintf("invalid value found in grid[%d][%d]: %b", i, j, cell))
                }
            }
            grid[0] = line
            continue
        } else {
            for j, cell := range line_str {
                if cell == '.' {
                    line[j] = '.'
                } else if cell == '^' {
                    line[j] = '^'
                } else {
                    panic(fmt.Sprintf("invalid value found in grid[%d][%d]: %b", i, j, cell))
                }
            }
            grid[i] = line
        }

        line_above := grid[i-1]
        // fmt.Println("i:", i)
        // fmt.Println("line_above:  ", string(line_above))
        // fmt.Println("line_before: ", string(line))
        for j, cell := range line {
            if cell == '.' {
                if line_above[j] == '|' {
                    line[j] = '|'
                } else {
                    line[j] = '.'
                }
            } else if cell == '^' {
                if line_above[j] == '|' {
                    if j > 0 && line[j-1] == '.' {
                        line[j-1] = '|'
                    }
                    if j < width - 1 && line[j+1] == '.' {
                        line[j+1] = '|'
                    }
                    split_count++
                }
                line[j] = '^'
            } else if cell == '|' {
                continue
            } else {
                panic(fmt.Sprintf("invalid value found in grid[%d][%d]: %s", i, j, string(cell)))
            }
        }
        // fmt.Println("line_after:  ", string(line), "\n")
    }

    fmt.Println("split_count:", split_count)

    // fmt.Println("grid:")
    // for _, line := range grid{
    //     fmt.Println(string(line))
    // }
}

func Puzzle2() {
    fmt.Printf("\nDay %d, Puzzle 2!\n", DAY)

    gridlines_strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    // fmt.Println("gridlines_strs:", gridlines_strs)

    first_line := gridlines_strs[0]
    width := len(first_line)
    path_counts := make([]uint, width)
    start_index := strings.Index(first_line, "S")
    path_counts[start_index] += 1

    for _, line_str := range gridlines_strs[1:] {
        for i, char := range line_str {
            if char == '^' {
                path_counts[i - 1] += path_counts[i]
                path_counts[i + 1] += path_counts[i]
                path_counts[i] = 0
            }
        }
    }

    // fmt.Println("path_counts:", path_counts)
    var total_paths uint
    for _, i := range path_counts {
        total_paths += i
    }
    fmt.Println("total_paths:", total_paths)
}

