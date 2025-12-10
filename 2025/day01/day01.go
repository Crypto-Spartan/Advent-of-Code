package day01

import (
    "fmt"
    "main/utils"
    "strconv"
)

const DAY uint = 1

func Puzzle1() {
    fmt.Println("Day 1, Puzzle 1!")

    rotations := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    // fmt.Println("rotations:", rotations)
    // fmt.Println("len(rotations):", len(rotations))

    var dial int = 50
    // fmt.Println("dial:", dial)
    // fmt.Println()
    var dial_zero_count uint = 0
    for _, r := range rotations {
        num_turns, err := strconv.Atoi(r[1:])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }
        // if num_turns > 100 {
        //     fmt.Println("r:", r)
        //     fmt.Println("num_turns:", num_turns)
        //     fmt.Println("num_turns//100:", num_turns/100)
        //     fmt.Println()
        // }

        if r[0] == 'L' {
            // dial = turn_left(dial, num_turns)
            dial = (((dial - num_turns) % 100) + 100) % 100
        } else if r[0] == 'R' {
            // dial = turn_right(dial, num_turns)
            dial = (dial + num_turns) % 100
        } else {
            panic(fmt.Sprintf("Rotation not L or R, got: %s", r[0]))
        }

        if dial == 0 {
            dial_zero_count += 1
        } else if dial < 0 || dial > 100 {
            panic(fmt.Sprintf("dial out of bounds - dial: %d", dial))
        }

        // fmt.Println("dial:", dial)
        // fmt.Println()
    }

    fmt.Println("dial_zero_count:", dial_zero_count)
}

func turn_left(dial int, num_turns int) (int) {
    for i := 0; i < num_turns; i++ {
        if dial == 0 {
            dial = 99
        } else {
            dial -= 1
        }
    }
    return dial
}

func turn_right(dial int, num_turns int) (int) {
    for i := 0; i < num_turns; i++ {
        if dial == 99 {
            dial = 0
        } else {
            dial += 1
        }
    }
    return dial
}

func Puzzle2() {
    fmt.Println("Day 1, Puzzle 2!")

    rotations := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))

    var dial int = 50
    var dial_zero_count int = 0
    for _, r := range rotations {
        num_turns, err := strconv.Atoi(r[1:])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }
        // if num_turns > 100 {
        // if true {
            // fmt.Println("rotation:", r)
            // fmt.Println("num_turns:", num_turns)
            // fmt.Println("num_turns//100:", num_turns/100)
            // fmt.Println()
        // }

        var hypothetical_dial int
        if r[0] == 'L' {
            // dial = turn_left(dial, num_turns)
            if dial == 0 {
                dial = 99
                num_turns -= 1
            }
            hypothetical_dial = dial - num_turns
            dial = ((hypothetical_dial % 100) + 100) % 100
            // fmt.Println("((hypothetical_dial - 100) / 100) * -1:", ((hypothetical_dial - 100) / 100) * -1)
            dial_zero_count += ((hypothetical_dial - 100) / 100) * -1
        } else if r[0] == 'R' {
            // dial = turn_right(dial, num_turns)
            hypothetical_dial = dial + num_turns
            dial = hypothetical_dial % 100
            // fmt.Println("hypothetical_dial / 100:", hypothetical_dial / 100)
            dial_zero_count += hypothetical_dial / 100
        } else {
            panic(fmt.Sprintf("Rotation not L or R, got: %s", r[0]))
        }

        if dial < 0 || dial > 100 {
            panic(fmt.Sprintf("dial out of bounds - dial: %d", dial))
        }

        // fmt.Println("dial:", dial)
        // fmt.Println("dial_zero_count:", dial_zero_count)
        // fmt.Println()
    }

    fmt.Println("final dial_zero_count:", dial_zero_count)
}

