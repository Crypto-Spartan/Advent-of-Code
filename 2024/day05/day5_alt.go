package day5

import (
    "fmt"
    "main/utils"
    "os"
    "slices"
    "strconv"
    "strings"
)

func AltSolution() {
    input, _ := os.ReadFile(utils.Get_Input_Filepath(DAY))
    split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

    cmp := func(a, b string) int {
        for _, s := range strings.Split(split[0], "\n") {
            if s := strings.Split(s, "|"); s[0] == a && s[1] == b {
                return -1
            }
        }
        return 0
    }

    run := func(sorted bool) (r int) {
        for _, s := range strings.Split(split[1], "\n") {
            if s := strings.Split(s, ","); slices.IsSortedFunc(s, cmp) == sorted {
                slices.SortFunc(s, cmp)
                n, _ := strconv.Atoi(s[len(s)/2])
                r += n
            }
        }
        return r
    }

    fmt.Println(run(true))
    fmt.Println(run(false))
}