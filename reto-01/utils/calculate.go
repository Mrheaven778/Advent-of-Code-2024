package utils

import (
	"sort"
	"strconv"
)

// Convierte un array de strings a un array de enteros
func ConvertToIntArray(strArray []string) ([]int, error) {
	intArray := make([]int, len(strArray))
	for i, str := range strArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intArray[i] = num
	}
	return intArray, nil
}

// Calcula las repeticiones de cada número en arr2
func calculateRepetitions(arr1, arr2 []int) map[int]int {
	repetitions := make(map[int]int)
	for _, num := range arr2 {
		repetitions[num]++
	}
	return repetitions
}

// Calcula el puntaje de similitud entre dos listas
func CalculateTotalDistance(arr1, arr2 []int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	repetitions := calculateRepetitions(arr1, arr2)
	similarityScore := 0

	for _, num := range arr1 {
		similarityScore += num * repetitions[num]
	}

	return similarityScore
}

// Calcula el valor absoluto de un número
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
