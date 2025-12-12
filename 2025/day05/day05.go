package day05

import (
    "fmt"
    "main/utils"
    "strconv"
    "strings"
    "sort"
)

const DAY uint = 5

type FreshIngredientsRange struct {
    Start int
    End int
}

func merge_ranges(ranges []FreshIngredientsRange) []FreshIngredientsRange {
    n := len(ranges)
    sort.Slice(ranges, func(i, j int) bool { return ranges[i].Start < ranges[j].Start })

    var ranges_merged []FreshIngredientsRange
    for i, r1 := range ranges {
        start := r1.Start
        end := r1.End

        ranges_merged_len := len(ranges_merged)
        if ranges_merged_len > 0 && ranges_merged[ranges_merged_len - 1].End >= end {
            continue
        }

        for j := i + 1; j < n; j++ {
            if ranges[j].Start <= end {
                end = max(end, ranges[j].End)
            }
        }
        ranges_merged = append(ranges_merged, FreshIngredientsRange{Start: start, End: end})
    }

    return ranges_merged
}

func Puzzle1() {
    fmt.Printf("Day %d, Puzzle 1!\n", DAY)

    strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    // fmt.Println("strs:", strs)
    var empty_line_idx int
    for idx, s := range strs {
        if len(s) == 0 {
            empty_line_idx = idx
        }
    }
    // fmt.Println("empty_line_idx:", empty_line_idx)

    fresh_ingredients_ranges_strs := strs[:empty_line_idx]
    ingredients_to_check_strs := strs[empty_line_idx+1:]

    fresh_ingredients_ranges := make([]FreshIngredientsRange, len(fresh_ingredients_ranges_strs))
    for i, range_str := range fresh_ingredients_ranges_strs {
        range_str_split := strings.Split(range_str, "-")
        if len(range_str_split) != 2 {
            panic("len(range_str_split) != 2")
        }

        range_start, err := strconv.Atoi(range_str_split[0])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        range_end, err := strconv.Atoi(range_str_split[1])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        fresh_ingredients_ranges[i] = FreshIngredientsRange{Start: range_start, End: range_end}
    }
    // fmt.Println("fresh_ingredients_ranges:", fresh_ingredients_ranges)

    fresh_ingredients_ranges_merged := merge_ranges(fresh_ingredients_ranges)
    // fmt.Println("fresh_ingredients_ranges_merged:", fresh_ingredients_ranges_merged)

    fresh_ingredients_count := uint(0)
    for _, ingredient_str := range ingredients_to_check_strs {
        ingredient_int, err := strconv.Atoi(ingredient_str)
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        for _, ingredient_range := range fresh_ingredients_ranges_merged {
            if ingredient_int >= ingredient_range.Start && ingredient_int <= ingredient_range.End {
                fresh_ingredients_count++
                break
            }
        }
    }

    fmt.Println("# of fresh ingredients (inventory):", fresh_ingredients_count)
}

func Puzzle2() {
    fmt.Println(fmt.Sprintf("Day %d, Puzzle 2!", DAY))

    strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    // fmt.Println("strs:", strs)
    var empty_line_idx int
    for idx, s := range strs {
        if len(s) == 0 {
            empty_line_idx = idx
        }
    }
    // fmt.Println("empty_line_idx:", empty_line_idx)

    fresh_ingredients_ranges_strs := strs[:empty_line_idx]
    fresh_ingredients_ranges := make([]FreshIngredientsRange, len(fresh_ingredients_ranges_strs))
    for i, range_str := range fresh_ingredients_ranges_strs {
        range_str_split := strings.Split(range_str, "-")
        if len(range_str_split) != 2 {
            panic("len(range_str_split) != 2")
        }

        range_start, err := strconv.Atoi(range_str_split[0])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        range_end, err := strconv.Atoi(range_str_split[1])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        fresh_ingredients_ranges[i] = FreshIngredientsRange{Start: range_start, End: range_end}
    }
    // fmt.Println("fresh_ingredients_ranges:", fresh_ingredients_ranges)

    fresh_ingredients_ranges_merged := merge_ranges(fresh_ingredients_ranges)
    // fmt.Println("fresh_ingredients_ranges_merged:", fresh_ingredients_ranges_merged)

    fresh_ingredients_count := uint(0)
    fmt.Println("Starting loop")
    for _, ingredient_range := range fresh_ingredients_ranges_merged {
        if ingredient_range.Start == ingredient_range.End {
            fresh_ingredients_count++
        } else if ingredient_range.Start > ingredient_range.End {
            fmt.Println("ingredient_range:", ingredient_range)
            panic("range start larger than range end")
        } else {
            fresh_ingredients_count += uint(ingredient_range.End + 1 - ingredient_range.Start)
            // for i := ingredient_range.Start; i <= ingredient_range.End; i++ {
            //     fresh_ingredients_count++
            // }
        }
    }

    fmt.Println("total # of fresh ingredients:", fresh_ingredients_count)
}

