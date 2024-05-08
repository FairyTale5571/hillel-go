package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func handleDaysOfWeek(reader *bufio.Reader, index int, name string) float64 {
	fmt.Printf("Коли співробітник прийшов на роботу в %s? Введіть в (ГГ:ХХ): ", daysOfWeek[index])
	comeInput, _ := reader.ReadString('\n')
	comeInput = strings.TrimSpace(comeInput)
	if comeInput == "-" {
		return 0
	}
	comeTime, err := time.Parse("15:04", comeInput)
	if err != nil {
		fmt.Println("Помилка вводу")
		return 0
	}

	fmt.Printf("Коли співробітник пішов з роботи в %s? Введіть в (ГГ:ХХ): ", daysOfWeek[index])
	wentInput, _ := reader.ReadString('\n')
	wentInput = strings.TrimSpace(wentInput)
	wentTime, err := time.Parse("15:04", wentInput)
	if err != nil {
		fmt.Println("Помилка вводу")
		return 0
	}

	if wentTime.Before(comeTime) {
		fmt.Printf("В %s співробітник не міг піти раніще, ніж прийшов!\n", daysOfWeek[index])
		return 0
	}
	hoursWorked := wentTime.Sub(comeTime).Hours()
	fmt.Printf("%s працював(ла) %.2f годин в %s.\n", name, hoursWorked, daysOfWeek[index])
	return hoursWorked
}

func handleInput(reader *bufio.Reader) {
	fmt.Print("Введіть ім'я співробітника: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Помилка вводу")
		return
	}
	name = strings.TrimSpace(name)
	weekHours := 0.0

	fmt.Println("Якщо співробітник не прийшов в цей день на роботу введіть ʼ-ʼ")
	for i := 0; i < len(daysOfWeek); i++ {
		weekHours += handleDaysOfWeek(reader, i, name)
	}
	fmt.Printf("Кількість годин, відпрацьованих %s за тиждень: %.2f\n", name, weekHours)

	fmt.Print("Бажаєте підрахувати години іншого співробітника? (так/ні): ")
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Помилка при читанні вводу:", err)
		return
	}
	if strings.TrimSpace(response) != "так" {
		return
	}
}

var daysOfWeek = []string{"Понеділок", "Вівторок", "Середа", "Четвер", "П'ятниця", "Субота", "Неділя"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		handleInput(reader)
	}
}
