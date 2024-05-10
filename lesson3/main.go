package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// This is main function
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введіть ім'я співробітника: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Помилка вводу")
			continue
		}
		name = strings.TrimSpace(name)
		weekHours := 0.0
		daysOfWeek := []string{"Понеділок", "Вівторок", "Середа", "Четвер", "П'ятниця", "Субота", "Неділя"}
		fmt.Println("Якщо співробітник не прийшов в цей день на роботу введіть ʼ-ʼ")
		for i := 0; i < len(daysOfWeek); i++ {
			for {
				fmt.Printf("Коли співробітник прийшов на роботу в %s? Введіть в (ГГ:ХХ): ", daysOfWeek[i])
				comelInput, _ := reader.ReadString('\n')
				if strings.TrimSpace(strings.TrimSpace(comelInput)) == "-" {
					break
				}
				comeTime, err := time.Parse("15:04", strings.TrimSpace(comelInput))
				if err != nil {
					fmt.Println("Помилка вводу")
					continue
				}

				fmt.Printf("Коли співробітник пішов з роботи в %s? Введіть в (ГГ:ХХ): ", daysOfWeek[i])
				wentInput, _ := reader.ReadString('\n')
				wentTime, err := time.Parse("15:04", strings.TrimSpace(wentInput))
				if err != nil {
					fmt.Println("Помилка вводу")
					continue
				}

				if wentTime.Before(comeTime) {
					fmt.Printf("В %s співробітник не міг піти раніще, ніж прийшов!\n", daysOfWeek[i])
					continue
				}
				hoursWorked := wentTime.Sub(comeTime).Hours()
				fmt.Printf("%s працював(ла) %.2f годин в %s.\n", name, hoursWorked, daysOfWeek[i])
				weekHours += hoursWorked
				break
			}
		}
		fmt.Printf("Кількість годин, відпрацьованих %s за тиждень: %.2f\n", name, weekHours)

		fmt.Print("Бажаєте підрахувати години іншого співробітника? (так/ні): ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Помилка при читанні вводу:", err)
			break
		}
		if strings.TrimSpace(response) != "так" {
			break
		}

	}
}
