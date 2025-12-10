package day10

import (
	"fmt"
	"main/utils"
	// "slices"
	// "strings"
)

const DAY uint = 10

var NUM_ROWS uint
var NUM_COLS uint
var MAX_ROWS_IDX uint
var MAX_COLS_IDX uint

type coordinate struct {
	x uint
	y uint
}
type search_node struct {
	coord coordinate
	gridmap_idx uint
}

// array[width * row + col]


func Puzzle1() {
	fmt.Println("Day 10, Puzzle 1!")

	gridmap := get_gridmap()
	trailheads := get_trailheads(gridmap)
	// fmt.Println("trailheads:", trailheads)
	score := uint(0)
	for _, th := range trailheads {
		score += uint(get_trailhead_score_pt1(th, gridmap))
	}

	fmt.Println("Sum of trailhead scores:", score)
}


// BFS (G, s)                   //Where G is the graph and s is the source node
//       let Q be queue.
//       Q.enqueue( s ) //Inserting s in queue until all its neighbour vertices are marked.
//
//       mark s as visited.
//       while ( Q is not empty)
//            //Removing that vertex from queue,whose neighbour will be visited now
//            v  =  Q.dequeue( )
//
//           //processing all the neighbours of v  
//           for all neighbours w of v in Graph G
//                if w is not visited 
//                         Q.enqueue( w )             //Stores w in Q to further visit its neighbour
//                         mark w as visited.

func get_trailhead_score_pt1(trailhead coordinate, gridmap []byte) byte {
	curr_map_idx :=  NUM_COLS * trailhead.y + trailhead.x
	init_node := search_node{coord: trailhead, gridmap_idx: curr_map_idx}
	queue := []search_node{init_node}
	unique_trail_ends := map[coordinate]struct{}{}

	for len(queue) > 0 {
		curr_node := queue[0]
		// fmt.Println("curr_node:", curr_node)
		// fmt.Println("curr_val:", gridmap[curr_node.gridmap_idx])
		queue = queue[1:]

		if gridmap[curr_node.gridmap_idx] == 9 {
			unique_trail_ends[curr_node.coord] = struct{}{}
			continue
		}

		next_nodes := get_next_nodes(curr_node)
		// fmt.Println("next_nodes:", next_nodes)
		for _, next_node := range next_nodes {
			if gridmap[curr_node.gridmap_idx] + 1 == gridmap[next_node.gridmap_idx] {
				queue = append(queue, next_node)
			}
		}
	}

	return byte(len(unique_trail_ends))
}


func Puzzle2() {
	fmt.Println("Day 10, Puzzle 2!")

	gridmap := get_gridmap()
	trailheads := get_trailheads(gridmap)
	// fmt.Println("trailheads:", trailheads)
	score := uint(0)
	for _, th := range trailheads {
		score += get_trailhead_score_pt2(th, gridmap)
	}

	fmt.Println("Sum of trailhead scores:", score)
}


func get_trailhead_score_pt2(trailhead coordinate, gridmap []byte) uint {
	// curr_map_idx :=  NUM_COLS * trailhead.y + trailhead.x
	// init_node := search_node{coord: trailhead, gridmap_idx: curr_map_idx}
	queue := [][]coordinate{{trailhead}}
	// unique_trails := [][]coordinate{}
	unique_trails := make(map[[10]coordinate]struct{})

	for len(queue) > 0 {
		curr_path := queue[0]
		// fmt.Println("curr_node:", curr_node)
		// fmt.Println("curr_val:", gridmap[curr_node.gridmap_idx])
		queue = queue[1:]

		curr_coord := curr_path[len(curr_path) - 1]
		curr_map_idx :=  NUM_COLS * curr_coord.y + curr_coord.x

		if gridmap[curr_map_idx] == 9 {
			if len(curr_path) != 10 {
				panic("invalid path length")
			}
			path_array := [10]coordinate{}
			copy(path_array[:], curr_path)
			unique_trails[path_array] = struct{}{}
			continue
		}

		next_nodes := get_next_nodes(search_node{coord: curr_coord, gridmap_idx: curr_map_idx})
		// fmt.Println("next_nodes:", next_nodes)
		for _, next_node := range next_nodes {
			if gridmap[curr_map_idx] + 1 == gridmap[next_node.gridmap_idx] {
				new_path := make([]coordinate, len(curr_path) + 1)
				copy(new_path, curr_path)
				new_path[len(curr_path)] = next_node.coord
				queue = append(queue, new_path)
			}
		}
	}

	return uint(len(unique_trails))
}


func get_next_nodes(curr_node search_node) []search_node {
	next_nodes := [4]search_node{}
	idx := 0
	curr_x := curr_node.coord.x
	curr_y := curr_node.coord.y
	curr_map_idx := curr_node.gridmap_idx

	// check left
	if curr_x > 0 {
		left_coord := coordinate{x: curr_x - 1, y: curr_y}
		left_node := search_node{coord: left_coord, gridmap_idx: curr_map_idx - 1} 
		next_nodes[idx] = left_node
		idx++
	}

	// check right
	if curr_x < MAX_COLS_IDX {
		right_coord := coordinate{x: curr_x + 1, y: curr_y}
		right_node := search_node{coord: right_coord, gridmap_idx: curr_map_idx + 1} 
		next_nodes[idx] = right_node
		idx++
	}

	// check up
	if curr_y > 0 {
		up_coord := coordinate{x: curr_x, y: curr_y - 1}
		up_node := search_node{coord: up_coord, gridmap_idx: curr_map_idx - NUM_COLS} 
		next_nodes[idx] = up_node
		idx++
	}

	// check down
	if curr_y < MAX_ROWS_IDX {
		down_coord := coordinate{x: curr_x, y: curr_y + 1}
		down_node := search_node{coord: down_coord, gridmap_idx: curr_map_idx + NUM_COLS} 
		next_nodes[idx] = down_node
		idx++
	}

	return next_nodes[:idx]
}


func get_trailheads(gridmap []byte) []coordinate {
	// size := NUM_COLS
	// gridmap_len := uint(len(gridmap))
	trailheads := []coordinate{}

	for row := uint(0); row < NUM_ROWS; row++ {
		i := row * NUM_COLS
		j := i + NUM_COLS
		chunk := gridmap[i:j]

		for col, n := range chunk {
			if n == 0 {
				c := coordinate{x: uint(col), y: row}
				trailheads = append(trailheads, c)
			}
		}
	}

	return trailheads
}


func get_gridmap() []byte {

	input_lines := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
	// for _, s := range input_lines {
	// 	fmt.Println(s)
	// }

	NUM_ROWS = uint(len(input_lines))
	NUM_COLS = uint(len(input_lines[0]))
	MAX_ROWS_IDX = NUM_ROWS - 1
	MAX_COLS_IDX = NUM_COLS - 1

	gridmap := make([]byte, NUM_COLS*NUM_ROWS)
	for i, r := range input_lines {
		// fmt.Println(fmt.Sprintf("%#v", r))
		for j, c := range r {
			gridmap[(NUM_COLS*uint(i))+uint(j)] = byte(c - '0')
		}
	}

	return gridmap
}
