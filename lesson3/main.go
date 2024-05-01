package main

import (
	"fmt"

	"github.com/maksymshtarkberg/Calculating-Work-Time/app"
)

func main() {
	calcTime()
}

func calcTime() {
	workerManager := app.NewWorkerManager()

	for {
		fmt.Println("\nTimeManager")
		fmt.Println("\nOptions:")
		fmt.Println("[1] Add a worker")
		fmt.Println("[2] Set working hours")
		fmt.Println("[3] View workers")
		fmt.Println("[4] View workers hours")
		fmt.Println("[5] Exit")
		fmt.Print("\nChoose an option: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			app.HandleAddWorker(workerManager)
		case 2:
			app.HandleSetWorkingHours(workerManager)
		case 3:
			app.GetWorkers(workerManager)
		case 4:
			app.GetWorkerTime(workerManager)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
