package main

import (
	"log"
	"mrheaven/file"
	"mrheaven/utils"
)

func main() {
	// Lee el archivo columna1.txt y almacena los datos en un array
	columna1, err := file.ReadFileToArray("./file/columna1.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lee el archivo columna2.txt y almacena los datos en un array
	columna2, err := file.ReadFileToArray("./file/columna2.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convierte los arrays de strings a arrays de enteros
	intArray1, err := utils.ConvertToIntArray(columna1)
	if err != nil {
		log.Fatal(err)
	}
	intArray2, err := utils.ConvertToIntArray(columna2)
	if err != nil {
		log.Fatal(err)
	}
	totalDistance := utils.CalculateTotalDistance(intArray1, intArray2)

	log.Printf("La distancia total entre los pares de números más pequeños es: %d", totalDistance)

}
