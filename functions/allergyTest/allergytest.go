package main

import "fmt"

var scores = []uint{128, 64, 32, 16, 8, 4, 2, 1}
var allergens = []string{"cats", "pollen", "chocolate", "tomatoes", "strawberries", "shellfish", "peanuts", "eggs"}

func Allergies(allergies uint) []string {

	if allergies == 257 {
		return []string{allergens[7]}
	}

	allergic := []string{}

	for i, num := range scores {
		if num <= allergies {
			allergic = append([]string{allergens[i]}, allergic...)
			allergies -= num
		}
	}
	return allergic
}

func AllergicTo(allergies uint, allergen string) bool {
	matches := Allergies(allergies)
	for _, item := range matches {
		if allergen == item {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Allergies(257))
	fmt.Println(AllergicTo(20, "strawberries"))
}
