// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

var limitFlag = pflag.Float64P("limit", "l", -1, `When the sum of the inputs
reaches this limit, program execution is stopped.`)

// Add creates a running total of items in exp, and returns total and total items.
func scanAndAdd(limit float64) (float64, int) {
	var total float64
	var count int
	// var args []float64
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		if t == "" {
			break
		}
		f, err := strconv.ParseFloat(t, 64)
		if err != nil {
			panic(err)
			// return nil
		}
		// args = append(args, float64(f))
		if limit > 0 && (total+f) > limit {
			return total, count
		}
		total += f
		count++

	}
	// for _, n := range exp {
	// 	total += n
	// }
	return total, count
}

/*
func bufioScan() []float64 {
	var args []float64
	var input = bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		if t == "" {
			break
		}
		f, err := strconv.ParseFloat(t, 64)
		if err != nil {
			panic(err)
			// return nil
		}
		args = append(args, float64(f))

	}
	return args
}
*/

func main() {
	pflag.Parse()
	fmt.Println("Enter operand to add, pressing ENTER after each one. Press ENTER with no operand to end.")
	if int(*limitFlag) <= 0 {
		total, count := scanAndAdd(0)
		fmt.Printf("Total Amount: %.2f, \n Number of items: %d\n", total, count)
	}
	total, count := scanAndAdd(*limitFlag)
	fmt.Printf("Total Amount: %.2f, \n Number of items: %d\n", total, count)

}
