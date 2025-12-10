package day1

import (
    "fmt"
    "main/utils"
    "slices"
    "strconv"
    "strings"
)

const DAY uint = 1

func Puzzle1() {
    fmt.Println("Day 1, Puzzle 1!")

    list1, list2 := get_lists()
    slices.Sort(list1)
    slices.Sort(list2)

    var sum_differences uint64 = 0
    for i := range list1 {
        if list1[i] > list2[i] {
            sum_differences += list1[i] - list2[i]
        } else {
            sum_differences += list2[i] - list1[i]
        }
    }

    fmt.Println("Sum of differences:", sum_differences)
}


func Puzzle2() {
    fmt.Println("Day 1, Puzzle 2!")

    list1, list2 := get_lists()

    list2_counter := make(map[uint64]uint64)
    for _, x := range list2 {
        list2_counter[x]++
    }
    // fmt.Println(list2_counter)

    var similarity_score uint64 = 0
    for _, x := range list1 {
        if c := list2_counter[x]; c > 0 {
            similarity_score += x * c
        }
    }

    fmt.Println("Similarity score:", similarity_score)
}


func get_lists() ([]uint64, []uint64) {
    input_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))

    list1, list2 := parse_input(input_lines)
    if len(list1) != len(list2) {
        panic("Lists are not the same length")
    }

    return list1, list2
}


func parse_input(input_lines []string) ([]uint64, []uint64) {
    var list1, list2 []uint64

    for _, v := range input_lines {
        //fmt.Println(fmt.Sprintf("%#v", line))
        line_split := strings.SplitN(v, " ", 2)
        //fmt.Println(fmt.Sprintf("%#v", line_split))

        str1, str2 := line_split[0], strings.TrimSpace(line_split[1])
        int1, err := strconv.ParseUint(str1, 10, 0)
        utils.Check(err)

        int2, err := strconv.ParseUint(str2, 10, 0)
        utils.Check(err)
        //fmt.Println(int1, int2)

        list1 = append(list1, int1)
        list2 = append(list2, int2)
    }

    //fmt.Printf("len=%d cap=%d %v\n", len(list1), cap(list1), list1)
    return list1, list2
}
