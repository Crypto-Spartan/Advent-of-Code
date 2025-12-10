package day3

import (
    "fmt"
    //"strconv"
    "main/utils"
    //"slices"
    "strconv"
    "strings"
)

const DAY uint = 3
const digits_str = "0123456789"


func Puzzle1() {
    fmt.Println("Day 3, Puzzle 1!")

    nums_to_mult := get_multiply_numbers_puz1()
    //fmt.Println("Numbers to multiply:", nums_to_mult)

    var sum_nums_mult uint64 = 0
    for _, v := range nums_to_mult {
        sum_nums_mult += v[0] * v[1]
    }

    fmt.Println("Sum of multiplication instructions:", sum_nums_mult)
}


func get_multiply_numbers_puz1() [][2]uint64 {
    input_str := utils.ReadInputToString(utils.Get_Input_Filepath(DAY))
    //bar := progressbar.Default(int64(len(input_str)))

    var nums_to_mult [][2]uint64
    for {
        // beginning of instruction
        start := strings.Index(input_str, "mul(")
        if start == -1 {
            break
        }
        input_str = input_str[start+4:]
        //bar.Add(start)

        // first num
        valid_count := check_digits(input_str)
        if valid_count == 0 {
            //fmt.Println("flag#1")
            continue
        }
        //fmt.Println("flag#2")
        first_num, err := strconv.ParseUint(input_str[:valid_count], 10, 0)
        utils.Check(err)
        input_str = input_str[valid_count:]
        //bar.Add(valid_count)

        // comma
        if input_str[0] != ',' {
            //fmt.Println("flag#3")
            continue
        }
        //fmt.Println("flag#4")
        input_str = input_str[1:]
        //bar.Add(1)

        // second num
        valid_count = check_digits(input_str)
        if valid_count == 0 {
            //fmt.Println("flag#5")
            continue
        }
        //fmt.Println("flag#6")
        second_num, err := strconv.ParseUint(input_str[:valid_count], 10, 0)
        utils.Check(err)
        input_str = input_str[valid_count:]
        //bar.Add(valid_count)

        // closing paren
        if input_str[0] != ')' {
            //fmt.Println("flag#7")
            continue
        }
        //fmt.Println("flag#8")
        input_str = input_str[1:]
        //bar.Add(1)

        nums_to_mult = append(nums_to_mult, [2]uint64{first_num, second_num})
        //fmt.Println("appended")
    }

    return nums_to_mult
}


func Puzzle2() {
    fmt.Println("Day 3, Puzzle 2!")

    nums_to_mult := get_multiply_numbers_puz2()
    //fmt.Println("Numbers to multiply:", nums_to_mult)

    var sum_nums_mult uint64 = 0
    for _, v := range nums_to_mult {
        sum_nums_mult += v[0] * v[1]
    }

    fmt.Println("Sum of enabled multiplication instructions:", sum_nums_mult)
}


func get_multiply_numbers_puz2() [][2]uint64 {
    input_str := utils.ReadInputToString(utils.Get_Input_Filepath(DAY))

    var nums_to_mult [][2]uint64
    enabled := true
    for {
        toggle_start := strings.Index(input_str, "do")
        mul_start := strings.Index(input_str, "mul(")
        if mul_start == -1 || (toggle_start == -1 && !enabled) {
            break
        }

        if mul_start > toggle_start && toggle_start >= 0 {
            // validate toggle instruction
            if input_str[toggle_start:toggle_start+4] == "do()" {
                enabled = true
                input_str = input_str[toggle_start+4:]
            } else if input_str[toggle_start:toggle_start+7] == "don't()" {
                enabled = false
                input_str = input_str[toggle_start+7:]
            }
            continue
        }

        input_str = input_str[mul_start+4:]

        if !enabled {
            continue
        }

        // first num
        valid_count := check_digits(input_str)
        if valid_count == 0 {
            continue
        }
        first_num, err := strconv.ParseUint(input_str[:valid_count], 10, 0)
        utils.Check(err)
        input_str = input_str[valid_count:]

        // comma
        if input_str[0] != ',' {
            continue
        }
        input_str = input_str[1:]

        // second num
        valid_count = check_digits(input_str)
        if valid_count == 0 {
            continue
        }
        second_num, err := strconv.ParseUint(input_str[:valid_count], 10, 0)
        utils.Check(err)
        input_str = input_str[valid_count:]

        // closing paren
        if input_str[0] != ')' {
            continue
        }
        input_str = input_str[1:]

        nums_to_mult = append(nums_to_mult, [2]uint64{first_num, second_num})
    }

    return nums_to_mult
}


func check_digits(input_str string) int {
    valid_count := 0
    for i := 0; i < 3; i++ {
        if strings.IndexByte(digits_str, input_str[i]) == -1 {
            break
        } else {
            valid_count++
        }
    }
    return valid_count
}
