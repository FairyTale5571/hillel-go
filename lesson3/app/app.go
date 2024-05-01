package app

import (
	"bufio"
	"fmt"
	"os"
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
	var enterTime, outTime float64

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

	fmt.Print("Enter the time the employee started work:")
	fmt.Scanln(&enterTime)
	fmt.Print("Enter the time the employee finished work:")
	fmt.Scanln(&outTime)

	if enterTime < 0 || outTime < 0 {
		fmt.Println("Negative time is not allowed")
		return
	}

	wm.SetWorkingHours(id, day, enterTime, outTime)
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
		fmt.Printf("%sEnter to work: %.2f, Out of work: %.2f\n", daysOfWeek[i], day.EnterTime, day.OutTime)

		if day.EnterTime > day.OutTime {
			totalHours += 24 - day.EnterTime + day.OutTime
		} else {
			totalHours += day.OutTime - day.EnterTime
		}
	}

	fmt.Printf("Total working hours for the week: %.2f\n", totalHours)
	fmt.Println()

}
