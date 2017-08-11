package life

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

type Universe struct {
	Generation, Height, Width int
	Space                     [][]int
}

func (uni *Universe) Live(i, j int) error {
	if i >= uni.Height || i < 0 || j >= uni.Width || j < 0 {
		return fmt.Errorf("(%d, %d) out of bounds", i, j)
	}
	uni.Space[i][j]++
	return nil
}

func (uni *Universe) Die(i, j int) error {
	if i >= uni.Height || i < 0 || j >= uni.Width || j < 0 {
		return fmt.Errorf("(%d, %d) out of bounds", i, j)
	}
	uni.Space[i][j] = 0
	return nil
}

func (uni *Universe) Get(i, j int) (int, error) {
	if i >= uni.Height || i < 0 || j >= uni.Width || j < 0 {
		return 0, fmt.Errorf("(%d, %d) out of bounds", i, j)
	}
	return uni.Space[i][j], nil
}

/**
 * [space description]
 * @type {[type]}
 */
func freshSpace(height, width int) [][]int {
	// Use make to allocate the slice of slices
	space := make([][]int, height)
	// Make a new zeroed row at each index
	for i := 0; i < height; i++ {
		space[i] = make([]int, width)
	}

	return space
}

/**
 * [space description]
 * @type {[type]}
 */
func NewUniverse(height, width int) *Universe {
	// Get some fresh space
	space := freshSpace(height, width)
	// Construct the universe
	uni := Universe{
		Generation: 0,
		Height: height,
		Width: width,
		Space:      space,
	}
	// Return a pointer to the new universe
	return &uni
}


/**
 * [func description]
 * @param  {[type]} uni [description]
 * @return {[type]}     [description]
 */
func (uni *Universe) RandomSeed(n int, r *rand.Rand) {
	// Randomly seed until n seeds are placed
	for i := 0; i < n; i++ {
		row := r.Intn(uni.Height)
		col := r.Intn(uni.Width)
		uni.Live(row, col)
	}
}

/**
 * [living description]
 * @type {[type]}
 */
func (uni *Universe) LiveNeighbors(i, j int) int {
	living := 0
	// Calculate the number of living neighbors in O(1) time
	for r := i - 1; r <= i+1; r++ {
		for c := j - 1; c <= j+1; c++ {
			// Ensure we aren't out of bounds and are not the central cell
			if (r >= 0 && r < uni.Height) && (c >= 0 && c < uni.Width) &&
				!(r == i && c == j) {
				cell, _ := uni.Get(r, c)
				if cell > 0 {
					living++
				}
			}
		}
	}

	return living
}

func (uni *Universe) String() string {
	// Declare a buffer to build the string with
	buf := new(bytes.Buffer)
	// Add starting curly and bracket
	buf.WriteString("{\"space\": [\n")
	for r := 0; r < uni.Height; r++ {
		// Add starting bracket and first column value
		buf.WriteString("\t[")
		buf.WriteString(strconv.Itoa(uni.Space[r][0]))
		for c := 1; c < uni.Width; c++ {
			// Add column val and
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(uni.Space[r][c]))
		}
		// Add ending bracket (and comma if needed)
		buf.WriteString("]")
		if r < uni.Height-1 {
			buf.WriteString(",\n")
		}
	}

	// Add ending bracket, comma, generation, and curly brace
	buf.WriteString("],\n\"generation\": ")
	buf.WriteString(strconv.Itoa(uni.Generation))
	buf.WriteString(",\n\"height\": ")
	buf.WriteString(strconv.Itoa(uni.Height))
	buf.WriteString(",\n\"width\": ")
	buf.WriteString(strconv.Itoa(uni.Width))
	buf.WriteString("}")

	return buf.String()
}


// func (uni *Universe) String() string {
// 	b, err := json.Marshal(uni)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	return string(b)
// }
