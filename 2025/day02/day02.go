package day02

import (
    "fmt"
    "main/utils"
    "strconv"
    "strings"
    // "slices"
)

const DAY uint = 2

func Puzzle1() {
    fmt.Println("Day 2, Puzzle 1!")

    ranges_str := strings.TrimSpace(utils.ReadInputToString(utils.Get_Input_Filepath(DAY)))
    // fmt.Println("ranges_str:", ranges_str)
    ranges_slc := strings.Split(ranges_str, ",")
    // fmt.Println("ranges_slc:", ranges_slc)

    var invalid_ids_sum uint = 0
    for _, r := range ranges_slc {
        r_split := strings.Split(r, "-")
        // fmt.Println("r_split:", r_split)

        first_id, err := strconv.Atoi(r_split[0])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }
        last_id, err := strconv.Atoi(r_split[1])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        for id := first_id; id <= last_id; id++ {
            id_str := strconv.Itoa(id)
            id_str_len := len(id_str)
            if id_str_len % 2 == 0 {
                id_str_len_half := id_str_len / 2
                if id_str[:id_str_len_half] == id_str[id_str_len_half:] {
                    // fmt.Println("id has repeated number:", id_str)
                    invalid_ids_sum += uint(id)
                }
            }
        }
    }

    fmt.Println("Sum of invalid ids:", invalid_ids_sum)
}


func Puzzle2() {
    fmt.Println("Day 2, Puzzle 2!")

    ranges_str := strings.TrimSpace(utils.ReadInputToString(utils.Get_Input_Filepath(DAY)))
    // fmt.Println("ranges_str:", ranges_str)
    ranges_slc := strings.Split(ranges_str, ",")
    // fmt.Println("ranges_slc:", ranges_slc)

    var invalid_ids_sum uint = 0
    for _, r := range ranges_slc {
        r_split := strings.Split(r, "-")
        // fmt.Println("r_split:", r_split)

        first_id, err := strconv.Atoi(r_split[0])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }
        last_id, err := strconv.Atoi(r_split[1])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        for id := first_id; id <= last_id; id++ {
            id_str := strconv.Itoa(id)
            id_str_len := uint(len(id_str))
            id_str_len_half := id_str_len / 2

            for chunk_size := uint(1); chunk_size <= id_str_len_half; chunk_size++ {
                id_str_chunks := chunkString(id_str, chunk_size)
                chunk_1 := id_str_chunks[0]
                chunk_eq := true

                for _, chunk_n := range id_str_chunks[1:] {
                    if chunk_1 != chunk_n {
                        chunk_eq = false
                        break
                    }
                }

                if !chunk_eq {
                    continue
                } else {
                    // fmt.Println("id has repeated number:", id_str)
                    invalid_ids_sum += uint(id)
                    break
                }
            }
        }
    }

    fmt.Println("Sum of invalid ids:", invalid_ids_sum)
}

func chunkString(s string, chunkSize uint) []string {
    var chunks []string
    if chunkSize <= uint(0) {
        return chunks // Return empty slice for invalid chunk size
    }

    for i := uint(0); i < uint(len(s)); i += chunkSize {
        end := i + chunkSize
        if end > uint(len(s)) {
            end = uint(len(s))
        }
        chunks = append(chunks, s[i:end])
    }
    return chunks
}

