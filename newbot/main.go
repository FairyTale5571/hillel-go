package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := generateTarget(1, 100)
	fmt.Println("Target:", target)
	fmt.Println("Гра 'Вгадай число'!")
	fmt.Println("Комп'ютер загадав число від 1 до 100.")

	attempts := playGame(target)
	fmt.Printf("Вітаємо! Уявний гравець вгадав число за %d спроб.\n", attempts)
}

// generateTarget створює випадкове число в заданому діапазоні.
func generateTarget(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// playGame виконує логіку гри, де уявний гравець намагається вгадати число.
func playGame(target int) int {
	min := 1
	max := 100
	attempts := 0

	for {
		time.Sleep(3 * time.Second)
		guess := (min + max) / 2
		attempts++
		fmt.Printf("Уявний гравець вгадує: %d\n", guess)

		if guess < target {
			fmt.Println("Занадто мало.")
			min = guess + 1
			fmt.Println("guess:", guess)
		} else if guess > target {
			fmt.Println("Занадто багато.")
			max = guess - 1
			fmt.Println("guess:", guess)
		} else {
			break
		}
	}
	return attempts
}
