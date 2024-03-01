package main

import (
	"errors"
	"fmt"
)

// WARMUP
func main_spiral() {
	fmt.Println("Starting...")

	n := 4
	spiral, err := createSpiral(n)
	if err != nil {
		fmt.Printf("Encountered error: %s", err.Error())
		return
	}

	printSpiral(spiral)
	fmt.Println("\nDone.")
}

func createSpiral(n int) ([][]int, error) {
	// exit if n invalid
	if n <= 0 {
		return nil, errors.New("invalid value for n, should be > 0")
	}

	// initialize some array
	spiral := make([][]int, n)
	for i := 0; i < n; i++ {
		spiral[i] = make([]int, n)
	}

	// initialize the bouncing indexes, these will change as we spiral
	bounce_right := n - 1
	bounce_bottom := n - 1
	bounce_left := 0
	bounce_top := 0
	current_row := 0
	current_col := 0

	// done when counter == n^2
	counter := 0

	// go ahead and start spiraling

	for counter < n*n {
		// go right
		for i := current_col; i <= bounce_right && counter < n*n; i++ {
			spiral[current_row][i] = counter
			if i <= bounce_right-1 || counter == n*n-1 {
				counter++
			}
			current_col = i
		}
		bounce_top++

		// go down
		for i := current_row; i <= bounce_bottom && counter < n*n; i++ {
			spiral[i][current_col] = counter
			if i <= bounce_bottom-1 || counter == n*n-1 {
				counter++
			}
			current_row = i
		}
		bounce_right--

		// go left
		for i := current_col; i >= bounce_left && counter < n*n; i-- {
			spiral[current_row][i] = counter
			if i >= bounce_left+1 || counter == n*n-1 {
				counter++
			}
			current_col = i
		}
		bounce_bottom--

		// go up
		for i := current_row; i >= bounce_top && counter < n*n; i-- {
			spiral[i][current_col] = counter
			if i >= bounce_top+1 || counter == n*n-1 {
				counter++
			}
			current_row = i
		}
		bounce_left++
	}

	return spiral, nil
}

func printSpiral(spiral [][]int) {
	rows := len(spiral)
	cols := 0
	if rows > 0 {
		cols = len(spiral[0])
		fmt.Printf("rows: %d, cols: %d \n", rows, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", spiral[i][j])
		}
		fmt.Println()
	}
}
