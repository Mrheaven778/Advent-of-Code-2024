package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReadFile reads the content of a file and returns it as a slice of strings
func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Convierte una cadena de informes a una matriz de enteros
func ParseReports(lines []string) [][]int {
	reports := make([][]int, len(lines))

	for i, line := range lines {
		levels := strings.Fields(line)
		report := make([]int, len(levels))
		for j, level := range levels {
			num, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
			report[j] = num
		}
		reports[i] = report
	}

	return reports
}

// Verifica si un informe es seguro
func isReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	orderedClone := make([]int, len(levels))
	copy(orderedClone, levels)
	sort.Ints(orderedClone)

	reversedClone := make([]int, len(orderedClone))
	copy(reversedClone, orderedClone)
	for i, j := 0, len(reversedClone)-1; i < j; i, j = i+1, j-1 {
		reversedClone[i], reversedClone[j] = reversedClone[j], reversedClone[i]
	}

	if !equal(levels, orderedClone) && !equal(levels, reversedClone) {
		return false
	}

	set := make(map[int]struct{})
	for _, v := range levels {
		set[v] = struct{}{}
	}
	if len(set) != len(levels) {
		return false
	}

	for i := 0; i < len(levels)-1; i++ {
		if abs(levels[i]-levels[i+1]) > 3 {
			return false
		}
	}

	return true
}

// Verifica si dos slices son iguales
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Verifica si un informe es seguro eliminando un solo nivel
func isReportSafeWithDampener(levels []int) bool {
	if isReportSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		modifiedReport := append(levels[:i], levels[i+1:]...)
		if isReportSafe(modifiedReport) {
			return true
		}
	}

	return false
}

// Calcula el valor absoluto de un número
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Cuenta cuántos informes son seguros
func CountSafeReports(reports [][]int) (int, int) {
	safeCountP1 := 0
	safeCountP2 := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeCountP1++
			safeCountP2++
		} else if isReportSafeWithDampener(report) {
			safeCountP2++
		}
	}
	return safeCountP1, safeCountP2
}
