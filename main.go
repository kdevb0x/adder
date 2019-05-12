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

var (
	// set upper limit on totals
	limitFlag = pflag.Float64P("limit", "l", -1, `When the sum of the inputs
reaches this limit, program execution is stopped.`)

	// echo current totals afer each input
	echoFlag = pflag.BoolP("print", "p", false, `Print current total after
each input.`)

	// take user input using fmt pkg instead of bufio
	fmtFlag = pflag.BoolP("alt", "a", false, `Use an alternative input
scanner (userfull for debugging).`)

	// reuse the same input line instead of scrolling
	noscroll = pflag.BoolP("noscroll", "n", false, `Don't add newline after
	each user input, instead reuse the same line`)
)

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
		if *echoFlag {
			fmt.Printf("Current total: %.2f\n", total)
		}
		if *noscroll {
			print("\033[1A")
			print("\033[K")
		}
	}
	// for _, n := range exp {
	// 	total += n
	// }
	return total, count
}

func fmtScan(limit float64) (total float64, count int) {
	var input float64
	for {
		// (kdd) TODO:
		// fmt.Scanf panics if input is a blank newline, which is our
		// signal to finish looking for input There has to be a better
		// way of doing this than catching the panic...
		defer func() {
			if err := recover(); err == "unexpected newline" {
				os.Exit(0)
			}
		}()
		_, err := fmt.Scanf("%f", &input)
		if err != nil {
			panic(err)
		}
		if input == 0 {
			break
		}
		if limit > 0 && (total+input) > limit {
			return total, count
		}
		total += input
		count++
		if *echoFlag {
			fmt.Printf("Current total: %.2f\n", total)
		}
		if *noscroll {
			print("\033[1A")
			print("\033[K")
		}
	}
	return
}

func main() {
	pflag.Parse()
	fmt.Println("Enter operand to add, pressing ENTER after each one. Press ENTER with no operand to end.")
	if int(*limitFlag) <= 0 {
		if *fmtFlag {
			total, count := fmtScan(0)

			fmt.Printf("Total Amount: %.2f, \n Number of items: %d\n", total, count)
			os.Exit(0)
		}
		total, count := scanAndAdd(0)
		fmt.Printf("Total Amount: %.2f, \n Number of items: %d\n", total, count)
		os.Exit(0)

	}
	if *fmtFlag {
		total, count := fmtScan(*limitFlag)
		fmt.Printf("Total Amount: %.2f, \n Number of items: %d\n", total, count)
		os.Exit(0)
	}
	total, count := scanAndAdd(*limitFlag)
	fmt.Printf("Total Amount: %.2f, \n Number of items: %d\n", total, count)

}
