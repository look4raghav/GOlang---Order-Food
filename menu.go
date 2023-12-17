package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	{1, "Adrakh Chai", 15},
	{2, "Filter Coffee", 25},
	{3, "Chhas", 40},
	{4, "Lassi", 35},
	{5, "Mango Lassi", 50},
	{6, "Kadi Chaawal", 40},
	{7, "Chhole Bhature", 90},
	{8, "Khasta Kachori", 20},
	{9, "Raj Kachori", 110},
	{10, "Veg. Sandwich", 30},
	{11, "Veg. Masala Maggi", 50.00},
	{12, "Samosa", 10},
	{13, "Cream Roll", 10},
}

func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}
