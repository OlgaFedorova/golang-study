package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село"},        // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик"},              // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Село":      100,
		"Миллионик": 500,
	}

	for _, city := range groupCity[10] {
		if _, inMap := cityPopulation[city]; inMap {
			delete(cityPopulation, city)
		}
	}

	for _, city := range groupCity[1000] {
		if _, inMap := cityPopulation[city]; inMap {
			delete(cityPopulation, city)
		}
	}
	fmt.Println(cityPopulation)
}
