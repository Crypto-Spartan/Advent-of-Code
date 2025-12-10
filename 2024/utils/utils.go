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
    return filepath.Join("/workspaces/Advent-of-Code/2024/", fmt.Sprintf("day%d/input.txt", day_num))
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
    return strings.Split(data_str, "\n")
}
