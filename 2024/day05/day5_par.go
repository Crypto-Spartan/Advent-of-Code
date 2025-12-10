package day5

import (
    "fmt"
    // "strings"
    //"strconv"
    // "main/utils"
    // "slices"
    // "strconv"
    // "strings"
)


func Puzzle1_par() {
    fmt.Println("Day 5, Puzzle 1! (parallel)")

    all_rules, all_updates := get_rules_and_updates()
    // fmt.Println("all_rules:", all_rules)
    // fmt.Println("all_updates:", all_updates)
    
    results := make(chan uint, len(all_updates))
    for _, pages := range all_updates {

        go func(p []uint){
            valid := check_update(p, all_rules)
            if valid {
                mid_idx := (len(p) - 1) / 2
                results <- p[mid_idx]
            } else {
                results <- 0
            }
        }(pages)
    }

    var sum_mid_pages uint = 0
    for i := 0; i < len(all_updates); i++ {
        sum_mid_pages += <-results
    } 

    fmt.Println("Sum of middle pages from valid updates:", sum_mid_pages)
}


func Puzzle2_par() {
    fmt.Println("Day 5, Puzzle 2! (parallel)")

    all_rules, all_updates := get_rules_and_updates()
    // fmt.Println("all_rules:", all_rules)
    // fmt.Println("all_updates:", all_updates)

    results := make(chan uint, len(all_updates))
    for _, pages := range all_updates {

        go func(p []uint){
            valid := check_update(p, all_rules)
            if !valid {
                p_fixed := fix_update(p, all_rules)
                if len(p) != len(p_fixed) {
                    panic("pages_fixed has incorrect length")
                }
                mid_idx := (len(p_fixed) - 1) / 2
                results <- p_fixed[mid_idx]
            } else {
                results <- 0
            }
        }(pages)
    }

    var sum_mid_pages uint = 0
    for i := 0; i < len(all_updates); i++ {
        sum_mid_pages += <-results
    } 

    fmt.Println("Sum of middle pages from fixed invalid updates:", sum_mid_pages)
}


