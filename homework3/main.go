package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type WorkTimeLog struct {
	currentName string
	journal     map[string]*TimeRecord
}

type TimeRecord struct {
	WeekTime  int
	StartTime string
	EndTime   string
}

func (w *WorkTimeLog) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введіть ім'я працівника: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	w.nameRegistration(name)
	w.userOptions()

	for i, val := range w.journal {
		fmt.Println(i)
		fmt.Println(val)
	}
}

func (w *WorkTimeLog) startTimeRegistration() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введіть час приходу у форматі '09:10': ")

	inputTime, _ := reader.ReadString('\n')

	ok := w.inputTimeCheck(inputTime)
	if !ok {
		fmt.Println("Час введено некорректно, спробуйте знову")
		return false
	}

	w.journal[w.currentName].StartTime = inputTime

	return true
}

func (w *WorkTimeLog) endTimeRegistration() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введіть час закінчення роботи у форматі '09:10': ")

	inputTime, _ := reader.ReadString('\n')

	ok := w.inputTimeCheck(inputTime)

	if !ok {
		fmt.Println("Час введено некорректно, спробуйте знову")
		return false
	}

	w.journal[w.currentName].EndTime = inputTime

	return true
}

func (w *WorkTimeLog) calculateHours() {
	start, _ := time.Parse("15:04", w.journal[w.currentName].StartTime)
	end, _ := time.Parse("15:04", w.journal[w.currentName].EndTime)
	duration := end.Sub(start).Hours()
	if duration < 0 {
		duration += 24
	}
	//fmt.Println("startTime: " + start.Format("15:04"))
	//fmt.Println("endTime: " + end.Format("15:04"))
	//fmt.Println(start)
	//fmt.Println(w.journal[w.currentName].StartTime)
	fmt.Printf("Відпрацьовано за день: %.2f годин\n", duration)
	w.journal[w.currentName].WeekTime += int(duration)
}

func (w *WorkTimeLog) inputTimeCheck(inputTime string) bool {
	inputTime = strings.TrimSpace(inputTime)

	_, err := time.Parse("15:04", inputTime)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func (w *WorkTimeLog) userOptions() {
	for {
		fmt.Println("\n1. Зареєструвати час початку робочого дня")
		fmt.Println("2. Час закінчення робочого дня")
		fmt.Println("3. Всього відпрацьовано за день")
		fmt.Println("4. Всього відпрацьовано за тижлдень")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			ok := w.startTimeRegistration()
			if !ok {
				continue
			}
		case 2:
			ok := w.endTimeRegistration()
			if !ok {
				continue
			}
		case 3:
			w.calculateHours()
		case 4:

		case 5:

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}

func (w *WorkTimeLog) nameRegistration(name string) {
	_, ok := w.journal[name]

	if !ok {
		w.journal[name] = &TimeRecord{}
	}
	w.currentName = name
}

func main() {
	workLogger := WorkTimeLog{journal: make(map[string]*TimeRecord)}
	workLogger.Run()
}
