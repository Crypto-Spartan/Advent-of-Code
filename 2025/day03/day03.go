package day03

import (
    "fmt"
    "main/utils"
    "strconv"
    "strings"
)

const DAY uint = 3

func get_slice_max(slc []uint) (uint, uint) {
    max := uint(0)
    idx := uint(0)
    for i, n := range slc {
        if n > max {
            max = n
            idx = uint(i)
        }
        if max == 9 {
            break
        }
    }
    return max, idx
}

func Puzzle1() {
    fmt.Println("Day 3, Puzzle 1!")

    battery_banks := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    for i := 0; i < len(battery_banks); i++ {
        battery_banks[i] = strings.TrimSpace(battery_banks[i])
    }
    fmt.Println("battery_banks:", battery_banks)

    bank_len := len(battery_banks[0])
    var largest_joltage_sum uint = 0
    for bank_idx, bank_str := range battery_banks {
        bank_str_slc := strings.Split(bank_str, "")
        bank_uint_slc := make([]uint, bank_len)
        for i, j_str := range bank_str_slc {
            j_int, err := strconv.Atoi(j_str)
            if err != nil {
                panic(fmt.Sprintf("Error converting string to int: %s", err))
            }
            bank_uint_slc[i] = uint(j_int)
        }

        bank_max_1, bank_max_1_idx := get_slice_max(bank_uint_slc[:bank_len - 1])
        bank_max_2, _ := get_slice_max(bank_uint_slc[bank_max_1_idx + 1:])

        bank_joltage := (bank_max_1 * 10) + bank_max_2
        fmt.Println("Bank", bank_idx, "joltage:", bank_joltage)
        largest_joltage_sum += bank_joltage
    }

    fmt.Println("Sum of largest joltages:", largest_joltage_sum)
}

const PUZZLE2_BATTERIES_NEEDED uint = 12

func UIntPow(n, e uint) uint {
    if e == 0 {
        return 1
    }

    if e == 1 {
        return n
    }

    result := n
    for i := uint(2); i <= e; i++ {
        result *= n
    }
    return result
}

func Puzzle2() {
    fmt.Println("Day 3, Puzzle 2!")

    battery_banks := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    for i := 0; i < len(battery_banks); i++ {
        battery_banks[i] = strings.TrimSpace(battery_banks[i])
    }
    fmt.Println("battery_banks:", battery_banks)

    bank_len := uint(len(battery_banks[0]))
    fmt.Println("bank_len:", bank_len)
    var largest_joltage_sum uint = 0
    for bank_idx, bank_str := range battery_banks {
        bank_str_slc := strings.Split(bank_str, "")
        bank_uint_slc := make([]uint, bank_len)
        for i, j_str := range bank_str_slc {
            j_int, err := strconv.Atoi(j_str)
            if err != nil {
                panic(fmt.Sprintf("Error converting string to int: %s", err))
            }
            bank_uint_slc[i] = uint(j_int)
        }

        batteries_needed := uint(PUZZLE2_BATTERIES_NEEDED)
        var batteries_chosen_joltage [PUZZLE2_BATTERIES_NEEDED]uint

        // fmt.Println("bank_uint_slc", bank_uint_slc)
        // bank_max_1, bank_max_1_idx := get_slice_max(bank_uint_slc[:uint(bank_len) - batteries_needed])
        // fmt.Println("bank_max_1:", bank_max_1, ", bank_max_1_idx:", bank_max_1_idx)
        // batteries_chosen_joltage[0] = bank_max_1
        // batteries_needed -= 1
        bank_max_n_idx := uint(0)
        for battery_idx := uint(0); battery_idx < PUZZLE2_BATTERIES_NEEDED; battery_idx++ {
            slc_end_bounds := bank_len - batteries_needed + 1
            // fmt.Println("slc_end_bounds:", slc_end_bounds)
            var slc_to_search []uint
            if slc_end_bounds < bank_len {
                slc_to_search = bank_uint_slc[bank_max_n_idx:slc_end_bounds]
            } else {
                slc_to_search = bank_uint_slc[bank_max_n_idx:]
            }
            bank_slc_max_n, bank_slc_max_n_idx := get_slice_max(slc_to_search)
            // fmt.Println(fmt.Sprintf("bank_uint_slc[%d:%d - %d + 1]", bank_max_n_idx, bank_len, batteries_needed), bank_uint_slc[bank_max_n_idx:bank_len - batteries_needed + 1])
            // fmt.Println("bank_slc_max_n:", bank_slc_max_n, ", bank_max_n_idx:", bank_max_n_idx)
            batteries_chosen_joltage[battery_idx] = bank_slc_max_n
            bank_max_n_idx += bank_slc_max_n_idx + 1
            batteries_needed -= 1
        }
        // fmt.Println("batteries_chosen_joltage:", batteries_chosen_joltage)

        bank_joltage := uint(0)
        for i, joltage := range batteries_chosen_joltage {
            bank_joltage += joltage * UIntPow(10, (PUZZLE2_BATTERIES_NEEDED - uint(i+1)))
        }
        fmt.Println("Bank", bank_idx, "joltage:", bank_joltage)
        // fmt.Println()
        largest_joltage_sum += bank_joltage
    }

    fmt.Println("Sum of largest joltages:", largest_joltage_sum)
}

