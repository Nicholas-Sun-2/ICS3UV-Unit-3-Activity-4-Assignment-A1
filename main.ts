/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program tests if a Discount Nuts and Bolts order is OK and checks the price.
 */

// Constants
const BOLT_COST: number = 10;
const NUT_COST: number = 5;
const WASHER_COST: number = 3;

// Input
const bolts: number = parseInt(prompt("How many bolts would you like to purchase?") || "0");
const nuts: number = parseInt(prompt("How many nuts would you like to purchase?") || "0");
const washers: number = parseInt(prompt("How many washers would you like to purchase?") || "0");

// Cost
const totalCost: number = bolts * BOLT_COST + nuts * NUT_COST + washers * WASHER_COST;

console.log(`Bolts: ${bolts}`);
console.log(`Nuts: ${nuts}`);
console.log(`Washers: ${washers}`);

// If statement
if (bolts > nuts) {
  console.log("Check the Order - not enough nuts for the bolts you purchased.");
}
if (washers < bolts) {
  console.log("Check the Order - not enough washers for the bolts you purchased.");
}
if (bolts <= nuts && washers >= bolts) {
  console.log("Order is OK.");
}

console.log(`Your total cost of the order is ${totalCost} cents.`);
