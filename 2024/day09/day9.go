package day9

import (
	"fmt"
	"main/utils"
	// "os"

	// "strconv"

	// "slices"
	"strings"
)

const DAY uint = 9

// array[width * row + col]

func Puzzle1() {
	fmt.Println("Day 9, Puzzle 1!")

	diskmap := get_diskmap()
	// fmt.Println(diskmap)
	expanded_blocks := get_expanded_blocks(diskmap)
	// fmt.Println("expanded_blocks:", expanded_blocks)
	diskmap_contiguous := make_contiguous(expanded_blocks)
	// fmt.Println("diskmap_contiguous:", diskmap_contiguous)
	checksum := calculate_checksum(diskmap_contiguous)
	fmt.Println("checksum:", checksum)

	// fmt.Println("Number of unique antinodes:", len(antinodes))
}


func calculate_checksum(diskmap []int) int {
	checksum := 0
	for i, v := range diskmap {
		if v == -1 {
			continue
		}
		checksum += i * v
	}
	return checksum
}


func make_contiguous(diskmap []int) []int {
	
	dot_idx := 0
	for i := len(diskmap) - 1; i >= 0; i-- {
		// fmt.Println("diskmap:", string(diskmap))

		n := diskmap[i]
		if n == -1 {
			continue
		}

		dot_idx += find_first_dot(diskmap[dot_idx:])
		if dot_idx == -2 {
			panic("no dots found")
		}

		contiguous := true
		for _, x := range diskmap[dot_idx+1:] {
			if x != -1 {
				contiguous = false
				break
			}
		}

		if contiguous {
			return diskmap
		}

		diskmap[dot_idx] = n
		diskmap[i] = -1
	}

	panic("unable to make diskmap contiguous")
}


func find_first_dot(diskmap []int) int {
	dot_idx := -2
	for i, v := range diskmap {
		if v == -1 {
			dot_idx = i
			break
		}
	}
	return dot_idx
}


func Puzzle2() {
	fmt.Println("Day 9, Puzzle 2!")

	diskmap := get_diskmap()
	// fmt.Println(diskmap)
	expanded_blocks := get_expanded_blocks(diskmap)
	// fmt.Println("expanded_blocks:", expanded_blocks)
	diskmap_compact := make_compact(expanded_blocks)
	// fmt.Println("diskmap_compact:", diskmap_compact)
	checksum := calculate_checksum(diskmap_compact)
	fmt.Println("checksum:", checksum)
}


func make_compact(diskmap []int) []int {
	
	
	// dot_idx := 0
	i := len(diskmap) - 1
	for i > 0 {
		// fmt.Println("diskmap:", diskmap)

		n := diskmap[i]
		// fmt.Println("fileid:", n)
		if n == -1 {
			i--
			continue
		} else if n == 0 {
			break
		}

		filesize := 1
		file_idx := -1
		for j := i-1; j >= 0; j-- {
			if diskmap[j] != n {
				file_idx = j + 1
				break
			}
			filesize++
		}
		// fmt.Println("filesize:", filesize)

		if file_idx == -1 {
			panic("file_idx not set")
		}

		spacesize := 0
		space_idx := -1
		for j := 0; j < file_idx; j++ {
			if diskmap[j] == -1 {
				spacesize++
				if space_idx == -1 {
					space_idx = j
				}
			} else {
				spacesize = 0
				space_idx = -1
				continue
			}

			if spacesize == filesize {
				for k := space_idx; k < space_idx + spacesize; k++ {
					diskmap[k] = n
				}

				// file_idx := i - spacesize
				for k := file_idx; k < file_idx + spacesize; k++ {
					diskmap[k] = -1
				}
				break
			}
		}


		i = file_idx - 1
	}

	return diskmap
}


func get_expanded_blocks(diskmap []byte) []int {
	expanded := []int{}
	is_file := true
	var i int
	for _, s := range diskmap {
		if is_file {
			// fmt.Println("s:", string(s))
			// fmt.Println("int(s-'0'):", int(s-'0'))
			for j := 0; j < int(s - '0'); j++ {
				expanded = append(expanded, i)
			}
			// expanded += strings.Repeat(strconv.Itoa(i), int(s - '0'))
			is_file = false
			i++
		} else {
			for j := 0; j < int(s - '0'); j++ {
				expanded = append(expanded, -1)
			}
			// expanded += strings.Repeat(".", int(s - '0'))
			is_file = true
		}
	}
	return expanded
}


func get_diskmap() []byte {
	input_str := utils.ReadInputToString(utils.Get_Input_Filepath(DAY))
	return []byte(strings.TrimSpace(input_str))
}
