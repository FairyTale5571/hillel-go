package main

import (
	"fmt"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks = make(map[int]Task)
var idCounter = 1

// Нове завдання (функція)
func addTask(title string) {
	tasks[idCounter] = Task{ID: idCounter, Title: title, Completed: false}
	fmt.Printf("Завдання \"%s\" додано з ID %d\n", title, idCounter)
	idCounter++
}

// Мінус завдання за ID (функція)
func deleteTask(id int) {
	_, exists := tasks[id]
	if exists {
		delete(tasks, id)
		fmt.Printf("Завдання з ID \"%d\" успішно видалено", id)
		return
	}
	fmt.Println("Завдання з таким ID не існує!")
}

// Перегляд усіх завдань (метод)
func (t Task) viewTasks() {
	fmt.Println("Усі завдання:")
	for _, task := range tasks {
		status := "не завершене"
		if task.Completed {
			status = "завершене"
		}
		fmt.Printf("ID: %d, Завдання: %s, Статус: %s\n", task.ID, task.Title, status)
	}
}

// Завершененя завдання (метод)
func (t *Task) markAsCompleted() {
	if t.Completed {
		fmt.Println("Це завдання вже виконане.")
		return
	}
	t.Completed = true
	fmt.Println("Завдання успішно позначене як завершене.")
}

func main() {
	for {
		fmt.Println("Виберіть дію:")
		fmt.Println("1 - Додати нове завдання")
		fmt.Println("2 - Видалити завдання")
		fmt.Println("3 - Перегляд усіх завданнь")
		fmt.Println("4 - Завершення завдання")
		fmt.Println("5 - Вийти")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Помилка вводу!")
			continue
		}
		switch choice {
		case 1:
			fmt.Println("Введіьть назву завдання")
			var title string
			_, err := fmt.Scan(&title)
			if err != nil {
				return
			}
			addTask(title)
		case 2:
			fmt.Println("ВВедіть ID завдання")
			var id int
			_, _ = fmt.Scan(&id)
			deleteTask(id)
		case 3:
			var task Task
			task.viewTasks()
		case 4:
			fmt.Println("Виберіть завдання з яким ID бажаєте завершити:")
			var id int
			_, err := fmt.Scan(&id)
			if err != nil {
				println("Такого завдання не існує!")
				return
			}
			task := tasks[id]
			task.markAsCompleted()
			tasks[id] = task
		case 5:
			return
		default:
			fmt.Println("Невідома команда!")
		}

	}
}
