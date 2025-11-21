/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program tests if a Discount Nuts and Bolts order is OK and checks the price.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Constants
	const BOLT_COST int = 10
	const NUT_COST int = 5
	const WASHER_COST int = 3

	// Declaration
	var boltsAsString string
	var bolts int
	var nutsAsString string
	var nuts int
	var washersAsString string
	var washers int

	// Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many bolts would you like to purchase? ")
	boltsAsString, _ = reader.ReadString('\n')
	boltsAsString = strings.TrimSpace(boltsAsString)
	bolts, _ = strconv.Atoi(boltsAsString)

	fmt.Print("How many nuts would you like to purchase? ")
	nutsAsString, _ = reader.ReadString('\n')
	nutsAsString = strings.TrimSpace(nutsAsString)
	nuts, _ = strconv.Atoi(nutsAsString)

	fmt.Print("How many washers would you like to purchase? ")
	washersAsString, _ = reader.ReadString('\n')
	washersAsString = strings.TrimSpace(washersAsString)
	washers, _ = strconv.Atoi(washersAsString)

	// Cost
	var totalCost int = bolts*BOLT_COST + nuts*NUT_COST + washers*WASHER_COST

	fmt.Printf("Bolts: %d\n", bolts)
	fmt.Printf("Nuts: %d\n", nuts)
	fmt.Printf("Washers: %d\n", washers)

	// If statement
	if bolts > nuts {
		fmt.Println("Check the Order - not enough nuts for the bolts you purchased.")
	}
	if washers < bolts {
		fmt.Println("Check the Order - not enough washers for the bolts you purchased.")
	}
	if bolts <= nuts && washers >= bolts {
		fmt.Println("Order is OK.")
	}

	fmt.Printf("Your total cost of the order is %d cents.\n", totalCost)
}
