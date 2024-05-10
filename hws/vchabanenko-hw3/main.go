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
	journal     map[string]map[int]*Week
}

type Week struct {
	Days      [7]*Day
	WeekHours float64
}

type Day struct {
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

	w.ensureWeekAndDay()
	w.journal[w.currentName][w.currentWeekNumber()].Days[time.Now().Weekday()].StartTime = inputTime

	hours, _ := w.calculateHours()
	w.addWeekHours(hours)

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

	w.ensureWeekAndDay()
	w.journal[w.currentName][w.currentWeekNumber()].Days[time.Now().Weekday()].EndTime = inputTime
	hours, _ := w.calculateHours()
	w.addWeekHours(hours)

	return true
}

func (w *WorkTimeLog) calculateHours() (float64, error) {
	startTime := w.journal[w.currentName][w.currentWeekNumber()].Days[time.Now().Weekday()].StartTime
	endTime := w.journal[w.currentName][w.currentWeekNumber()].Days[time.Now().Weekday()].EndTime

	if startTime == "" || endTime == "" {
		return 0, fmt.Errorf("час початку або завершення робочого дня не зареєстровано")
	}

	start, errStart := time.Parse("15:04", strings.TrimSpace(startTime))
	end, errEnd := time.Parse("15:04", strings.TrimSpace(endTime))
	if errStart != nil || errEnd != nil {
		return 0, fmt.Errorf("невдалий парсинг часу: %v, %v", errStart, errEnd)
	}

	duration := end.Sub(start).Hours()
	if duration < 0 {
		duration += 24
	}

	return duration, nil
}

func (w *WorkTimeLog) dayHoursDisplay() {
	duration, err := w.calculateHours()
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	fmt.Printf("Відпрацьовано за день: %.2f годин\n", duration)
}

func (w *WorkTimeLog) addWeekHours(duration float64) {
	w.journal[w.currentName][w.currentWeekNumber()].WeekHours += duration
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
		fmt.Println("4. Всього відпрацьовано за тиждень")
		fmt.Println("5. Exit")
		fmt.Print("Оберіть опцію: ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Невірний ввід. Спробуйте ще")
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
			w.dayHoursDisplay()
		case 4:
			w.weekTimeDisplay()
		case 5:
			fmt.Println("Вихід з програми.")
			return
		default:
			fmt.Println("Невірна опція. Спробуйте ще раз.")
		}
	}

}

func (w *WorkTimeLog) nameRegistration(name string) {
	_, ok := w.journal[name]

	if !ok {
		w.journal[name] = make(map[int]*Week)
	}

	w.currentName = name
}

func (w *WorkTimeLog) weekTimeDisplay() {
	fmt.Printf("Всього відпрацьовано за тиждень: %0.2f", w.journal[w.currentName][w.currentWeekNumber()].WeekHours)
}

func (w *WorkTimeLog) currentWeekNumber() int {
	_, weekNumber := time.Now().ISOWeek()

	return weekNumber
}

func (w *WorkTimeLog) ensureWeekAndDay() {
	if w.journal[w.currentName] == nil {
		w.journal[w.currentName] = make(map[int]*Week)
	}
	if w.journal[w.currentName][w.currentWeekNumber()] == nil {
		w.journal[w.currentName][w.currentWeekNumber()] = &Week{}
	}
	dayIndex := time.Now().Weekday()
	if w.journal[w.currentName][w.currentWeekNumber()].Days[dayIndex] == nil {
		w.journal[w.currentName][w.currentWeekNumber()].Days[dayIndex] = &Day{}
	}
}

func main() {
	workLogger := WorkTimeLog{journal: make(map[string]map[int]*Week)}
	workLogger.Run()
}
