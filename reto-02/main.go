package main

import (
	"fmt"
	utils "mrheaven/ultils"
)

func main() {
	lines, err := utils.ReadFile("./ultils/datos.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	reports := utils.ParseReports(lines)
	safeCountP1, safeCountP2 := utils.CountSafeReports(reports)
	fmt.Println("Number of safe reports (Part 1):", safeCountP1)
	fmt.Println("Number of safe reports (Part 2):", safeCountP2)
}
