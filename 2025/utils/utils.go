package utils

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var HOME string

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func Get_Input_Filepath(day_num uint) string {
    return filepath.Join("/workspaces/Advent-of-Code/2025/", fmt.Sprintf("day%02d/input.txt", day_num))
}

func ReadInputToBytes(input_filepath string) []byte {
    data_bytes, err := os.ReadFile(input_filepath)
    Check(err)
    return data_bytes
}

func ReadInputToString(input_filepath string) string {
    return string(ReadInputToBytes(input_filepath))
}

func ReadInputToLines(input_filepath string) []string {
    data_str := ReadInputToString(input_filepath)
    data_arr := strings.Split(data_str, "\n")
    data_arr_len := len(data_arr)
    last_idx := data_arr_len - 1
    if len(data_arr[last_idx]) == 0 {
        data_arr = data_arr[:last_idx]
    }
    return data_arr
}
