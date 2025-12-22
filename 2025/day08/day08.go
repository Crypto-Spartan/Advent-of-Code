package day08

import (
    "cmp"
    "fmt"
    "main/utils"
    "math"
    "slices"
    "strconv"
    "strings"

    "github.com/schollz/progressbar/v3"
)

const DAY uint = 8

type BoxCoord [3]int

func get_box_coords() []BoxCoord {
    coord_strs := utils.ReadInputToLines(utils.Get_Input_Filepath(DAY))
    // fmt.Println("coord_strs:", coord_strs)

    boxes := make([]BoxCoord, len(coord_strs))
    for box_idx, s := range coord_strs {
        coord_str_split := strings.Split(s, ",")
        if len(coord_str_split) != 3 {
            panic("len(coord_str_split) != 3")
        }

        var box_coord BoxCoord
        for coord_idx := 0; coord_idx < 3; coord_idx++ {
            i, err := strconv.Atoi(coord_str_split[coord_idx])
            if err != nil {
                panic(fmt.Sprintf("Error converting string to int: %s", err))
            }

            box_coord[coord_idx] = i
        }

        boxes[box_idx] = box_coord
    }

    return boxes
}

// generateCombinationsOf2 generates all pairs of BoxCoord.
func generateCombinationsOf2(boxes []BoxCoord) [][2]BoxCoord {
    var result [][2]BoxCoord
    n := len(boxes)
    // Generate combinations of 2
    for i := 0; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            result = append(result, [2]BoxCoord{boxes[i], boxes[j]})
        }
    }
    return result
}

func calc_distance(bc1, bc2 BoxCoord) uint {
    x := bc1[0] - bc2[0]
    y := bc1[1] - bc2[1]
    z := bc1[2] - bc2[2]

    x_sqd := uint(x * x)
    y_sqd := uint(y * y)
    z_sqd := uint(z * z)

    sum_sqrs := x_sqd + y_sqd + z_sqd
    // return math.Sqrt(float64(sum_sqrs))
    return sum_sqrs
}

func get_min_distance(boxes_combinations [][2]BoxCoord) ([2]BoxCoord, [][2]BoxCoord) {
    min_distance_val := uint(math.MaxUint)
    min_distance_idx := -1
    for i, combo := range boxes_combinations {
        d := calc_distance(combo[0], combo[1])

        if d < min_distance_val {
            min_distance_val = d
            min_distance_idx = i
        }
    }

    min_distance_coord := boxes_combinations[min_distance_idx]
    boxes_combinations = slices.Delete(boxes_combinations, min_distance_idx, min_distance_idx + 1)
    return min_distance_coord, boxes_combinations
}

// Adjacency list type to hold linked coordinates.
type Graph map[BoxCoord][]BoxCoord

func removeDuplicate[T comparable](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

// AddLink adds a link (edge) between two coordinates.
func (g Graph) AddLink(from, to BoxCoord) {
    to_in_from := false
    g_from := g[from]
    for _, c := range g_from {
        if to == c {
            to_in_from = true
            break
        }
    }
    if !to_in_from {
        g[from] = append(g_from, to)
    }

    // If undirected graph, also add the reverse link.
    g[to] = append(g[to], from)
    from_in_to := false
    g_to := g[to]
    for _, c := range g_to {
        if from == c {
            from_in_to = true
            break
        }
    }
    if !from_in_to {
        g[to] = append(g_to, from)
    }
}

// PrintGraph prints the graph (adjacency list).
func (g Graph) PrintGraph() {
    for coord, links := range g {
        fmt.Printf("Coordinate %v is linked to: ", coord)
        for _, link := range links {
            fmt.Printf("%v ", link)
        }
        fmt.Println()
    }
}

func MergeGraphs(graph1, graph2 Graph) Graph {
    // Create a new graph for merging
    mergedGraph := make(Graph)

    // Add all nodes and links from graph1 to the merged graph
    for node, links := range graph1 {
        mergedGraph[node] = links
    }

    // Add all nodes and links from graph2 to the merged graph
    for node, links := range graph2 {
        // If the node already exists in the merged graph, append the links
        if mgn, exists := mergedGraph[node]; exists {
            for _, c1 := range links {
                already_in_slice := false
                for _, c2 := range mgn {
                    if c1 == c2 {
                        already_in_slice = true
                        break
                    }
                }
                if !already_in_slice {
                    mergedGraph[node] = append(mergedGraph[node], c1)
                }
            }
            // mergedGraph[node] = append(mergedGraph[node], links...)
        } else {
            // If the node doesn't exist, add it with its links
            mergedGraph[node] = links
        }
    }

    return mergedGraph
}

func CanMerge(graph1, graph2 Graph) bool {
    // Loop through all nodes in graph1
    for node1 := range graph1 {
        // Check if any node in graph1 is connected to a node in graph2
        for node2 := range graph2 {
            // Check if there's a link from node1 in graph1 to node2 in graph2
            for _, link := range graph1[node1] {
                if link == node2 {
                    return true
                }
            }
        }
    }
    return false
}

func MergeGraphsIfNecessary(graphs []Graph) []Graph {
    // Iterate over the slice of graphs and attempt to merge them if necessary
    merged := true
    for merged {
        merged = false
        // We need to check each pair of graphs for potential merging
        for i := 0; i < len(graphs); i++ {
            for j := i + 1; j < len(graphs); j++ {
                if CanMerge(graphs[i], graphs[j]) {
                    // Merge the two graphs and replace them in the slice
                    mergedGraph := MergeGraphs(graphs[i], graphs[j])

                    // Remove the merged graphs and add the new merged graph
                    // graphs = append(graphs[:i], graphs[i+1:j]...)
                    // graphs = append(graphs[:j-1], graphs[j+1:]...)

                    graphs = slices.Delete(graphs, i, i + 1)
                    graphs = slices.Delete(graphs, j - 1, j)
                    graphs = append(graphs, mergedGraph)

                    // Set merged to true to continue the process
                    merged = true
                    break
                }
            }
            if merged {
                break
            }
        }
    }

    return graphs
}

func is_box_in_circuits(box BoxCoord, circuits []Graph) (bool, int) {
    for c_idx, c := range circuits {
        if _, exists := c[box]; exists {
            return true, c_idx
        }
    }
    return false, -1
}

type Distance struct {
    box_pair [2]BoxCoord
    dist     uint
}

func Puzzle1() {
    fmt.Printf("Day %d, Puzzle 1!\n", DAY)

    boxes := get_box_coords()
    // fmt.Println("boxes:", boxes)

    var iter_count int
    if len(boxes) == 20 {
        iter_count = 10
    } else {
        iter_count = 1000
    }

    boxes_combinations := generateCombinationsOf2(boxes)
    distances := make([]Distance, len(boxes_combinations))
    for i, bc := range boxes_combinations {
        distances[i] = Distance{bc, calc_distance(bc[0], bc[1])}
    }
    slices.SortFunc(distances, func(a, b Distance) int {
        return cmp.Compare(a.dist, b.dist)
    })
    distances = distances[:iter_count]

    var circuits []Graph
    bar := progressbar.Default(int64(len(distances)))
    for _, dist := range distances {
        min_distance_coord := dist.box_pair

        mdc0_connected, mdc0_c_idx := is_box_in_circuits(min_distance_coord[0], circuits)
        mdc1_connected, mdc1_c_idx := is_box_in_circuits(min_distance_coord[1], circuits)
        if mdc0_connected && mdc1_connected {
            if mdc0_c_idx == mdc1_c_idx {
                bar.Add(1)
                continue
            } else {
                c0 := circuits[mdc0_c_idx]
                c1 := circuits[mdc1_c_idx]
                c0.AddLink(min_distance_coord[0], min_distance_coord[1])
                c1.AddLink(min_distance_coord[0], min_distance_coord[1])
                cm := MergeGraphs(c0, c1)

                if mdc0_c_idx < mdc1_c_idx {
                    circuits = slices.Delete(circuits, mdc0_c_idx, mdc0_c_idx + 1)
                    circuits = slices.Delete(circuits, mdc1_c_idx - 1, mdc1_c_idx)
                } else {
                    circuits = slices.Delete(circuits, mdc1_c_idx, mdc1_c_idx + 1)
                    circuits = slices.Delete(circuits, mdc0_c_idx - 1, mdc0_c_idx)
                }
                circuits = append(circuits, cm)
            }
        } else if mdc0_connected {
            c := circuits[mdc0_c_idx]
            c.AddLink(min_distance_coord[0], min_distance_coord[1])
            circuits[mdc0_c_idx] = c
        } else if mdc1_connected {
            c := circuits[mdc1_c_idx]
            c.AddLink(min_distance_coord[0], min_distance_coord[1])
            circuits[mdc1_c_idx] = c
        } else {
            c := make(Graph)
            c.AddLink(min_distance_coord[0], min_distance_coord[1])
            circuits = append(circuits, c)
        }

        circuits = MergeGraphsIfNecessary(circuits)
        bar.Add(1)
    }
    bar.Close()
    // fmt.Println("circuits:")
    // for _, c := range circuits {
    //     fmt.Println(c)
    // }

    mult_largest_3_circuits_size := 1
    for i := 0; i < 3; i++ {
        largest_circuit_len := 0
        largest_circuit_idx := 0
        for j, c := range circuits {
            c_len := len(c)
            if c_len > largest_circuit_len {
                largest_circuit_len = c_len
                largest_circuit_idx = j
            }
        }
        circuits = slices.Delete(circuits, largest_circuit_idx, largest_circuit_idx + 1)

        fmt.Println("largest_circuit_len:", largest_circuit_len)
        mult_largest_3_circuits_size *= largest_circuit_len
    }

    fmt.Println("mult_largest_3_circuits_size:", mult_largest_3_circuits_size)
}

func Puzzle2() {
    fmt.Printf("\nDay %d, Puzzle 2!\n", DAY)

    boxes := get_box_coords()
    // fmt.Println("boxes:", boxes)

    num_boxes := len(boxes)

    boxes_combinations := generateCombinationsOf2(boxes)
    distances := make([]Distance, len(boxes_combinations))
    for i, bc := range boxes_combinations {
        distances[i] = Distance{bc, calc_distance(bc[0], bc[1])}
    }
    slices.SortFunc(distances, func(a, b Distance) int {
        return cmp.Compare(a.dist, b.dist)
    })

    var circuits []Graph
    var last_combo [2]BoxCoord
    bar := progressbar.Default(int64(num_boxes - 1))
    for _, dist := range distances {
        min_distance_coord := dist.box_pair

        mdc0_connected, mdc0_c_idx := is_box_in_circuits(min_distance_coord[0], circuits)
        mdc1_connected, mdc1_c_idx := is_box_in_circuits(min_distance_coord[1], circuits)
        if mdc0_connected && mdc1_connected {
            if mdc0_c_idx == mdc1_c_idx {
                continue
            } else {
                c0 := circuits[mdc0_c_idx]
                c1 := circuits[mdc1_c_idx]
                c0.AddLink(min_distance_coord[0], min_distance_coord[1])
                c1.AddLink(min_distance_coord[0], min_distance_coord[1])
                cm := MergeGraphs(c0, c1)

                if mdc0_c_idx < mdc1_c_idx {
                    circuits = slices.Delete(circuits, mdc0_c_idx, mdc0_c_idx + 1)
                    circuits = slices.Delete(circuits, mdc1_c_idx - 1, mdc1_c_idx)
                } else {
                    circuits = slices.Delete(circuits, mdc1_c_idx, mdc1_c_idx + 1)
                    circuits = slices.Delete(circuits, mdc0_c_idx - 1, mdc0_c_idx)
                }
                circuits = append(circuits, cm)
            }
        } else if mdc0_connected {
            c := circuits[mdc0_c_idx]
            c.AddLink(min_distance_coord[0], min_distance_coord[1])
            circuits[mdc0_c_idx] = c
        } else if mdc1_connected {
            c := circuits[mdc1_c_idx]
            c.AddLink(min_distance_coord[0], min_distance_coord[1])
            circuits[mdc1_c_idx] = c
        } else {
            c := make(Graph)
            c.AddLink(min_distance_coord[0], min_distance_coord[1])
            circuits = append(circuits, c)
        }

        circuits = MergeGraphsIfNecessary(circuits)
        bar.Add(1)
        num_boxes--
        if num_boxes == 1 {
            // fmt.Println("min_distance_coord:", min_distance_coord)
            // fmt.Println("circuits:")
            // for _, c := range circuits {
            //     fmt.Println(c)
            // }
            // fmt.Println()
            last_combo = min_distance_coord
            break
        }
        // if i > 10 && len(circuits) == 1 {
        // }
    }
    bar.Close()
    // fmt.Println("circuits:", circuits)

    fmt.Println("last_combo:", last_combo)
    mult_x_last_boxes := last_combo[0][0] * last_combo[1][0]
    fmt.Println("mult_x_last_boxes:", mult_x_last_boxes)
}

