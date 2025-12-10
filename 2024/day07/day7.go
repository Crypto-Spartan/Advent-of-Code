package day7

import (
    "fmt"
    "main/utils"
    // "math"
    // "slices"
    "strconv"
    "strings"
    "time"
)

const DAY uint = 7
type calibration struct {
    test_val uint
    eqn_nums []uint
}
var ADD_MULT_OPS = make(map[uint8][][]uint8)
var ADD_MULT_CONCAT_OPS = make(map[uint8][][]uint8)


func Puzzle1() {
    fmt.Println("Day 7, Puzzle 1!")

    calibrations := get_calibrations()
    // fmt.Println(fmt.Sprintf("%#v", calibrations))
    // fmt.Println("calibrations:", calibrations)

    var sum_vals uint = 0
    for _, c := range calibrations {
        // fmt.Println("calibration:", c)
        if test_calibration_part1(c) {
            sum_vals += c.test_val
        }
        // fmt.Println()
    }

    fmt.Println("Sum of test_vals from valid equations", sum_vals)
}


func test_calibration_part1(c calibration) bool {
    ops_list := get_add_mult_ops(uint8(len(c.eqn_nums) - 1))
    // fmt.Println("ops_list:", ops_list)
    if len(ops_list) == 0 {
        panic(fmt.Sprintf("no operations found for %v", c))
    }

    for _, ops := range ops_list {
        // fmt.Println("ops:", ops)
        var temp_total uint = c.eqn_nums[0]
        for i, o := range ops {
            switch o {
            case 0:
                temp_total += c.eqn_nums[i+1]
            case 1:
                temp_total = temp_total * c.eqn_nums[i+1]
            default:
                panic("should be unreachable")
            }
        }

        if temp_total == c.test_val {
            // fmt.Println("TRUE")
            return true
        }
    }

    return false
}


func Puzzle2() {
    fmt.Println("Day 7, Puzzle 2!")

    calibrations := get_calibrations()
    // fmt.Println(fmt.Sprintf("%#v", calibrations))
    // fmt.Println("calibrations:", calibrations)

    start := time.Now()

    var sum_vals uint = 0
    for _, c := range calibrations {
        // fmt.Println("calibration:", c)
        if test_calibration_part2(c) {
            sum_vals += c.test_val
        }
        // fmt.Println()
    }

    fmt.Println("Sum of test_vals from valid equations", sum_vals)
    fmt.Printf("%s took %v\n", "Puzzle2", time.Since(start))
}


func Puzzle2_par() {
    fmt.Println("Day 7, Puzzle 2! (parallel)")

    calibrations := get_calibrations()
    // fmt.Println(fmt.Sprintf("%#v", calibrations))
    // fmt.Println("calibrations:", calibrations)

    start := time.Now()

    results := make(chan uint)
    for _, c := range calibrations {
        // fmt.Println("calibration:", c)
        go func(c calibration) {
            if test_calibration_part2(c) {
                results <- c.test_val
            } else {
                results <- 0
            }
        }(c)
        // fmt.Println()
    }

    var sum_vals uint = 0
    for i := 0; i < len(calibrations); i++ {
        sum_vals += <-results
    }

    fmt.Println("Sum of test_vals from valid equations", sum_vals)
    fmt.Printf("%s took %v\n", "Puzzle2_par", time.Since(start))
}


func test_calibration_part2(c calibration) bool {
    ops_list := get_add_mult_concat_ops(uint8(len(c.eqn_nums) - 1))
    // fmt.Println("len(ops_list):", len(ops_list))
    // fmt.Println("ops_list:", ops_list)
    if len(ops_list) == 0 {
        panic(fmt.Sprintf("no operations found for %v", c))
    }

    for _, ops := range ops_list {
        // fmt.Println("ops:", ops)
        var temp_total uint = c.eqn_nums[0]
        for i, op := range ops {
            switch op {
            case 0:
                temp_total += c.eqn_nums[i+1]
            case 1:
                temp_total = temp_total * c.eqn_nums[i+1]
            case 2:
                s_concat := fmt.Sprintf("%d%d", temp_total, c.eqn_nums[i+1])
                v, err := strconv.ParseUint(s_concat, 10, 0)
                utils.Check(err)
                temp_total = uint(v)
            default:
                panic("should be unreachable")
            }
        }

        if temp_total == c.test_val {
            // fmt.Println("TRUE")
            return true
        }
    }

    return false
}


func get_add_mult_concat_ops(n_ops uint8) [][]uint8 {
    if n_ops == 0 {
        return nil
    }
    if ops, ok := ADD_MULT_CONCAT_OPS[n_ops]; ok {
        return ops
    }

    var choices = [3]uint8{0, 1, 2}
    ops := [][]uint8{}
    if n_ops == 1 {
        for _, c := range choices {
            ops = append(ops, []uint8{c})
        }
        // fmt.Println("ops:", ops)
        return ops
    }

    smaller_product := get_add_mult_concat_ops(n_ops - 1)
    for _, c := range choices {
        for _, x := range smaller_product {
            temp := make([]uint8, len(x))
            copy(temp, x)
            temp = append(temp, c)
            temp = temp[:len(x)+1]
            ops = append(ops, temp)
        }
    }
    // fmt.Println("ops:", ops)

    ADD_MULT_CONCAT_OPS[n_ops] = ops
    return ops
}


func get_add_mult_ops(n_ops uint8) [][]uint8 {
    if ops, ok := ADD_MULT_OPS[n_ops]; ok {
        return ops
    }

    // fmt.Println("n_ops:", n_ops)
    ops := [][]uint8{}

    q := 1 << n_ops
    // fmt.Println("q:", q)
    for i := 0; i < q; i++ {
        // fmt.Println("i:", i)
        x := make([]uint8, n_ops)
        for xi := 0; xi < len(x); xi++ {
            if (i >> xi) & 0x1 == 0x1 {
                x[xi] = 0
            } else {
                x[xi] = 1
            }
        }
        // fmt.Println("x:", x)
        ops = append(ops, x)
    }

    ADD_MULT_OPS[n_ops] = ops
    return ops
}


func get_calibrations() []calibration {
    input_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))

    all_calibrations := []calibration{}
    for _, s := range input_lines {

        colon_idx := strings.IndexByte(s, ':')
        if colon_idx == -1 {
            fmt.Println("s:", s)
            panic("colon seperator not found")
        }

        t_val, err := strconv.ParseUint(s[:colon_idx], 10, 0)
        utils.Check(err)

        eqn_nums_as_strs := strings.Split(s[colon_idx + 1:], " ")
        // fmt.Println("eqn_nums_as_strs:", fmt.Sprintf("%#v", eqn_nums_as_strs))
        eqn_nums := []uint{}
        for _, ns := range eqn_nums_as_strs {
            if ns == "" {
                continue
            }
            // fmt.Println("ns:", fmt.Sprintf("%#v", ns))

            nui, err := strconv.ParseUint(ns, 10, 0)
            utils.Check(err)

            eqn_nums = append(eqn_nums, uint(nui))
        }

        c := calibration{test_val: uint(t_val), eqn_nums: eqn_nums}
        all_calibrations = append(all_calibrations, c)
    }

    return all_calibrations
}
