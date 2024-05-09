package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type weekDays int

const (
	Sunday weekDays = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (day weekDays) String() string {
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if day < Sunday || day > Saturday {
		return "Unknown"
	}
	return names[day]
}

type TimeEntry struct {
	ArrivalTime   time.Time
	DepartureTime time.Time
	TimeWorked    time.Duration
}

type Employee struct {
	TimeEntries     map[weekDays]TimeEntry
	TotalTimeWorked time.Duration
}

type Report map[string]*Employee

var timeLayout string = "15:04"
var defaultArrivalTime, _ = time.Parse(timeLayout, "09:00")
var defaultDepartureTime, _ = time.Parse(timeLayout, "17:00")
var workdayTime = defaultDepartureTime.Sub(defaultArrivalTime)

var report = Report{
	"John Doe": &Employee{
		TimeEntries: map[weekDays]TimeEntry{
			Monday: {
				ArrivalTime:   defaultArrivalTime,
				DepartureTime: defaultDepartureTime,
				TimeWorked:    workdayTime,
			},
		},
		TotalTimeWorked: workdayTime,
	},
	"Jane Smith": &Employee{
		TimeEntries: map[weekDays]TimeEntry{
			Tuesday: {
				ArrivalTime:   defaultArrivalTime,
				DepartureTime: defaultDepartureTime,
				TimeWorked:    workdayTime,
			},
		},
		TotalTimeWorked: workdayTime,
	},
}

func formatDuration(d time.Duration) string {
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	return fmt.Sprintf("%dh%dm", hours, minutes)
}

func printReport() {
	for name, employee := range report {
		fmt.Println("Employee:", name)
		for day, timeEntry := range employee.TimeEntries {
			fmt.Println("Day:", day, timeEntry.ArrivalTime.Format(timeLayout), "-", timeEntry.DepartureTime.Format(timeLayout), "Time worked:", formatDuration(timeEntry.TimeWorked))
		}
		fmt.Println("Week Total:", formatDuration(employee.TotalTimeWorked))
	}
}

func getWeekdayInput() weekDays {
	fmt.Println("Enter weekday: 0 - Sunday, 1 - Monday, 2 - Tuesday, 3 - Wednesday, 4 - Thursday, 5 - Friday, 6 - Saturday")
	var weekday int
	_, err := fmt.Scanln(&weekday)
	if err != nil {
		fmt.Println("\033[1;31mInvalid input. Please enter a number.\033[0m")
		return getWeekdayInput()
	}
	return weekDays(weekday)
}

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("\033[1;31mInvalid input. Please enter a string.\033[0m")
		return getInput(prompt)
	}
	return strings.TrimSpace(input)
}

func getTimeInput(prompt string) time.Time {
	for {
		input := getInput(prompt)
		parsedTime, err := time.Parse(timeLayout, input)
		if err != nil {
			fmt.Println("\033[1;31mInvalid input. Please enter time in format hh:mm.\033[0m")
			continue
		}
		return parsedTime
	}
}

func addEmployee(name string) {
	employee := Employee{
		TimeEntries: map[weekDays]TimeEntry{},
	}

	report[name] = &employee
}

func main() {
	for {
		for {
			printReport()
			fmt.Println("\033[1;36mChoose an option:\033[0m")
			fmt.Println("1) Add a new employee")
			fmt.Println("2) Add arrival/departure time(hh:mm)")
			fmt.Println("3) Show working hours report")
			fmt.Println("0) Exit")
			choice := getInput("Enter your choice:\n")

			switch choice {
			case "1":
				name := getInput("Enter employee name:\n")
				addEmployee(name)
			case "2":
				fmt.Println("Choose employee:")
				employees := make([]string, 0, len(report))
				for name := range report {
					employees = append(employees, name)
				}
				for i, name := range employees {
					fmt.Println("Index: ", i, "Name: ", name)
				}
				employeeIndexStr := getInput("Enter employee index:\n")
				employeeIndex, err := strconv.Atoi(employeeIndexStr)
				if err != nil || employeeIndex < 0 || employeeIndex >= len(employees) {
					fmt.Println("\033[1;31mInvalid input. Please enter a valid index.\033[0m")
					continue
				}
				employee := employees[employeeIndex]
				weekday := getWeekdayInput()
				var arrivalTime, departureTime time.Time
				for {
					arrivalTime = getTimeInput("Enter arrival time(hh:mm):\n")
					departureTime = getTimeInput("Enter departure time(hh:mm):\n")
					if departureTime.Before(arrivalTime) {
						fmt.Println("\033[1;31mEmployee can't leave before arrival.\033[0m")
						continue
					}
					break
				}

				report[employee].TimeEntries[weekday] = TimeEntry{
					ArrivalTime:   arrivalTime,
					DepartureTime: departureTime,
					TimeWorked:    departureTime.Sub(arrivalTime),
				}
				report[employee].TotalTimeWorked += report[employee].TimeEntries[weekday].TimeWorked
			case "3":
				printReport()
			case "0":
				fmt.Println("\033[1;31mExiting...\033[0m")
				return
			default:
				fmt.Println("\033[1;31mInvalid option chosen.\033[0m")
			}
		}

	}
}
