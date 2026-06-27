package main

import (
	"fmt"
	"sort"
)

var Nextid = 3

type Task struct {
	Title       string
	Description string
	Assignee    string
	Deadline    string
	Priority    string
	Status      string
}

var t1 = Task{
	Title:       "Task 1",
	Description: "Description of Task 1",
	Assignee:    "John Doe",
	Deadline:    "17 июля",
	Priority:    "High",
	Status:      "In Progress",
}
var t2 = Task{
	Title:       "Task 2",
	Description: "Description of Task 2",
	Assignee:    "Jane Smith",
	Deadline:    "30 июля",
	Priority:    "Medium",
	Status:      "Not Started",
}

var Tasklist = map[int]Task{
	1: t1,
	2: t2,
}

func viewTasks() { //Функция
	fmt.Println("Вот все задачи: ")
	for id, task := range Tasklist {
		fmt.Printf("ID: %d, Название: %s, Описание: %s, Исполнитель: %s, Срок выполнения: %s, Приоритет: %s, Статус: %s\n", id, task.Title, task.Description, task.Assignee, task.Deadline, task.Priority, task.Status)
	}
}
func addTask() { //Функция добавляющая таск
	var task Task
	fmt.Println("Название:")
	fmt.Scanln(&task.Title)
	fmt.Println("Описание:")
	fmt.Scanln(&task.Description)
	fmt.Println("Исполнитель:")
	fmt.Scanln(&task.Assignee)
	fmt.Println("Срок выполнения:")
	fmt.Scanln(&task.Deadline)
	fmt.Println("Приоритет:")
	fmt.Scanln(&task.Priority)
	fmt.Println("Статус:")
	fmt.Scanln(&task.Status)
	Tasklist[Nextid] = task
	Nextid++
}
func editTask() { //Функция изменяющая таск
	var id int
	var choise int
	var newValue string

	fmt.Println("Введите ID задачи, которую хотите отредактировать:")
	fmt.Scanln(&id)

	task, ok := Tasklist[id]
	if ok {
		fmt.Printf("Название: %s,\nОписание: %s,\nИсполнитель: %s, Срок выполнения: %s,\nПриоритет: %s, Статус: %s\n", task.Title, task.Description, task.Assignee, task.Deadline, task.Priority, task.Status)

		fmt.Printf("Выберите какой пункт вы хотите отредактировать:\n1. Название\n2. Описание\n3. Исполнитель\n4. Срок выполнения\n5. Приоритет\n6. Статус\n")
		fmt.Scanln(&choise)
		fmt.Println("Введите новое содержимое:")
		fmt.Scanln(&newValue)
		switch choise {
		case 1:
			task.Title = newValue
		case 2:
			task.Description = newValue
		case 3:
			task.Assignee = newValue
		case 4:
			task.Deadline = newValue
		case 5:
			task.Priority = newValue
		case 6:
			task.Status = newValue
		}
		Tasklist[id] = task
	} else {
		fmt.Println("Нет такого ID")
	}
}
func deleteTask() { //Функция удаляющая таск
	var id int
	fmt.Println("Введите ID задачи, которую хотите удалить:")
	fmt.Scanln(&id)
	_, ok := Tasklist[id]
	if ok {
		delete(Tasklist, id)
		fmt.Println("Задание удалено!")
	} else {
		fmt.Println("Задание не найдено!")
	}
}
func searchTaskByTitle() { //Функция выдающая таск по заголовку
	var title string
	fmt.Println("Введите название задачи, которую хотите найти:")
	fmt.Scanln(&title)
	found := false
	for id, task := range Tasklist {
		if task.Title == title {
			fmt.Printf("ID: %d, Название: %s, Описание: %s, Исполнитель: %s, Срок выполнения: %s, Приоритет: %s, Статус: %s\n", id, task.Title, task.Description, task.Assignee, task.Deadline, task.Priority, task.Status)
			found = true
		}
	}
	if !found {
		fmt.Println("Задача с таким названием не найдена!")
	}

}
func searchTaskByAssignee() { //Функция выдающая таск по исполнителю
	var assignee string
	fmt.Println("Введите название задачи, которую хотите найти:")
	fmt.Scanln(&assignee)
	found := false
	for id, task := range Tasklist {
		if task.Assignee == assignee {
			fmt.Printf("ID: %d, Название: %s, Описание: %s, Исполнитель: %s, Срок выполнения: %s, Приоритет: %s, Статус: %s\n", id, task.Title, task.Description, task.Assignee, task.Deadline, task.Priority, task.Status)
			found = true
		}
	}
	if !found {
		fmt.Println("Задача с таким исполнителем не найдена!")
	}

}
func weight(p string) int { //Нужно будет разделить функции по типу сортировки
	switch p {
	case "High":
		return 1
	case "Done":
		return 1
	case "Medium":
		return 2
	case "In progress":
		return 2
	case "Low":
		return 3
	case "Not started":
		return 3
	default:
		return 4
	}

}
func sortTasks() { //Функция сортирующая по приоритету
	var userInput string

	tasks := []Task{}
	for _, task := range Tasklist {
		tasks = append(tasks, task)
	}

	fmt.Println("По какому признаку сортировать:\n1. Приоритет\n2. Статус")
	fmt.Scanln(&userInput)
	switch userInput {
	case "1":
		sort.Slice(tasks, func(i, j int) bool {
			return weight(tasks[i].Priority) < weight(tasks[j].Priority)
		})

		for _, task := range tasks {
			fmt.Printf("Название: %s, Приоритет: %s\n", task.Title, task.Priority)
		}
	case "2":
		sort.Slice(tasks, func(i, j int) bool {
			return weight(tasks[i].Status) < weight(tasks[j].Status)
		})

		for _, task := range tasks {
			fmt.Printf("Название: %s, Приоритет: %s\n", task.Title, task.Status)
		}
	}
}

func main() {
	var userInput string
	for {
		fmt.Println("Выберите действие:\n1. Просмотр всех задач \n2. Добавление новой задачи \n3. Редактирование существующей задачи \n4. Удаление задачи \n5. Поиск задачи по названию \n6. Поиск задачи по исполнителю \n7. Сортировка задач по приоритету\n8. Выход из программы")
		fmt.Scanln(&userInput)
		switch userInput {
		case "1":
			viewTasks()
		case "2":
			addTask()
		case "3":
			editTask()
		case "4":
			deleteTask()
		case "5":
			searchTaskByTitle()
		case "6":
			searchTaskByAssignee()
		case "7":
			sortTasks()
		case "8":
			fmt.Println("Выход из программы!")
			return
		default:
			fmt.Println("Вы ввели неверное значение! Попробуйте еще раз!")

		}
	}
}
