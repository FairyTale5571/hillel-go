package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TimeAtWork struct {
	EnterTime float64
	OutTime   float64
}

type WorkSchedule struct {
	Monday    TimeAtWork
	Tuesday   TimeAtWork
	Wednesday TimeAtWork
	Thursday  TimeAtWork
	Friday    TimeAtWork
	Saturday  TimeAtWork
	Sunday    TimeAtWork
}

type Worker struct {
	ID        int
	Name      string
	Surname   string
	WorkHours WorkSchedule
}

type WorkerManager struct {
	workers map[int]Worker
}

func NewWorkerManager() *WorkerManager {
	return &WorkerManager{
		workers: make(map[int]Worker),
	}
}

func (wm *WorkerManager) AddWorker(name, surname string) {
	id := len(wm.workers) + 1
	worker := Worker{
		ID:      id,
		Name:    name,
		Surname: surname,
	}
	wm.workers[id] = worker
}

func HandleAddWorker(wm *WorkerManager) {
	var name string
	var surname string

	fmt.Print("Enter employee name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter employee surname: ")
	surname, _ = reader.ReadString('\n')
	surname = strings.TrimSpace(surname)

	if name == "" || surname == "" {
		fmt.Print("\nName or surname cannot be empty\n")
		return
	}

	wm.AddWorker(name, surname)
}

func (wm *WorkerManager) SetWorkingHours(id, day int, enterTime, outTime float64) {
	worker, exists := wm.workers[id]
	if !exists {
		return
	}
	switch day {
	case 1:
		worker.WorkHours.Monday.EnterTime = enterTime
		worker.WorkHours.Monday.OutTime = outTime
	case 2:
		worker.WorkHours.Tuesday.EnterTime = enterTime
		worker.WorkHours.Tuesday.OutTime = outTime
	case 3:
		worker.WorkHours.Wednesday.EnterTime = enterTime
		worker.WorkHours.Wednesday.OutTime = outTime
	case 4:
		worker.WorkHours.Thursday.EnterTime = enterTime
		worker.WorkHours.Thursday.OutTime = outTime
	case 5:
		worker.WorkHours.Friday.EnterTime = enterTime
		worker.WorkHours.Friday.OutTime = outTime
	case 6:
		worker.WorkHours.Saturday.EnterTime = enterTime
		worker.WorkHours.Saturday.OutTime = outTime
	case 7:
		worker.WorkHours.Sunday.EnterTime = enterTime
		worker.WorkHours.Sunday.OutTime = outTime
	default:
		fmt.Println("Invalid day")
		return
	}
	wm.workers[id] = worker
}

func HandleSetWorkingHours(wm *WorkerManager) {
	var id, day int
	var enterTimeStr, outTimeStr string

	fmt.Print("Enter id of the employee:")
	fmt.Scanln(&id)

	if _, ok := wm.workers[id]; !ok {
		fmt.Printf("Employee with id %d not found\n", id)
		return
	}

	fmt.Print("Enter the day (1 for Monday, 2 for Tuesday, etc.):")
	fmt.Scanln(&day)

	if day <= 0 || day > 7 {
		fmt.Println("Invalid day")
		return
	}

	fmt.Print("Enter the time the employee started work (hh:mm):")
	fmt.Scanln(&enterTimeStr)

	fmt.Print("Enter the time the employee finished work (hh:mm):")
	fmt.Scanln(&outTimeStr)

	enterHour, enterMinute, err := parseTime(enterTimeStr)
	if err != nil {
		fmt.Println("Invalid time format:", err)
		return
	}

	outHour, outMinute, err := parseTime(outTimeStr)
	if err != nil {
		fmt.Println("Invalid time format:", err)
		return
	}

	enterTime := float64(enterHour) + float64(enterMinute)/60
	outTime := float64(outHour) + float64(outMinute)/60

	wm.SetWorkingHours(id, day, enterTime, outTime)
}

func parseTime(timeStr string) (hour, minute int, err error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid time format")
		return
	}

	hour, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	minute, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}

	if hour < 0 || hour >= 24 || minute < 0 || minute >= 60 {
		err = fmt.Errorf("invalid time values")
		return
	}

	return
}

func GetWorkers(wm *WorkerManager) {
	if len(wm.workers) == 0 {
		fmt.Println("\nNo workers found")
		return
	}
	fmt.Println("\nOur staff:")
	for _, worker := range wm.workers {
		fmt.Printf("ID: %d, Name: %s, Surname: %s\n", worker.ID, worker.Name, worker.Surname)
	}
}

func GetWorkerTime(wm *WorkerManager) {
	var id int
	var totalHours float64

	fmt.Print("Enter id of the employee:")
	fmt.Scanln(&id)

	worker, exists := wm.workers[id]
	if !exists {
		fmt.Println("Employee with id", id, "not found")
		return
	}

	fmt.Printf("\nWorking week for %s %s\n", worker.Name, worker.Surname)

	totalHours = 0
	workDays := []TimeAtWork{
		wm.workers[id].WorkHours.Monday,
		wm.workers[id].WorkHours.Tuesday,
		wm.workers[id].WorkHours.Wednesday,
		wm.workers[id].WorkHours.Thursday,
		wm.workers[id].WorkHours.Friday,
		wm.workers[id].WorkHours.Saturday,
		wm.workers[id].WorkHours.Sunday,
	}

	daysOfWeek := []string{"Monday\n", "Tuesday\n", "Wednesday\n", "Thursday\n", "Friday\n", "Saturday\n", "Sunday\n"}

	for i, day := range workDays {
		enterHour := int(day.EnterTime)
		enterMinute := int((day.EnterTime - float64(enterHour)) * 60)
		outHour := int(day.OutTime)
		outMinute := int((day.OutTime - float64(outHour)) * 60)

		fmt.Printf("%sEnter to work: %02d:%02d, Out of work: %02d:%02d\n", daysOfWeek[i], enterHour, enterMinute, outHour, outMinute)

		if day.EnterTime > day.OutTime {
			totalHours += float64(outHour) + float64(outMinute)/60 - (float64(enterHour) + float64(enterMinute)/60) + 24.0
		} else {
			totalHours += day.OutTime - day.EnterTime
		}
	}

	totalHoursInt := int(totalHours)
	totalMinutes := int((totalHours - float64(totalHoursInt)) * 60)

	fmt.Printf("Total working hours for the week: %02d:%02d\n", totalHoursInt, totalMinutes)
	fmt.Println()

}
