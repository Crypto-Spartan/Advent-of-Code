package day5

import (
	"fmt"
	"strings"
	//"strconv"
	"main/utils"
	"slices"
	"strconv"
	// "strings"
)

const DAY uint = 5
type rule_obj struct {
    before uint
    after uint
}


func Puzzle1() {
    fmt.Println("Day 5, Puzzle 1!")

    all_rules, all_updates := get_rules_and_updates()
    // fmt.Println("all_rules:", all_rules)
    // fmt.Println("all_updates:", all_updates)

    var sum_mid_pages uint = 0
    for _, pages := range all_updates {

        valid := check_update(pages, all_rules)
        if valid {
            mid_idx := (len(pages) - 1) / 2
            sum_mid_pages += pages[mid_idx]
        }
    }

    fmt.Println("Sum of middle pages from valid updates:", sum_mid_pages)
}


func Puzzle2() {
    fmt.Println("Day 5, Puzzle 2!")

    all_rules, all_updates := get_rules_and_updates()
    // fmt.Println("all_rules:", all_rules)
    // fmt.Println("all_updates:", all_updates)

    var sum_mid_pages uint = 0
    for _, pages := range all_updates {

        valid := check_update(pages, all_rules)
        if !valid {
            pages_fixed := fix_update(pages, all_rules)
            if len(pages) != len(pages_fixed) {
                panic("pages_fixed has incorrect length")
            }
            mid_idx := (len(pages_fixed) - 1) / 2
            sum_mid_pages += pages_fixed[mid_idx]
        }
    }

    fmt.Println("Sum of middle pages from fixed invalid updates:", sum_mid_pages)
}


func check_update(pages []uint, all_rules []rule_obj) bool {

    valid := true
    for _, r := range all_rules {
        before_i := slices.Index(pages, r.before)
        if before_i == -1 {
            continue
        }
        after_i := slices.Index(pages, r.after)
        if after_i == -1 {
            continue
        }

        if before_i > after_i {
            valid = false
            break
        }
    }

    return valid
}


func fix_update(pages []uint, all_rules []rule_obj) []uint {

    valid := false
    for valid == false {
        for _, r := range all_rules {
            before_i := slices.Index(pages, r.before)
            if before_i == -1 {
                continue
            }
            after_i := slices.Index(pages, r.after)
            if after_i == -1 {
                continue
            }

            if before_i > after_i {
                pages[before_i], pages[after_i] = pages[after_i], pages[before_i]
                break
            }
        }

        valid = check_update(pages, all_rules)
    }

    return pages
}


func get_rules_and_updates() ([]rule_obj, [][]uint) {
    input_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))

    var rules_lines []string
    var page_updates_lines []string
    for i, r := range input_lines {
        if strings.TrimSpace(r) == "" {
            rules_lines = input_lines[:i]
            page_updates_lines = input_lines[i+1:]
            break
        }
    }
    if len(rules_lines) == 0 || len(page_updates_lines) == 0 {
        panic("unable to find rules or page lines")
    }

    rules := parse_rules(rules_lines)
    page_updates := parse_page_updates(page_updates_lines)

    return rules, page_updates
}


func parse_rules(rules_lines []string) []rule_obj {
    rules := make([]rule_obj, len(rules_lines))

    for i, r := range rules_lines {
        before_str, after_str, found := strings.Cut(r, "|")
        if !found {
            panic(fmt.Sprintf("error with string cut for rule `%s`", r))
        }

        before_u64, err := strconv.ParseUint(before_str, 10, 0)
        utils.Check(err)
        after_u64, err := strconv.ParseUint(after_str, 10, 0)
        utils.Check(err)
        

        rule := rule_obj{
            before: uint(before_u64),
            after: uint(after_u64),
        }

        rules[i] = rule
    }

    return rules
}


func parse_page_updates(page_updates_lines []string) [][]uint {
    page_updates := make([][]uint, len(page_updates_lines))

    for i, r := range page_updates_lines {
        split := strings.Split(r, ",")

        as_uints := make([]uint, len(split))
        for j, s := range split {
            n, err := strconv.ParseUint(s, 10, 0)
            utils.Check(err)

            as_uints[j] = uint(n)
        }

        page_updates[i] = as_uints
    }

    return page_updates
}

