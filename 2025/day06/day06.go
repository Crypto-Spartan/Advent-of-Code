package day06

import (
    "fmt"
    "main/utils"
    "strconv"
    "strings"
)

const DAY uint = 6

func Puzzle1() {
    fmt.Printf("Day %d, Puzzle 1!\n", DAY)

    homework_problems_lines_strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    fmt.Println("homework_problems_lines_strs:", homework_problems_lines_strs)

    num_int_lines := len(homework_problems_lines_strs) - 1
    fmt.Println("num_int_lines:", num_int_lines)
    homework_problems_lines_ints := make([][]int, num_int_lines)
    line_length := len(strings.Fields(homework_problems_lines_strs[0]))
    for i, line := range homework_problems_lines_strs[:num_int_lines] {
        line_split := strings.Fields(line)
        if len(line_split) != line_length {
            panic("lines of different lengths")
        }

        line_ints := make([]int, line_length)
        for j, n_str := range line_split {
            n_int, err := strconv.Atoi(n_str)
            if err != nil {
                fmt.Println("Error converting string to int:", err)
                return
            }
            line_ints[j] = n_int
        }

        homework_problems_lines_ints[i] = line_ints
    }
    fmt.Println("homework_problems_lines_ints:", homework_problems_lines_ints)

    homework_total := uint(0)
    ops_strs := strings.Fields(homework_problems_lines_strs[num_int_lines])
    fmt.Println("ops_strs:", ops_strs)
    for i, op_str := range ops_strs {
        col_total := uint(0)
        if op_str == "+" {
            for j := 0; j < num_int_lines; j++ {
                col_total += uint(homework_problems_lines_ints[j][i])
            }
        } else if op_str == "*" {
            col_total += uint(1)
            for j := 0; j < num_int_lines; j++ {
                col_total *= uint(homework_problems_lines_ints[j][i])
            }
        } else {
            panic(fmt.Sprintf("Unknown operator: %s", op_str))
        }

        homework_total += col_total
    }

    fmt.Println("math homework total:", homework_total)
}

func Puzzle2() {
    fmt.Printf("\nDay %d, Puzzle 2!\n", DAY)

    homework_problems_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    fmt.Println("homework_problems_lines:", homework_problems_lines)
    width := len(homework_problems_lines[0])
    // fmt.Println("width:", width)
    length := len(homework_problems_lines)

    columns := make([][]rune, width)
    for _, line := range homework_problems_lines {
        for i, char := range line {
            columns[i] = append(columns[i], char)
        }
    }
    // fmt.Println("columns:", columns)

    total := uint(0)
    new_problem := true
    var current_op byte
    var current_total uint
    for _, col_runes := range columns {
        col_str := string(col_runes)
        if len(strings.ReplaceAll(col_str, " ", "")) == 0 {
            total += current_total
            new_problem = true
            continue
        }

        if new_problem {
            current_op = col_str[length - 1]
            if current_op == '+' {
                current_total = 0
            } else if current_op == '*' {
                current_total = 1
            } else {
               panic(fmt.Sprintf("Unknown operator: %b", current_op))
            }
            new_problem = false
        }

        int_str_no_spaces := strings.ReplaceAll(col_str[:length - 1], " ", "")
        current_int, err := strconv.Atoi(int_str_no_spaces)
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        if current_op == '+' {
            current_total += uint(current_int)
        } else if current_op == '*' {
            current_total *= uint(current_int)
        } else {
            panic(fmt.Sprintf("Unknown operator: %b", current_op))
        }
    }
    total += current_total
    fmt.Println("part 2 total:", total)
}

