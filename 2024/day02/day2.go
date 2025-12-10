package day2

import (
	"fmt"
	"strconv"
	"main/utils"
    "slices"
	"strings"
)

const DAY uint = 2

func Puzzle1() {
    fmt.Println("Day 2, Puzzle 1!")

    reports := get_reports()

    var num_safe_reports uint = 0
    for _, v := range reports {
        if check_report_safe(v) {
            num_safe_reports++
        }
    }

    fmt.Println("Number of safe reports:", num_safe_reports)
}


func Puzzle2() {
    fmt.Println("Day 2, Puzzle 2!")

    reports := get_reports()

    var num_safe_reports uint = 0
    for _, v := range reports {
        if check_report_safe(v) {
            num_safe_reports++
        } else {
            for i := range v {
                // fmt.Println("i:", i)
                // fmt.Println("v:", v, "len(v):", len(v))
                x := slices.Clone(v)
                x = append(x[:i], x[i+1:]...)
                // fmt.Println("x:", x)
                // fmt.Println()
                if check_report_safe(x) {
                    num_safe_reports++
                    break
                }
            }
        }
    }

    fmt.Println("Number of safe reports:", num_safe_reports)
}


func get_reports() [][]uint8 {
    input_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))

    var parsed_reports [][]uint8
    for _, v := range input_lines {
        r := parse_single_report(v)
        parsed_reports = append(parsed_reports, r)
    }

    return parsed_reports
}


func parse_single_report(raw_report string) []uint8 {
    report_split := strings.Split(raw_report, " ")
    var report_uints []uint8
    for _, s := range report_split {
        level, err := strconv.ParseUint(s, 10, 8)
        utils.Check(err)
        report_uints = append(report_uints, uint8(level))
    }
    return report_uints
}


func check_report_safe(report []uint8) bool {
    if len(report) < 3 {
        panic("Report is not long enough")
    }

    report_0 := report[0]
    report_1 := report[1]
    var ascending bool

    if report_0 < report_1 {
        ascending = true
    } else if report_0 > report_1 {
        ascending = false
    } else {
        return false
    }

    for i := 0; i <= len(report)-2; i++ {
        window_0 := report[i]
        window_1 := report[i+1]

        if window_0 == window_1 {
            return false
        } else if window_0 < window_1 {
            if !ascending || window_1 - window_0 > 3 {
                return false
            }
        } else if window_0 > window_1 {
            if ascending || window_0 - window_1 > 3 {
                return false
            }
        } else {
            panic("should be unreachable")
        }
    }

    return true
}
